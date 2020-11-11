.DEFAULT_GOAL := help
SHELL = bash

# The details of the application:
binary:=rhmas

managedservices_client_dir=./pkg/api/managedservices/client

# Enable Go modules:
export GO111MODULE=on

# Prints a list of useful targets.
help:
	@echo ""
	@echo "OpenShift Managed Services CLI"
	@echo ""
	@echo "make verify               	verify source code"
	@echo "make lint                 	run golangci-lint"
	@echo "make binary               	compile binaries"
	@echo "make test                 	run  tests"
	@echo "make format             		format files"
	@echo "make openapi/pull					pull openapi definition"
	@echo "make openapi/generate     	generate openapi modules"
	@echo "make openapi/validate     	validate openapi schema"
	@echo "make clean                	delete temporary generated files"
						
	@echo "$(fake)"
.PHONY: help


# Verifies that source passes standard checks.
verify:
	go vet \
		./cmd/... \
		./pkg/... \
		./test/...
.PHONY: verify

# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint 
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run cmd/... client/... pkg/...
.PHONY: lint

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary:
	go build -o ${binary} ./cmd/rhmas
.PHONY: binary

install:
	go install ./cmd/rhmas
.PHONY: install

# Runs the unit tests.
test: install
	go test ./test/integration
.PHONY: test


openapi/pull:
	wget -P ./openapi -O managed-services-api.yaml --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/managed-services-api.yaml
.PHONY: openapi/pull

# validate the openapi schema
openapi/validate:
	openapi-generator validate -i openapi/managed-services-api.yaml
.PHONY: openapi/validate

# generate the openapi schema
openapi/generate:
	openapi-generator generate -i openapi/managed-services-api.yaml -g go --package-name msapi --ignore-file-override=$$(pwd)/.openapi-generator-ignore -o ${managedservices_client_dir}
	openapi-generator validate -i openapi/managed-services-api.yaml
	gofmt -w ${managedservices_client_dir}
.PHONY: openapi/generate

# clean up code and dependencies
format:
	@go mod tidy
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
.PHONY: format

format/check:
	test -z $(gofmt -l ./cmd ./pkg)
.PHONY: format/check

