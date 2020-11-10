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

## Architecture

![./architecture.png](./resources/architecture.png)


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
go run ./cmd/rhmas kafka create --name=test --multi-az="true" --provider=aws --region=eu-west-1
```

## Generating documentation

You can generate the latest commands documentation by executing the CLI with the `GENERATE_DOCS` environment variable
```shell
GENERATE_DOCS=true rhmas docs
```

Aftr running the command, the documentation should be updated. If there were new commands added we need to update `sidebars.json` file. 
with content that was printed into stdout.
