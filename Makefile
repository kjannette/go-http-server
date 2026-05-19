.PHONY: *

help: ## This help dialog.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 | "sort -u"}' $(MAKEFILE_LIST)

run: ## Start the application
	docker compose up --build --remove-orphans

down: ## Stop the application
	docker compose down --remove-orphans

setup-local: copy-config install-dependencies run ## Sets up your local environment

install-lint: ## Install Go lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-dependencies: ## Install go project dependencies 
	go mod download
	go mod vendor

copy-config: ## Copy config file
	cp .env.dist .env

lint: install-lint ## Run lint
	golangci-lint run

tests: down ## Run tests
	go test -v ./...
