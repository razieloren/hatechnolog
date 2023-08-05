MAJOR_VERSION := 1
MINOR_VERSION := 1

OUT_DIR := out
PWD=$(shell pwd)
MODULE_NAME := query
SOURCES := $(shell find . -name '*.go')

LINUX_OUTPUT_DIR := ${OUT_DIR}/linux
LINUX_TARGET := ${LINUX_OUTPUT_DIR}/${MODULE_NAME}
LINUX_IMAGE_NAME := hatechnolog/modules/${MODULE_NAME}-linux:${MAJOR_VERSION}.${MINOR_VERSION}
DARWIN_OUTPUT_DIR := ${OUT_DIR}/darwin
DARWIN_TARGET := ${DARWIN_OUTPUT_DIR}/${MODULE_NAME}
DARWIN_IMAGE_NAME := hatechnolog/modules/${MODULE_NAME}-drawin:${MAJOR_VERSION}.${MINOR_VERSION}

GO_BUILD_CMD := CGO_ENABLED=0 go build -trimpath -a -ldflags="-w -s -buildid=" -o

.PHONY: all clean clean-binaries clean-images binaries images linux-image darwin-image
default: all

${LINUX_OUTPUT_DIR}: |
	mkdir -p $@

${DARWIN_OUTPUT_DIR}: |
	mkdir -p $@

${LINUX_TARGET}: ${SOURCES} | ${LINUX_OUTPUT_DIR}
	GOOS=linux GOARCH=amd64 ${GO_BUILD_CMD} $@

${DARWIN_TARGET}: ${SOURCES} | ${LINUX_OUTPUT_DIR}
	GOOS=darwin GOARCH=arm64 ${GO_BUILD_CMD} $@

binaries: ${LINUX_TARGET} ${DARWIN_TARGET}

linux-image: ${LINUX_TARGET} config.yaml
	docker buildx build --platform linux/amd64 --output type=docker -t ${LINUX_IMAGE_NAME} .

darwin-image: ${DARWIN_TARGET} config.yaml
	docker build -t ${DARWIN_IMAGE_NAME} .

images: linux-image darwin-image

all: binaries images

clean-binaries:
	rm -rf ${OUT_DIR}

clean-images:
	docker image rm `docker image ls | grep hatechnolog/modules/${MODULE_NAME} | tr -s ' ' | cut -d' ' -f3 | tr '\n' ' '`

clean: clean-binaries clean-images


