
DOCKER_COMPOSE_BINARY := $(shell docker compose version > /dev/null 2>&1 && echo "docker compose" || (which docker-compose > /dev/null 2>&1 && echo "docker-compose" || (echo "docker compose not found. Aborting." >&2; exit 1)))

PROJECT_NAME := go-ansible-$(shell basename ${PWD})

# dafault target
.DEFAULT_GOAL: help

help: ## Lists available targets
	@echo
	@echo "Makefile usage:"
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[1;32m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

build: ## Build the Docker compose environment
	@$(DOCKER_COMPOSE_BINARY) build

up: ## Create and start containers
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) up --detach --build

down: ## Stop and remove containers, networks, and volumes
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) down --volumes --remove-orphans --timeout 3

restart: down up ## Restart the containers

ps: ## List containers
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) ps

logs: ## Show all logs
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) logs

attach-ansible: ## Attach to the ansible container
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) exec --workdir /code/examples/$$(basename $$(pwd)) ansible /bin/sh

project-name: ## Show the project name
	@echo $(PROJECT_NAME)

run: ## Run the playbook
	@$(DOCKER_COMPOSE_BINARY) --project-name $(PROJECT_NAME) run --build --rm --workdir /code/examples/$$(basename $$(pwd)) ansible go run $$(basename $$(pwd)).go
