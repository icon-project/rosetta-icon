#-------------------------------------------------------------------------------
#
# 	Makefile for building target binaries.
#

# Configuration
BUILD_ROOT = $(abspath ./)
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
BUILD_NAME = "rosetta-icon"

.DEFAULT_GOAL := all
all: clean build

.PHONY: build clean
build:
	$(eval GOBUILD_LDFLAGS=$(LDFLAGS))
	@ echo "[#] go build main.go as $(BIN_DIR)/$(BUILD_NAME)"
	$(GOBUILD_ENVS) $(GOBUILD) $(GOBUILD_FLAGS) -o $(BIN_DIR)/$(BUILD_NAME)

clean:
	@$(RM) -r $(BIN_DIR)
