PACKAGE = go-hbtcp
ROOT_DIR = $(shell pwd)
CLIENT_BINARY = client
SERVER_BINARY = server

GO = go
GOBIN ?= .

VERSION = $(shell git tag --points-at HEAD)
VERSION := $(if $(VERSION),$(VERSION),(no_tag))
COMMIT = $(shell git rev-parse --short HEAD)
BUILDTIME = $(shell date +%s)

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Revision=$(COMMIT) -X main.BuildTime=$(BUILDTIME)"

all: build

build: build_server build_client

build_client: $(ROOT_DIR)/vendor
	$(GO) build $(LDFLAGS) -o $(GOBIN)/$(CLIENT_BINARY) $(PACKAGE)/cmd/$(CLIENT_BINARY)

build_server: $(ROOT_DIR)/vendor
	$(GO) build $(LDFLAGS) -o $(GOBIN)/$(SERVER_BINARY) $(PACKAGE)/cmd/$(SERVER_BINARY)

clean:
	if [ -f $(CLIENT_BINARY) ] ; then rm $(CLIENT_BINARY) ; fi
	if [ -f $(SERVER_BINARY) ] ; then rm $(SERVER_BINARY) ; fi

mod:
	$(GO) mod init

$(ROOT_DIR)/glide.lock: $(ROOT_DIR)/glide.yaml
	$(GLIDE) update
	@touch $@

$(ROOT_DIR)/vendor: $(ROOT_DIR)/glide.lock
	$(GLIDE) --quiet install
	@touch $@

vendor: $(ROOT_DIR)/vendor

test:
	$(GO) test `$(GO) list ./... | grep -v mock` -v

benchmark:
	$(GO) test ./... -v -bench . -benchmem