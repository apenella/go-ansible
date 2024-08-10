# dafault target
.DEFAULT_GOAL: help

DOCKER_COMPOSE_BINARY := $(shell docker compose version > /dev/null 2>&1 && echo "docker compose" || (which docker-compose > /dev/null 2>&1 && echo "docker-compose" || (echo "docker compose not found. Aborting." >&2; exit 1)))

GOLANG_BINARY := $(DOCKER_COMPOSE_BINARY) run --rm golang-ci go
GOLANGCI_LINT_BINARY := $(DOCKER_COMPOSE_BINARY) run --rm golangci-lint

## Colors
COLOR_GREEN=\033[0;32m
COLOR_RED=\033[0;31m
COLOR_BLUE=\033[0;34m
COLOR_END=\033[0m

help: ## Lists available targets
	@echo
	@echo "Makefile usage:"
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1;32m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

static-analysis: vet golangci-lint ## Executes all static analysis tools
test: unit-test

vet: ## Executes the go vet
	@echo
	@echo "$(COLOR_GREEN) Executing go vet $(COLOR_END)"
	@echo
	@$(GOLANG_BINARY) vet ./pkg/...

golangci-lint: ## Executes golangci-lint
	@echo
	@echo "$(COLOR_GREEN) Executing golangci-lint $(COLOR_END)"
	@echo
	@$(GOLANGCI_LINT_BINARY) run

unit-test: ## Run unit tests
	@echo
	@echo "$(COLOR_GREEN) Running unit tests...$(COLOR_END)"
	@echo
	@$(GOLANG_BINARY) test ./pkg/... -cover -count=1

list-examples: ## List all examples
	@echo
	@echo "$(COLOR_GREEN) Listing all examples...$(COLOR_END)"
	@echo
	@ls -1 examples

attach-golang-ci: ## Attach to golang-ci Docker compose service
	@echo
	@echo "$(COLOR_GREEN) Attaching to golangci-lint Docker compose service...$(COLOR_END)"
	@echo
	@$(DOCKER_COMPOSE_BINARY) run --rm golang-ci /bin/sh
