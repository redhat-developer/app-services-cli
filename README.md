## RHMAS CLI

This repository contains prototypes for CLI, Operator and many other artifacts 
used by developer experience team.

## Prerequisites

* [OpenAPI Generator](https://openapi-generator.tech/docs/installation/)
* [Golang](https://golang.org/dl/)


## Repository structure

- CLI (./cmd) - CLI for managed-services-api written in Golang
- SDK for MAS (./client) - API client written in golang that will be used in the CLI
- Mock  (./mas-mock) - Mock server for managed API (used only to demo flows an extra cases)
- website - documentation for the cli

## Development commands

```
OpenShift Managed Services CLI

make verify                     verify source code
make lint                       run golangci-lint
make binary                     compile binaries
make install                    compile binaries and install in GOPATH bin
make test                       run unit tests
make test-integration           run integration tests
make code/fix                   format files
make generate                   generate go and openapi modules
make openapi/pull                       pull openapi definition
make openapi/generate           generate openapi modules
make openapi/validate           validate openapi schema
make clean                      delete temporary generated files
```

## Development setup

1. Start mock server
```
cd mock
yarn install
yarn start
```

2. Build cli

```
make binary
```

3. Execute commands

```
./rhmas kafka list
```

## Development commands

When working with cmd we can execute commands using go run

```
go run ./cmd/rhmas kafka create --name=test --multi-az="true" --provider=aws --region=us-east-1
```

## Generating documentation

You can generate the latest commands documentation by executing the CLI with the `GENERATE_DOCS` environment variable
```shell
GENERATE_DOCS=true rhmas docs
```

Aftr running the command, the documentation should be updated. If there were new commands added we need to update `sidebars.json` file. 
with content that was printed into stdout.

## Performing releases

Releases can be triggered directing using github releases. 
Before performing release please make sure that numer of actions were performed.

- Run ./scripts/pullapi.sh script to make sure that CLI is using latest version of the MAS API
- Change ./pkg/version.go that will correspond to the semver
- Push all required changes to main branch

After Go to github releases and create new release.

### Releasing snapshot version

For testing purposes we should always release snapshot version that will not be used by end users.
To release snapshot version please execute

```
goreleaser --snapshot --rm-dist
```