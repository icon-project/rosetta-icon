#-------------------------------------------------------------------------------
#
# 	Makefile for building target binaries.
#

# Configuration
BUILD_ROOT = $(abspath ./)
BIN_DIR = ./bin
DST_DIR = /usr/local/bin

UNAME = $(shell uname)

GOBUILD = go build
GOTEST = go test
GOTOOL = go tool
GOMOD = go mod
GOBUILD_TAGS =
GOBUILD_ENVS = CGO_ENABLED=0 GO111MODULE=on
GOBUILD_LDFLAGS = -ldflags ""
GOBUILD_FLAGS = -mod vendor $(GOBUILD_TAGS) $(GOBUILD_LDFLAGS)

# Build flags
GL_VERSION ?= $(shell git describe --always --tags --dirty)
GL_TAG ?= latest
BUILD_INFO = tags($(GOBUILD_TAGS))-$(shell date '+%Y-%m-%d-%H:%M:%S')

# Build flags for each command
LDFLAGS = -ldflags "-X 'main.version=$(GL_VERSION)' -X 'main.build=$(BUILD_INFO)'"
BUILD_NAME = "rosetta-icon"

.DEFAULT_GOAL := all
all : clean build ## Build the tools for current OS

.PHONY: build linux darwin help
build :
	$(eval GOBUILD_LDFLAGS=$(LDFLAGS))
	@ echo "[#] go build main.go as $(BIN_DIR)/$(BUILD_NAME)"
	$(GOBUILD_ENVS) $(GOBUILD) $(GOBUILD_FLAGS) -o $(BIN_DIR)/$(BUILD_NAME)

linux : ## Build the tools for linux
	@ echo "[#] build for $@"
	@ make GOBUILD_ENVS="$(GOBUILD_ENVS) GOOS=$@ GOARCH=amd64"

darwin : ## Build the tools for OS X
	@ echo "[#] build for $@"
	@ make GOBUILD_ENVS="$(GOBUILD_ENVS) GOOS=$@ GOARCH=amd64"

modules : ## Update modules in vendor/
	$(GOMOD) tidy
	$(GOMOD) vendor

clean : ## Remove generated files
	@$(RM) -r $(BIN_DIR)

TARGET_MAX_CHAR_NUM=20
help : ## This help message
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
	IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/:.*//' -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "  %-20s %s %s\n" $$help_command $$help_info; \
	done
