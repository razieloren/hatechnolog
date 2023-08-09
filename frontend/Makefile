MAJOR_VERSION := 1
MINOR_VERSION := 6

LINUX_IMAGE_NAME := hatechnolog/frontend-linux:${MAJOR_VERSION}.${MINOR_VERSION}
DARWIN_IMAGE_NAME := hatechnolog/frontend-drawin:${MAJOR_VERSION}.${MINOR_VERSION}

.PHONY: all clean linux-image darwin-image images
default: all

linux-image: node_modules .next
	docker buildx build --platform linux/amd64 --output type=docker -t ${LINUX_IMAGE_NAME} .

darwin-image: node_modules .next
	docker build -t ${DARWIN_IMAGE_NAME} .

images: linux-image darwin-image

all: images

clean-images:
	docker image rm `docker image ls | grep hatechnolog/frontend | tr -s ' ' | cut -d' ' -f3 | tr '\n' ' '`

clean: clean-images
	rm -rf .next


