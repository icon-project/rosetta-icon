#-------------------------------------------------------------------------------
#
# 	Makefile for building target binaries.
#

# Configuration
PWD = $(abspath ./)
BIN_DIR = ./bin

GOBUILD = go build
GOBUILD_TAGS =
GOBUILD_ENVS = CGO_ENABLED=0 GO111MODULE=on
GOBUILD_LDFLAGS = -ldflags ""
GOBUILD_FLAGS = $(GOBUILD_TAGS) $(GOBUILD_LDFLAGS)

# Build flags
GL_VERSION ?= $(shell git describe --always --tags --dirty)
GL_TAG ?= latest
BUILD_INFO = tags($(GOBUILD_TAGS))-$(shell date '+%Y-%m-%d-%H:%M:%S')

# Build flags for each command
LDFLAGS = -ldflags "-X 'main.version=$(GL_VERSION)' -X 'main.build=$(BUILD_INFO)'"
BUILD_NAME = rosetta-icon

.DEFAULT_GOAL := all
all: clean build

.PHONY: build clean
build:
	$(eval GOBUILD_LDFLAGS=$(LDFLAGS))
	@ echo "[#] go build main.go as $(BIN_DIR)/$(BUILD_NAME)"
	$(GOBUILD_ENVS) $(GOBUILD) $(GOBUILD_FLAGS) -o $(BIN_DIR)/$(BUILD_NAME)

clean:
	@$(RM) $(BIN_DIR)/$(BUILD_NAME)

build-docker:
	docker build -t $(BUILD_NAME):latest https://github.com/icon-project/rosetta-icon.git#main

build-local:
	docker build -t $(BUILD_NAME):latest .

run-mainnet-online:
	docker run -d --rm -v "$(PWD)/icon-data/mainnet:/data" \
	    -p 7080:7080 -p 8080:8080 \
	    -e ENDPOINT=http://localhost:9080 \
	    -e NETWORK=MAINNET -e MODE=ONLINE -e PORT=8080 \
	    --name rosetta-icon \
	    rosetta-icon:latest

run-lisbon-online:
	docker run -d --rm -v "$(PWD)/icon-data/lisbon:/data" \
	    -p 7080:7080 -p 8080:8080 \
	    -e ENDPOINT=http://localhost:9080 \
	    -e NETWORK=LISBON -e MODE=ONLINE -e PORT=8080 \
	    --name rosetta-icon \
	    rosetta-icon:latest

# Run Rosetta CLI
ROSETTA_CLI := rosetta-cli
ROSETTA_CLI_CONF ?= rosetta-cli-conf
NETWORK ?= local
check_spec:
	$(ROSETTA_CLI) check:spec --configuration-file $(ROSETTA_CLI_CONF)/config_$(NETWORK).json --all

check_construction:
	$(ROSETTA_CLI) check:construction --configuration-file $(ROSETTA_CLI_CONF)/config_$(NETWORK).json

check_data:
	$(ROSETTA_CLI) check:data --configuration-file $(ROSETTA_CLI_CONF)/config_$(NETWORK).json
