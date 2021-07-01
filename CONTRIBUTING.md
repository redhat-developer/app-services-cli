# Contributing to rhoas CLI

Thank you for contributing to the RHOAS CLI. See below for guides to help you contribute.

## Prerequisites

The following will need to be installed on your device in order to contribute to this project.

- [Go >= v1.16](https://golang.org/dl)
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

### `make generate`

Generates code based on comments in code. This is primarily used to generate interface stubs using [moq](https://github.com/matryer/moq).

### Testing

If you have the Go extension for VS Code, you can generate test stubs for a file, package or function. See [Go#Test](https://code.visualstudio.com/docs/languages/go#_test)

### `make test/unit`

Runs unit tests

## Using CLI with Mock RHOAS API

RHOAS SDK provides mock for all supported APIs.
To use mock you need to have NPM installed on your system and have free port 8000
To work and test CLI locally please follow the [mock readme](https://github.com/redhat-developer/app-services-sdk-js/tree/main/packages/api-mock) and then login into cli using dev profile:

```shell
rhoas login --api-gateway=http://localhost:8000
```

### `make mock-api/start`

Starts the mock all services Manager API and Instance API at [`http://localhost:8000`](http://localhost:8000).

### Logging in

To log in to the mock API, run `rhoas login against the local server` with your authentication token:

```shell
rhoas login --api-gateway=http://localhost:8000
```

## Internationalization

All text strings are placed in `./pkg/localize/locales` directory.

## Documentation

The main CLI documentation source files are stored in the `./pkg/localize/locales/en/cmd/` directory.

The CLI documentation output is generated in the `./docs` directory.

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

Execute `git tag v0.20.1` to create the release tag. Then execute `git push origin v0.20.1` to push to the tag to your remote (GitHub).
Once pushed, a [GitHub Action](https://github.com/redhat-developer/app-services-cli/actions/workflows/release.yml) will create a release on GitHub and upload the binaries.

> NOTE: To create a pre-release, the tag should have appropriate suffix, e.g v0.20.1-alpha1

### Generate a changelog

> NOTE: This step is not required for pre-releases.

[git-chglog](https://github.com/git-chglog/git-chglog) is used to generate a changelog for the current release.

Run `./scripts/generate-changelog.sh` to output the changes between the current and last stable releases. Paste the output into the description of the [release on GitHub](https://github.com/redhat-developer/app-services-cli/releases/tag/latest).
