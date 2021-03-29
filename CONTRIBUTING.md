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
go run ./cmd/rhoas kafka create "test" --provider=aws --region=us-east-1
```

### Development commands

#### `make githooks`

Run `make githooks` to install important Githooks
This will symlink the files from `./githooks` to `.git/hooks`

#### `make lint`

Runs a linter on the Go source code. Configuration can be found in `.golangci.yaml`.
There are a number of lint rules enabled. You can find a full list of rules [here](https://golangci-lint.run/usage/linters/) with usage and configuration guides.

#### `make install`

Builds a binary in the `$GOPATH/bin` directory. Can be executed globally as it is in your `$PATH`.

#### `make binary`

Builds an executable binary `rhoas` of the CLI in the project root. Executable only inside the workspace.

#### `make format`

Formats source code.

#### `make pkger`

Packages static assets into the binary.

### `make generate`

Generates code based on comments in code. This is primarily used to generate interface stubs using [moq](https://github.com/matryer/moq).

### Testing

If you have the Go extension for VS Code, you can generate test stubs for a file, package or function. See [Go#Test](https://code.visualstudio.com/docs/languages/go#_test)

### `make test/unit`

Runs unit tests

### `make test/integration`

Runs integration tests

## Kafka Service Fleet Manager API (KAS)

- The CLI communicates with the Kafka Service Fleet Manager (KAS) API. For this there is a generated API client in `./pkg/api/kas/client`. 
- The CLI uses the Strimzi API for topic management. The client is located in `./pkg/api/strimzi-admin/client`

### Updating the API client

Please ensure you have the latest generated version. Follow the steps below to update all API clients.

#### `make openapi/pull`

Saves the latest OpenAPI specification files to the `./openapi` directory.

#### `make openapi/validate`

Validates the OpenAPI specification files.

#### `make openapi/generate`

Generates a Golang API client in `./pkg/api/kas/client` and `./pkg/api/strimzi-admin/client`.

## Mock API

The repo has a local mocked version of the KAS API in `./mas-mock`.
To work and test CLI locally please follow the mock readme and then login into cli using dev profile:

> rhoas login --url=http://locahost:8080

> The mock API can become outdated from the current state of the KAS API. If you want to work with it please ensure it uses the latest OpenAPI spec, making changes where necessary.

## Mock Authentication

The repo has a local Keycloak instance which replicates the production environment. To start the server run `make mock-api/start`.

You will need to update the `config.AuthURL` variable to `http://localhost:8080/auth/realms/redhat-external`.

```shell
rhoas login
```

### `make mock-api/start`

Starts the mocked KAS API server at [`http://localhost:8000`](http://localhost:8000).

This will also start a local Keycloak instance at [`http://localhost:8000`](http://localhost:8000) for authentication from the CLI.
When Keycloak is up and running, log in as an admin (username: `admin`, password: `admin`).
Next, you will need to import the custom realm and client for the RHOAS CLI by running `make mock-api/keycloak/import-realm`.
Once complete you should see the `rhoas-cli-prod` client in a `sso-external` realm from the Keycloak admin panel.

### `make mock-api/keycloak/import-realm`

Imports a Keycloak realm and `rhoas-cli-prod` client.

### Logging in

To log in to the mock API, run `rhoas login against the local server` with your authentication token:

```shell
rhoas login --url http://localhost:8000
```

If you don't have an authentication token, you can still use a faked one, provided it has the correct payload:

```shell
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImQ4NDgzMTAyLTRhYzAtNDQ0Mi1hZjMwLTAwYWExMDdjZDc5MCJ9.eyJpYXQiOjE2MDQ1OTAzNDAsImp0aSI6ImNjZjg1MmM5LWI5YWEtNDE3Ny1hYmU0LWZkYWU0NmZmNmIxMSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsInN1YiI6ImY6LTMzMGVkMmRiLWEwMWUtNDI2OC04ZTkzLTE5ZjhmOGM2YTUxYzpkZXZlbG9wZXIiLCJ0eXAiOiJPZmZsaW5lIiwiYXpwIjoiYXBpLXNlcnZpY2VzIiwibm9uY2UiOiI5OTBjOTI1NS0xNTI3LTRiMTItOTM5OS02YWM2ZGZkMDJmZWQiLCJzZXNzaW9uX3N0YXRlIjoiYTY4Y2U1ZjktZTBiNi00MTc0LTg1YWItMDdmNzBkOGYxZmU2Iiwic2NvcGUiOiJvcGVuaWQgb2ZmbGluZV9hY2Nlc3MifQ.WTfFifDGnPkJX-IQSbzPWRhBKE7Gq5E6SKq3e70jbNc"
```

## Internationalization

All text strings are placed in `./locales` directory. Group files into usage groups: commands go together under `./locales/cmd/{whoami,kafka/serviceaccount}` etc. All others can be grouped by their package.

When you make a change to a locale file, you _must_ run `make pkger` to embed the static files into the Go code. Please ensure this step is done before committing.

## Documentation

The main CLI documentation can be found in the `./docs` folder.

### Generating documentation

Documentation can be generated from the CLI commands.

```shell
make docs/generate
```

#### `make docs/generate`

After running the command, the documentation should be generated in AsciiDoc format.

#### `make docs/generate-modular-docs`

After running the command, the `dist` directory will contain the documentation conforming to the modular docs specification.

## Best practices

- [Command Line Interface Guidelines](https://clig.dev/) is a great resource for writing command-line interfaces.
- Write clear and meaningful Git commit messages following the [Conventional Commits specification](https://www.conventionalcommits.org)
- Provide clear documentation comments.
- Make sure you include a clear and detailed PR description, linking to the related issue when it exists.
- Check out [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) when writing and reviewing Go code.

## Releases

This project follows [Semantic Versioning](https://semver.org/). Before creating a release, identify if it will be a major, minor, or patch release. In the following example, we will create a patch release `0.20.1`.

> NOTE: When creating a release it is good practice to create a pre-release first to monitor the integration tests to ensure everything works as it should.

### Create snapshot

For testing purposes we should always release a local snapshot version for testing (requires [GoReleaser](https://goreleaser.com/))

```shell
goreleaser --snapshot --rm-dist
```

### Creating the release

Execute `git tag 0.20.1` to create the release tag. Then execute `git push origin 0.20.1` to push to the tag to your remote (GitHub).
Once pushed, a [GitHub Action](https://github.com/bf2fc6cc711aee1a0c2a/cli/actions/workflows/release.yml) will create a release on GitHub and upload the binaries.

### Generate a changelog

> NOTE: This step is not required for pre-releases.

[git-chglog](https://github.com/git-chglog/git-chglog) is used to generate a changelog for the current release.

Run `./scripts/generate-changelog.sh` to output the changes between the current and last stable releases.

### Make the release public

> NOTE: This step is not required for pre-releases.

Because the bf2 GitHub organization is private, we must clone the release to a public repository [bf3fc6c/cli](https://github.com/bf3fc6c/cli/releases).

Once all assets have been upload to the release on GitHub, run `./scripts/clone-release.sh` to clone the release to bf3fc6c/cli.

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
      "clusterId": "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg"
    }
  }
  ...
```