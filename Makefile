.PHONY: *

help: 
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 | "sort -u"}' $(MAKEFILE_LIST)

check-docker: ## Verify Docker CLI available
	@command -v docker >/dev/null 2>&1 || { \
		echo "Error: docker not found."; \
		echo "Install Docker Desktop: https://docs.docker.com/desktop/setup/install/mac-install/"; \
		echo "Or run without Docker: make setup-local && make run-local"; \
		exit 1; \
	}

run: check-docker ## Start the application with Docker Compose
	docker compose up --build --remove-orphans

run-local: ## Start the application locally (without Docker)
	go run ./cmd/app

down: check-docker 
	docker compose down --remove-orphans

setup-local: copy-config install-dependencies ## Install config and Go dependencies

setup-docker: setup-local run ## Install dependencies and start with Docker Compose

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-dependencies: ## Install Go project dependencies
	go mod download

copy-config:
	@test -f .env || cp .env.dist .env

lint: install-lint ## Run lint
	golangci-lint run

tests: ## Run tests
	go test -v ./...

validation: lint tests ## Run lint and tests

build: ## Build the server binary
	go build -o bin/server ./cmd/app
