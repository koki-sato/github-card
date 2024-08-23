.DEFAULT_GOAL := help

.PHONY: setup
setup: ## Resolve dependencies with `go mod` and install tools.
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/polyfloyd/go-errorlint@latest

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
	goimports -w .

.PHONY: lint
lint: ## Run all linters.
	go vet ./...
	staticcheck ./...
	go-errorlint -errorf=true ./...
	gosec -quiet ./...

.PHONY: fix
fix: ## Fix all linter errors.
	go-errorlint -errorf=true -fix ./...

.PHONY: help
help: ## Show this help message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
