.DEFAULT_GOAL := help

.PHONY: build
build: ## Build application.
	go build -trimpath -ldflags "-w -s" -o bin/

.PHONY: start
start: ## Start application.
	go run ./...

.PHONY: test
test: ## Test all packages.
	go test -cover -race ./...

.PHONY: fmt
fmt: ## Format all Go source codes.
	go tool goimports -w .

.PHONY: lint
lint: ## Run all linters.
	go vet ./...
	go tool staticcheck ./...
	go tool go-errorlint -errorf=true ./...
	go tool gosec -quiet ./...

.PHONY: fix
fix: ## Fix all linter errors.
	go tool go-errorlint -errorf=true -fix ./...

.PHONY: codegen
codegen: ## Generate code.
	go tool generate-github-colors
	$(MAKE) fmt

.PHONY: help
help: ## Show this help message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
