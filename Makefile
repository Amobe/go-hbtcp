ROOT_DIR = $(shell pwd)
GOPATH = ${ROOT_DIR}
GOBIN = ${ROOT_DIR}/bin
export GOPATH
export GOBIN
export GO111MODULE=on

PROJECT_NAME = go-hbtcp
PROJECT_DIR = ${ROOT_DIR}/src/${PROJECT_NAME}

all build clean mod test benchmark:
	@make -C ${PROJECT_DIR} $@