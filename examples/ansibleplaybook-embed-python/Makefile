DOCKER_COMPOSE_BINARY := $(shell docker compose version > /dev/null 2>&1 && echo "docker compose" || (which docker-compose > /dev/null 2>&1 && echo "docker-compose" || (echo "docker compose not found. Aborting." >&2; exit 1)))

PROJECT_NAME := go-ansible-$(shell basename ${PWD})

# dafault target
.DEFAULT_GOAL: help

help: ## Lists available targets
	@echo
	@echo "Makefile usage:"
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1;32m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

run: ## Run the example
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) run --build app
