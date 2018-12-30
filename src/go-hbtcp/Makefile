PACKAGE = go-hbtcp
ROOT_DIR = $(shell pwd)
SERVER_BINARY = server
CLIENT_BINARY = client

GO = go
GOBIN ?= .

VERSION = $(shell git tag --points-at HEAD)
VERSION := $(if $(VERSION),$(VERSION),(no_tag))
COMMIT = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +%s)

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Revision=$(COMMIT) -X main.BuildTime=$(BUILDTIME)"

all: build

build: build_server build_client

build_server:
	$(GO) build $(LDFLAGS) -o $(GOBIN)/$(SERVER_BINARY) $(PACKAGE)/cmd/$(SERVER_BINARY)

build_client:
	$(GO) build $(LDFLAGS) -o $(GOBIN)/$(CLIENT_BINARY) $(PACKAGE)/cmd/$(CLIENT_BINARY)

clean:
	if [ -f $(CLIENT_BINARY) ] ; then rm $(CLIENT_BINARY) ; fi
	if [ -f $(SERVER_BINARY) ] ; then rm $(SERVER_BINARY) ; fi
	if [ -f $(CLIENT_BINARY) ] ; then rm $(CLIENT_BINARY) ; fi

mod:
	$(GO) mod init

test:
	$(GO) test `$(GO) list ./... | grep -v mock` -v

benchmark:
	$(GO) test ./... -v -bench . -benchmem
