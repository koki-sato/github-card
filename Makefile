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
	go tool golangci-lint fmt
	# TODO: Use the following command when the issue is fixed. https://github.com/golang/go/issues/73279
	# go tool modernize -fix -test ./...
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@v0.18.1 -fix -test ./...

.PHONY: lint
lint: ## Lint all Go source codes.
	go tool golangci-lint run
	# TODO: Use the following command when the issue is fixed. https://github.com/golang/go/issues/73279
	# go tool modernize -test ./...
	go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@v0.18.1 -test ./...

.PHONY: lint-fix
lint-fix: ## Run golangci-lint and perform fixes
	go tool golangci-lint run --fix

.PHONY: lint-config
lint-config: ## Verify golangci-lint configuration
	go tool golangci-lint config verify

.PHONY: codegen
codegen: ## Generate code.
	go tool generate-github-colors
	$(MAKE) fmt

.PHONY: help
help: ## Show this help message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
