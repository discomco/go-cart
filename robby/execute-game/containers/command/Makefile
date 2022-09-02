.PHONY: default
default: help

.PHONY: install-tools
install-tools: ## Install tools
	go install gotest.tools/gotestsum
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: tests
tests: ## Run all tests
	gotestsum -- -vet=off -race ./...

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: generate
generate: ## Generate implementations defined by schemas. 
	go generate -v ./...


.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'


.PHONY: build-dev
build-dev: ## Build the project
	go build -o ~/work/apps/ce/cmd/ce-cmd ./cmd  &&\
    cp -r ./config ~/work/apps/ce/cmd/
