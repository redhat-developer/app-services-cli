.DEFAULT_GOAL := help
SHELL = bash

# see internal/build.go on build configurations
RHOAS_VERSION ?= "dev"
REPOSITORY_OWNER ?= "redhat-developer"
REPOSITORY_NAME ?= "app-services-cli"
TERMS_REVIEW_EVENT_CODE ?= "onlineService"
TERMS_REVIEW_SITE_CODE ?= "ocm"

GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.Version=$(RHOAS_VERSION) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.RepositoryOwner=$(REPOSITORY_OWNER) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.RepositoryName=$(REPOSITORY_NAME) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.TermsReviewEventCode=$(TERMS_REVIEW_EVENT_CODE) $(GO_LDFLAGS)
GO_LDFLAGS := -X github.com/redhat-developer/app-services-cli/internal/build.TermsReviewSiteCode=$(TERMS_REVIEW_SITE_CODE) $(GO_LDFLAGS)

BUILDFLAGS :=

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

# The details of the application:
binary:=rhoas

strimzi_admin_api_dir=./pkg/api/strimzi-admin/client
amsapi_dir=./pkg/api/ams/amsclient

# Enable Go modules:
export GO111MODULE=on

# Prints a list of useful targets.
help:
	@echo ""
	@echo "RHOAS CLI"
	@echo ""
	@echo "make lint                 	run golangci-lint"
	@echo "make binary               	compile binaries"
	@echo "make test                 	run  tests"
	@echo "make format             		format files"
	@echo "make openapi/pull					pull openapi definition"
	@echo "make openapi/generate     	generate openapi modules"
	@echo "make openapi/validate     	validate openapi schema"
	@echo "make pkger									bundle static assets"
	@echo "make docs/check						check if docs need to be updated"
	@echo "make docs/generate					generate the docs"

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

openapi/pull: openapi/strimzi-admin/pull
.PHONY: openapi/pull

openapi/validate: openapi/strimzi-admin/validate
.PHONY: openapi/validate

openapi/generate: openapi/strimzi-admin/generate
.PHONY: openapi/validate

openapi/strimzi-admin/pull:
	wget -O ./openapi/strimzi-admin.yaml https://raw.githubusercontent.com/strimzi/strimzi-admin/e45b7410c36a96866a417e7adb8646f05d8293b9/rest/src/main/resources/openapi-specs/rest.yaml
.PHONY: openapi/strimzi-admin/pull

# validate the openapi schema
openapi/strimzi-admin/validate:
	openapi-generator-cli validate -i openapi/strimzi-admin.yaml
.PHONY: openapi/strimzi-admin/validate

# generate the openapi schema
openapi/strimzi-admin/generate:
	openapi-generator-cli generate -i openapi/strimzi-admin.yaml -g go --package-name strimziadminclient -p="generateInterfaces=true" --ignore-file-override=$$(pwd)/.openapi-generator-ignore -o ${strimzi_admin_api_dir}
	openapi-generator-cli validate -i openapi/strimzi-admin.yaml
	# generate mock
	moq -out ${strimzi_admin_api_dir}/default_api_mock.go ${strimzi_admin_api_dir} DefaultApi
	gofmt -w ${strimzi_admin_api_dir}
.PHONY: openapi/strimzi-admin/generate

openapi/ams/generate:
	openapi-generator-cli generate -i openapi/ams.json -g go --package-name amsclient -p="generateInterfaces=true" --ignore-file-override=$$(pwd)/.openapi-generator-ignore -o ${amsapi_dir}
	# generate mock
	moq -out ${amsapi_dir}/default_api_mock.go ${amsapi_dir} DefaultApi
	gofmt -w ${amsapi_dir}
.PHONY: openapi/strimzi-admin/generate

mock-api/start: mock-api/client/start
.PHONY: mock-api/start

mock-api/server/start:
	cd mas-mock && docker-compose up -d
.PHONY: mock-api/server/start

mock-api/client/start:
	cd mas-mock && yarn && yarn start
.PHONY: mock-api/client/start

# clean up code and dependencies
format:
	@go mod tidy
	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format

# Symlink common git hookd into .git directory
githooks:
	ln -fs $$(pwd)/githooks/pre-commit .git/hooks
.PHONY: githooks

docs/check: docs/generate
	./scripts/check-docs.sh
.PHONY: docs/check

docs/generate:
	GENERATE_DOCS=true go run ./cmd/rhoas
.PHONY: docs/generate

docs/generate-modular-docs: docs/generate
	SRC_DIR=$$(pwd)/docs/commands DEST_DIR=$$(pwd)/dist go run ./cmd/modular-docs
.PHONY: docs/generate-modular-docs