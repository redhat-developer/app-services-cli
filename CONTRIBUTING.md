# Contributing to rhoas CLI

Thank you for contributing to the RHOAS CLI. See below for guides to help you contribute.

## Prerequisites

The following will need to be installed on your device in order to contribute to this project.

- [Go >= v1.15](https://golang.org/dl)
- [golangci-lint](https://golangci-lint.run)
- [OpenAPI Generator](https://openapi-generator.tech/)
- [Yarn](https://classic.yarnpkg.com)
- [GoReleaser](https://goreleaser.com/) (optional)

## Development

### Running CLI commands

You can run the CLI commands files directly with `go run`. All commands and subcommands are in the `./pkg/cmd` folder.

```shell
go run ./cmd/rhoas kafka create --name=test --multi-az="true" --provider=aws --region=us-east-1
```

### Development commands

#### `make lint`

Runs a linter on the Go source code. Configuration can be found in `.golangci.yaml`.
There are a number of lint rules enabled. You can find a full list of rules [here](https://golangci-lint.run/usage/linters/) with usage and configuration guides.

#### `make install`

Builds a binary in the `$GOPATH/bin` directory. Can be executed globally as it is in your `$PATH`.

#### `make binary`

Builds an executable binary `rhoas` of the CLI in the project root. Executable only inside the workspace.

### `make test`

Runs unit and integration tests.

#### `make format`

Formats source code.

## Managed Services API

The CLI communicates with the Managed Services API. For this there is a generated API client in `./pkg/api/managedservices/client`. 

### Updating the API client

Please ensure you have the latest generated version. Follow the steps below to update the Managed Services API version.

#### `make openapi/pull`

Saves the latest version of the Managed Services OpenAPI specification file to the `./openapi` directory.

#### `make openapi/validate`

Validates the Managed Services OpenAPI specification file.

#### `make openapi/generate`

Generates a Golang API client in `./pkg/api/managedservices/client`.

## Mock API

The repo has a local mocked version of the Managed Services API in `./mas-mock`.
To work and test CLI locally please follow the mock readme and then login into cli using dev profile:

> rhoas login --url=http://locahost:8080

> The mock API can become outdated from the current state of the Managed Services API. If you want to work with it please ensure it uses the latest OpenAPI spec, making changes where necessary.

## Mock Authentication

The repo has a local Keycloak instance which replicates the production environment. To start the server run `make mock-api/start`.

```shell
rhoas login --auth-url http://localhost:80
```

### `make mock-api/start`

Starts the mocked Managed Services API server at [`http://localhost:8000`](http://localhost:8000).

This will also start a local Keycloak instance at [`http://localhost:8000`](http://localhost:8000) for authentication from the CLI.
When Keycloak is up and running, log in as an admin (username: `admin`, password: `admin`).
Next, you will need to import the custom realm and client for the RHOAS CLI by running `make mock-api/keycloak/import-realm`.
Once complete you should see the `rhoas-cli` client in a `sso-external` realm from the Keycloak admin panel.

### `make mock-api/keycloak/import-realm`

Imports a Keycloak realm and `rhoas-cli` client.

### Logging in

To log in to the mock API, run `rhoas login against the local server` with your authentication token:

```shell
rhoas login --url http://localhost:8000
```

If you don't have an authentication token, you can still use a faked one, provided it has the correct payload:

```shell
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImQ4NDgzMTAyLTRhYzAtNDQ0Mi1hZjMwLTAwYWExMDdjZDc5MCJ9.eyJpYXQiOjE2MDQ1OTAzNDAsImp0aSI6ImNjZjg1MmM5LWI5YWEtNDE3Ny1hYmU0LWZkYWU0NmZmNmIxMSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsInN1YiI6ImY6LTMzMGVkMmRiLWEwMWUtNDI2OC04ZTkzLTE5ZjhmOGM2YTUxYzpkZXZlbG9wZXIiLCJ0eXAiOiJPZmZsaW5lIiwiYXpwIjoiYXBpLXNlcnZpY2VzIiwibm9uY2UiOiI5OTBjOTI1NS0xNTI3LTRiMTItOTM5OS02YWM2ZGZkMDJmZWQiLCJzZXNzaW9uX3N0YXRlIjoiYTY4Y2U1ZjktZTBiNi00MTc0LTg1YWItMDdmNzBkOGYxZmU2Iiwic2NvcGUiOiJvcGVuaWQgb2ZmbGluZV9hY2Nlc3MifQ.WTfFifDGnPkJX-IQSbzPWRhBKE7Gq5E6SKq3e70jbNc"
```

## Documentation

The main CLI documentation can be found in the `./docs` folder.
The documentation website (`./website`) is built dynamically from the documentation files using [Docusaurus](https://docusaurus.io/).

### `make docs/open`

Opens the documentation website at [`http://localhost:3000`](http://localhost:3000).

### Generating documentation

Documentation can be generated from the CLI commands.

```shell
make docs/generate
```

#### `make docs/generate`

After running the command, the documentation should be updated. If there were new commands added we need to update `./website/sidebars.json` file
with content that was printed into stdout.

## Best practices

- Write clear and meaningful Git commit mesages following the [Conventional Commits specification](https://www.conventionalcommits.org)
- Provide clear documentation comments.
- Make sure you include a clear and detailed PR description, linking to the related issue when it exists.

## Releases

Releases can be triggered directing using Github Releases. 
Before performing release, do the following:

1. Run `./scripts/pullapi.sh` script to make sure that CLI is using latest version of the Managed Services API.
2. Change `./pkg/version.go` that will correspond to the new version.
3. Push all required changes to main branch

After that, go to Github Releases and create the new release.

> Note: The project follows [semantic versioning](https://semver.org/)

### Releasing snapshot version

For testing purposes we should always release snapshot version that will not be used by end users.
To release snapshot version please execute:

```shell
goreleaser --snapshot --rm-dist
```

### Working with mocked Kafka

For testing you can use localy hosted Kafka

1. Run local kafka 

```
cd mas-mock
docker-compose up -d
```

2. Use Kafdrop to monitor it

http://localhost:9000

3. In CLI execute use command

```
rhoas kafka use 324234234
```

4. Edit `clusterHost` in `~/.rhoascli.json` to point to `localhost:9092`

Your config should look as follows:
```
  ...
  "services": {
    "kafka": {
      "clusterHost": "localhost:9092",
      "clusterId": "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
      "clusterName": "serviceapi"
    }
  }
  ...
```