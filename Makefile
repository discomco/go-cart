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


.PHONY: build-qr-cmd
build-CMD: ## Build the project
	go build -o ~/work/apps/quadratic-roots/CMD/runme ./examples/quadratic-roots/delivery/CMD/cartwheel/cmd  &&\
    cp -r ./examples/quadratic-roots/delivery/CMD/cartwheel/config ~/work/apps/quadratic-roots/CMD/

.PHONY: build-qr-cmd-img
build-qr-cmd-img: ## Build the OCI image
	docker build . \
    --build-arg CR_PAT=$(CR_PAT) \
	--build-arg APP=./examples/quadratic-roots/delivery/command/cartwheel \
	-t local/quadratic-cmd