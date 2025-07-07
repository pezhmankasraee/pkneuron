APP_NAME := pkneuron
BIN_DIR := bin

# Get the latest tag (if any)
TAG := $(shell git describe --tags --abbrev=0 2>/dev/null)
# Get the short commit hash
HASH := $(shell git rev-parse --short HEAD)
# Is HEAD exactly at the tag?
AT_TAG := $(shell [ "$$(git rev-parse HEAD)" = "$$(git rev-list -n 1 $(TAG) 2>/dev/null)" ] && echo 1 || echo 0)

ifeq ($(AT_TAG),1)
	VERSION := $(TAG)
else
	VERSION := $(HASH)
endif

BINARY := $(BIN_DIR)/$(APP_NAME)-$(VERSION)

.PHONY: all build run clean help

all: build

build:
	@echo "Building $(APP_NAME) with version $(VERSION)..."
	@mkdir -p $(BIN_DIR)
	GO111MODULE=on go build -o $(BINARY) .

run: build
	@echo "Running $(BINARY)..."
	@$(BINARY)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build   Build the application binary with the latest tag or commit hash"
	@echo "  run     Build and run the application"
	@echo "  clean   Remove built binaries"
	@echo "  help    Show this help message"
