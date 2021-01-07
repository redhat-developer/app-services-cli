.DEFAULT_GOAL := help
SHELL = bash

RHOAS_VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)

GO_LDFLAGS := -X github.com/bf2fc6cc711aee1a0c2a/cli/internal/build.Version=$(RHOAS_VERSION) $(GO_LDFLAGS)

BUILDFLAGS :=

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

# The details of the application:
binary:=rhoas

managedservices_client_dir=./pkg/api/managedservices

# Enable Go modules:
export GO111MODULE=on

# Prints a list of useful targets.
help:
	@echo ""
	@echo "OpenShift Managed Services CLI"
	@echo ""
	@echo "make lint                 	run golangci-lint"
	@echo "make binary               	compile binaries"
	@echo "make test                 	run  tests"
	@echo "make format             		format files"
	@echo "make openapi/pull					pull openapi definition"
	@echo "make openapi/generate     	generate openapi modules"
	@echo "make openapi/validate     	validate openapi schema"

	@echo "$(fake)"
.PHONY: help

# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run cmd/... pkg/... internal/...
.PHONY: lint

generate:
	go generate ./...

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary:
	go build $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" -o ${binary} ./cmd/rhoas
.PHONY: binary

install:
	go install -trimpath $(BUILDFLAGS) -ldflags "${GO_LDFLAGS}" ./cmd/rhoas
.PHONY: install

# Runs the integration tests.
test/integration: install
	go test ./test/integration
.PHONY: test/integration

# Runs the integration tests.
test/unit: install
	go test ./pkg/...
.PHONY: test/unit

openapi/pull:
	wget -O ./openapi/managed-services-api.yaml --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/managed-services-api.yaml
.PHONY: openapi/pull

# validate the openapi schema
openapi/validate:
	openapi-generator-cli validate -i openapi/managed-services-api.yaml
.PHONY: openapi/validate

# generate the openapi schema
openapi/generate:
	openapi-generator-cli generate -i openapi/managed-services-api.yaml -g go --package-name managedservices -p="generateInterfaces=true" --ignore-file-override=$$(pwd)/.openapi-generator-ignore -o ${managedservices_client_dir}
	openapi-generator-cli validate -i openapi/managed-services-api.yaml
	# generate mock
	moq -out ./pkg/api/managedservices/default_api_mock.go ./pkg/api/managedservices DefaultApi
	gofmt -w ${managedservices_client_dir}
.PHONY: openapi/generate

mock-api/start: mock-api/server/start mock-api/client/start
.PHONY: mock-api/start

mock-api/server/start:
	cd mas-mock && docker-compose up -d
.PHONY: mock-api/server/start

mock-api/client/start:
	cd mas-mock && yarn && yarn start
.PHONY: mock-api/client/start

mock-api/keycloak/import-realm:
	node mas-mock/keycloak/initKeycloak.js

# clean up code and dependencies
format:
	@go mod tidy
	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format


docs/generate:
	GENERATE_DOCS=true go run ./cmd/rhoas
	./scripts/pandoc.sh
.PHONY: docs/generate
