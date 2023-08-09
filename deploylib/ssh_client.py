import time
import paramiko

from .exceptions import CommandFailedException, ContainerStillRunningException

class SSHClient(paramiko.SSHClient):
    _DEFAULT_CMD_TIMEOUT_SEC = 10
    _CONTAINER_STOP_GRACE_SEC = 10

    def run_command(self, command: str, raise_on_error: bool = True) -> str:
        print(f'Running command: "{command}"')
        _, stdout, _ = self.exec_command(
            command, timeout=self._DEFAULT_CMD_TIMEOUT_SEC)
        status_code = stdout.channel.recv_exit_status()
        if status_code != 0 and raise_on_error:
            raise CommandFailedException(f'Command "{command}" failed with {status_code}')
        return stdout.read().decode('utf-8').strip()
    
    def get_module_active_container(self, module_name: str) -> str:
        try:
            container_id = self.run_command(f'docker ps -a | grep {module_name}_').split()[0]
            return container_id
        except CommandFailedException:
            return None
    
    def stop_module_container(self, module_name: str):
        container_id = self.get_module_active_container(module_name)
        if container_id is None:
            return
        self.run_command(f'docker stop {container_id}')
        time.sleep(self._CONTAINER_STOP_GRACE_SEC)
        container_id = self.get_module_active_container(module_name)
        if container_id is not None:
            raise ContainerStillRunningException(container_id)