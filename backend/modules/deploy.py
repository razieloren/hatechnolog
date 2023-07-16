#!/usr/bin/env python3

import yaml
import time
import docker
import logging
import argparse
import paramiko
import coloredlogs

from docker.models.images import Image

logger = logging.getLogger('deploy')
coloredlogs.install(level='DEBUG', logger=logger, fmt='%(asctime)s %(name)s %(levelname)s %(message)s')

class CommandFailedException(Exception):
    pass

class ContainerStillRunningException(Exception):
    pass

class SSHClient(paramiko.SSHClient):
    _DEFAULT_CMD_TIMEOUT_SEC = 10
    _CONTAINER_STOP_GRACE_SEC = 5

    def run_command(self, command: str, raise_on_error: bool = True) -> str:
        _, stdout, _ = self.exec_command(
            command, timeout=self._DEFAULT_CMD_TIMEOUT_SEC)
        status_code = stdout.channel.recv_exit_status()
        if status_code != 0 and raise_on_error:
            raise CommandFailedException(f'Command "{command}" failed with {status_code}')
        return stdout.read().decode('utf-8').strip()
    
    def get_module_active_container(self, module_name: str) -> str:
        try:
            container_id = self.run_command(f'docker ps -a | grep {module_name}_').split()[0]
            logger.info(f'Module "{module_name}" is running under container {container_id}')
            return container_id
        except CommandFailedException:
            logger.info(f'Module "{module_name}" does not seem to run')
        return None
    
    def stop_module_container(self, module_name: str):
        container_id = self.get_module_active_container(module_name)
        if container_id is None:
            return
        logger.info(f'Stopping container {container_id}')
        self.run_command(f'docker stop {container_id}')
        time.sleep(self._CONTAINER_STOP_GRACE_SEC)
        container_id = self.get_module_active_container(module_name)
        if container_id is not None:
            raise ContainerStillRunningException(container_id)


def yaml_file(file_path: str):
    with open(file_path, 'r') as stream:
        return yaml.safe_load(stream)['config']

def parse_args():
    args = argparse.ArgumentParser(description='Modules deployment script')
    args.add_argument('--config', '-c', help='Path to "deploy.yaml" file (Defauly: "./deploy.yaml")', default="deploy.yaml", type=yaml_file)
    return args.parse_args()

def create_ssh_client(host: str, port: int, username: str, password: str, key_file: str, key_pass: str) -> SSHClient:
    client = SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(host, port, username=username, password=password, 
                   key_filename=key_file, passphrase=key_pass)
    return client

def push_docker_image(module_name: str, version: str, host: str, registry: str, username: str, password: str) -> str:
    image_name = f'hatechnolog/modules/{module_name}-linux:{version}'
    docker_client = docker.from_env()
    try:
        image: Image = docker_client.images.get(image_name)
        taggedImage = f'{host}/{registry}/{image_name}'
        image.tag(taggedImage)
        try:
            docker_client.login(registry=host, username=username, password=password)
            docker_client.images.push(taggedImage)
        finally:
            docker_client.images.remove(taggedImage)
        return taggedImage
    finally:
        docker_client.close()


def main():
    args = parse_args()
    docker = args.config['docker']
    for module_name, module_props in args.config['modules'].items():
        try:
            metadata = module_props['metadata']
            version = f'{metadata["majorVersion"]}.{metadata["minorVersion"]}'
            logger.info(f'Deploying module "{module_name}:{version}"')
            taggedImage = push_docker_image(module_name, version, docker['host'], docker['registry'], docker['user'], docker['password'])
            ssh_props = module_props['ssh']
            ssh_client = create_ssh_client(ssh_props['host'], ssh_props['port'], 
                                        ssh_props['user'], ssh_props['password'], 
                                        ssh_props['identityFile'], 
                                        ssh_props['identityFilePassword'])
            try:
                ssh_client.stop_module_container(module_name)
                ssh_client.run_command(f'docker login {docker["host"]} -u {docker["user"]} -p {docker["password"]}')
                ssh_client.run_command(f'docker pull {taggedImage}')
                container_name = f'{module_name}_{version}'
                container_id = ssh_client.run_command(f'docker run --name {container_name} --rm -d {taggedImage}')[:12]
                logger.info(f'Module "{module_name}" new container ID: {container_id}')
                time.sleep(10)
                if ssh_client.get_module_active_container(module_name) is None:
                    logger.error(f'Cannot find "{module_name}"\'s container, it might be dead :(')
            finally:
                ssh_client.close()
        except Exception as e:
            logger.error(f'Unexpected error occured: {e}')
        logger.info(f'Done deploying module "{module_name}"')

if __name__ == '__main__':
    main()