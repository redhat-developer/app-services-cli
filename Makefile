.DEFAULT_GOAL := help
SHELL = bash

# The details of the application:
binary:=rhmas

# The version needs to be different for each deployment because otherwise the
# cluster will not pull the new image from the internal registry:
version:=0.1.0


# Enable Go modules:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export GOPRIVATE=gitlab.cee.redhat.com

ifndef SERVER_URL
	SERVER_URL:=http://localhost:8000
endif


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
	@echo "make code/fix             	format files"
	@echo "make openapi/pull			pull openapi definition"
	@echo "make openapi/generate     	generate openapi modules"
	@echo "make openapi/validate     	validate openapi schema"
	@echo "make clean                	delete temporary generated files"
						
	@echo "$(fake)"
.PHONY: help


# Checks if a GOPATH is set, or emits an error message
check-gopath:
ifndef GOPATH
	$(error GOPATH is not set)
endif
.PHONY: check-gopath

# Verifies that source passes standard checks.
verify: check-gopath
	go vet \
		./cmd/... \
		./client/... \
		./test/...
.PHONY: verify

# Runs our linter to verify that everything is following best practices
# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
lint:
	$(GOLANGCI_LINT_BIN) run \
		./cmd/... \
		./client/... \
		./test/...
.PHONY: lint

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary: check-gopath
	go build -o ${binary} ./cmd 
.PHONY: binary

# Runs the unit tests.
#
# Args:
#   TESTFLAGS: Flags to pass to `go test`. The `-v` argument is always passed.
#
# Examples:
#   make test TESTFLAGS="-run TestSomething"
test: install
	OCM_ENV=testing gotestsum --format $(TEST_SUMMARY_FORMAT) -- -p 1 -v -count=1 $(TESTFLAGS) \
		./client/... \
		./cmd/...
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
code/fix:
	@go mod tidy
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
.PHONY: code/fix



