#!/usr/bin/env python3

import yaml
import time
import logging
import argparse
import paramiko
import coloredlogs

from deploylib.ssh_client import SSHClient
from deploylib.docker_client import DockerClient

logger = logging.getLogger('deploy')
coloredlogs.install(level='DEBUG', logger=logger, fmt='%(asctime)s %(name)s %(levelname)s %(message)s')


def yaml_file(file_path: str):
    with open(file_path, 'r') as stream:
        return yaml.safe_load(stream)['config']

def parse_args():
    args = argparse.ArgumentParser(description='Modules deployment script')
    args.add_argument('--config', '-c', help='Path to "deploy.yaml" file (Default: "./deploy.yaml")', default="deploy.yaml", type=yaml_file)
    args.add_argument('--module', '-m', help='Deploy only this module (if exists)', default=None)
    return args.parse_args()

def create_ssh_client(host: str, port: int, username: str, password: str, key_file: str, key_pass: str) -> SSHClient:
    client = SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(host, port, username=username, password=password, 
                   key_filename=key_file, passphrase=key_pass)
    return client

def main():
    args = parse_args()
    docker = args.config['docker']
    for module_name, module_props in args.config['modules'].items():    
        try:
            metadata = module_props['metadata']
            docker_args = module_props.get('docker', '')
            version = f'{metadata["majorVersion"]}.{metadata["minorVersion"]}'
            container_name = f'{module_name}_{version}'.replace('/', '_')
            module_logger = logger.getChild(container_name)
            if args.module is not None and module_name != args.module:
                module_logger.info("Skipping this module")
                continue
            module_logger.info('Deploy started')
            docker_client = DockerClient()
            try:
                module_logger.info('Pushing image to registry')
                taggedImage = docker_client.push_image(module_name, version, docker['host'], docker['registry'], docker['user'], docker['password'])
            finally:
                docker_client.close()
            ssh_props = module_props['ssh']
            ssh_client = create_ssh_client(ssh_props['host'], ssh_props['port'], 
                                        ssh_props['user'], ssh_props['password'], 
                                        ssh_props['identityFile'], 
                                        ssh_props['identityFilePassword'])
            try:
                module_logger.info('Stopping remote container')
                ssh_client.stop_module_container(module_name.replace('/', '_'))
                module_logger.info('Logging in to docker registry')
                ssh_client.run_command(f'docker login {docker["host"]} -u {docker["user"]} -p {docker["password"]}')
                module_logger.info(f'Pulling latest image: {taggedImage}')
                ssh_client.run_command(f'docker pull {taggedImage}')
                module_logger.info(f'Running container with name: {container_name}')
                container_id = ssh_client.run_command(f'docker run {docker_args} --name {container_name} --rm -d {taggedImage}')[:12]
                module_logger.info(f'New container ID: {container_id}')
                time.sleep(10)
                if ssh_client.get_module_active_container(module_name.replace('/', '_')) is None:
                    logger.error(f'Container is gone, it might be dead :(')
                module_logger.info(f'Done deploying :)')
            finally:
                ssh_client.close()
        except Exception as e:
            logger.error(f'Unexpected error occured: {e}')
    logger.info('All modules finished deploying, thanks god')

if __name__ == '__main__':
    main()