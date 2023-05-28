#!/usr/bin/make -f

DOCKER := $(shell which docker)
DOCKERCOMPOSE := $(shell which docker-compose)
BUILDDIR ?= $(CURDIR)/build

clean:
	@echo "Cleaning up..."
	@rm -rf .testnets
	@echo "Done"

build:
	@echo "Building project..."
	@$(DOCKER) build -t vedhavyas/cf-sre . &> /dev/null
	@echo "Done"

setup-testnet:
	@echo "Creating node configurations..."
	@$(DOCKER) run -v $(CURDIR):/root vedhavyas/cf-sre:latest cf-sred testnet init-files --v 4
	@rm -rf .cf-sre
	@echo "Done"

start-testnet: build setup-testnet
	@echo "Starting local testnet..."
	@$(DOCKERCOMPOSE) up -d
	@echo "Done"

stop-testnet:
	@echo "Stopping local testnet..."
	@$(DOCKERCOMPOSE) down
	@$(MAKE) clean
	@echo "Done"

.PHONY: clean build setup-testnet start-testnet stop-testnet
