.DEFAULT_GOAL := help
SHELL = bash

# The details of the application:
binary:=rhmas


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
	@echo "make test                 	run unit tests"
	@echo "make test-integration     	run integration tests"
	@echo "make format             		format files"
	@echo "make openapi/pull			pull openapi definition"
	@echo "make openapi/generate     	generate openapi modules"
	@echo "make openapi/validate     	validate openapi schema"
	@echo "make clean                	delete temporary generated files"
						
	@echo "$(fake)"
.PHONY: help



# Verifies that source passes standard checks.
verify:
	go vet \
		./cmd/... \
		./client/... \
		./test/...
.PHONY: verify

# Runs our linter to verify that everything is following best practices
# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
lint:
	golangci-lint 
		./cmd/... \
		./client/... \
		./test/...
.PHONY: lint

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary:
	go build -o ${binary} ./cmd 
.PHONY: binary

install:
	go install ./cmd
.PHONY: install

# Runs the unit tests.
#
# Args:
#   TESTFLAGS: Flags to pass to `go test`. The `-v` argument is always passed.
#
# Examples:
#   make test TESTFLAGS="-run TestSomething"
test: install
	echo "go test - TODO"
.PHONY: test

# Precompile everything required for development/test.
test-prepare: install
	go test -i ./test/integration/...
.PHONY: test-prepare

openapi/pull:
	wget -P ./openapi -O managed-services-api.yaml --no-check-certificate https://gitlab.cee.redhat.com/service/managed-services-api/-/raw/master/openapi/managed-services-api.yaml
.PHONY: openapi/pull

# validate the openapi schema
openapi/validate:
	openapi-generator validate -i openapi/managed-services-api.yaml
.PHONY: openapi/validate

# generate the openapi schema
openapi/generate:
	openapi-generator generate -i openapi/managed-services-api.yaml -g go -o client/mas
	openapi-generator validate -i openapi/managed-services-api.yaml
	gofmt -w client/mas
.PHONY: openapi/generate

# clean up code and dependencies
format:
	@go mod tidy
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
.PHONY: format

formatCheck:
	test -z $(gofmt -l ./cmd)
.PHONY: formatCheck





