.DEFAULT_GOAL := binary
SHELL = bash

# The details of the application:
binary:=rhrequest

# Enable Go modules:
export GO111MODULE=on

# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint: lint-lang ## Lint Go files for errors
	golangci-lint run cmd/... pkg/... internal/...
# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary: ## Compile the binary into the local project directory
	go build $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" -o ${binary} ./cmd/${binary}
.PHONY: binary

install: ## Compile and install and add it to the PAth 
	go install -trimpath $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" ./cmd/${binary}
.PHONY: install

test: ## Run unit tests
	go test ./pkg/... ./internal/...
.PHONY: test

format: ## Clean up code and dependencies
	@go mod tidy

	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format

# Check http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

