.DEFAULT_GOAL := binary
SHELL = bash

# see internal/build.go on build configurations
RHOAS_VERSION ?= "dev"
REPOSITORY_OWNER ?= "redhat-developer"
REPOSITORY_NAME ?= "app-services-cli"
TERMS_SPEC_URL ?= "https://console.redhat.com/apps/application-services/terms-conditions-spec.json"
SSO_REDIRECT_PATH ?= "sso-redhat-callback"
MAS_SSO_REDIRECT_PATH ?= "mas-sso-callback"
BUILD_SOURCE ?= "local"

# see pkg/cmdutil/constants.go
DEFAULT_PAGE_NUMBER ?= "1"
DEFAULT_PAGE_SIZE ?= "10"

GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.Version=$(RHOAS_VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.RepositoryOwner=$(REPOSITORY_OWNER) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.RepositoryName=$(REPOSITORY_NAME) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.TermsReviewSpecURL=$(TERMS_SPEC_URL) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.DefaultPageSize=$(DEFAULT_PAGE_SIZE) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.DefaultPageNumber=$(DEFAULT_PAGE_NUMBER) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.SSORedirectPath=$(SSO_REDIRECT_PATH) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.MASSSORedirectPath=$(MAS_SSO_REDIRECT_PATH) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.BuildSource=$(BUILD_SOURCE) $(GO_LDFLAGS)

BUILDFLAGS :=

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

# The details of the application:
binary:=rhoas

amsapi_dir=./pkg/api/ams/amsclient
rbacapi_dir=./pkg/api/rbac/rbacclient

# Enable Go modules:
export GO111MODULE=on

# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint: ## Lint Go files for errors
	golangci-lint run cmd/... pkg/... internal/...

generate: ## Scan code for generate comments and run generators
	go generate ./...

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary: ## Compile the rhoas binary into the local project directory
	go build $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" -o ${binary} ./cmd/rhoas
.PHONY: binary

install: ## Compile and install rhoas and add it to the PAth 
	go install -trimpath $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" ./cmd/rhoas
.PHONY: install

test: ## Run unit tests
	go test ./pkg/...
.PHONY: test

generate-ams-sdk: ## Generate the Account Management Service SDK
	openapi-generator-cli generate -i openapi/ams.json -g go --package-name amsclient -p="generateInterfaces=true" --ignore-file-override=$$(pwd)/.openapi-generator-ignore -o ${amsapi_dir}
	# generate mock
	moq -out ${amsapi_dir}/default_api_mock.go ${amsapi_dir} DefaultApi
	gofmt -w ${amsapi_dir}
.PHONY: generate-ams-sdk

start-mock-api: ## Start the mock rhoas server
	npm install -g @rhoas/api-mock
	asapi --pre-seed
.PHONY: start-mock-api

format: ## Clean up code and dependencies
	@go mod tidy

	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format

check-docs: generate-docs ## Check whether reference documentation needs to be generated
	./scripts/check-docs.sh
.PHONY: check-docs

generate-docs: ## Generate command-line reference documentation
	rm -rf ./docs/commands/*
	go run ./cmd/rhoas docs --dir ./docs/commands --file-format adoc
.PHONY: generate-docs

generate-modular-docs: generate-docs ## Generate modular command-line reference documentation
	SRC_DIR=$$(pwd)/docs/commands DEST_DIR=$$(pwd)/dist go run ./cmd/modular-docs
.PHONY: generate-modular-docs

lint-lang:
	go install github.com/redhat-developer/app-services-go-linter/cmd/app-services-go-linter@latest
	app-services-go-linter
.PHONY: lint-lang

# Check http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

