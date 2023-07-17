BASE = $(CURDIR)
SRC = $(BASE)/src
MODULE = movies-api

.PHONY: all $(MODULE) swagger test-all coverage coverage-ci lint clean help dev
.DEFAULT_GOAL := help

all: $(MODULE) swagger

$(MODULE): ## Build the binary file
	@go build -v -o $(BASE)/bin/$@ $(BASE)/src/cmd/http/main.go

swagger: ## Generate swagger docs
	@swag init -d $(SRC)/cmd/http,$(SRC)/ -o $(SRC)/docs

test-all: ## Run all tests
	@go test -v $(SRC)/...

coverage: ## Run all tests with coverage
	@go test -coverprofile=coverage.out -v $(SRC)/...
	@go tool cover -html=coverage.out

coverage-ci: ## Run all tests with coverage for CI
	CVPKG=$(go list $(SRC)/... | grep -v mocks | tr '\n' ',')
	@go test -coverpkg=${CVPKG} -coverprofile=coverage.out -covermode=count  $(SRC)/...

lint: ## Run linter
	@revive -config ../../revive.toml --formatter friendly ./...

dev: ## Run dev server
	@air -c .air.toml

clean: ## Clean all generated files
	@rm -rfv bin
	@exit 0

help:
	@grep -hE '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

