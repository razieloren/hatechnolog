import time
import docker

from docker.models.images import Image

class DockerClient():
    def __init__(self):
        self._client = docker.from_env()

    def close(self):
        self._client.close()

    def push_image(self, module_name: str, version: str, host: str, registry: str, username: str, password: str) -> str:
        image_name = f'hatechnolog/{module_name}-linux:{version}'
        image: Image = self._client.images.get(image_name)
        taggedImage = f'{host}/{registry}/{image_name}'
        image.tag(taggedImage)
        try:
            self._client.login(registry=host, username=username, password=password)
            self._client.images.push(taggedImage)
            time.sleep(1)
        finally:
            self._client.images.remove(taggedImage)
        return taggedImage

    