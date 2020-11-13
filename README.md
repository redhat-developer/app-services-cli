# RHMAS CLI

This repository contains prototypes for CLI, Operator and many other artifacts used by the Developer Experience team.

## Repository structure

- (`./cmd`, `./pkg`) - CLI for managed-services-api written in Golang
- (`./pkg/api/managedservices/client`) - Managed Services API client written used in the CLI.
- (`./mas-mock`) - Mock server for Managed Services API (used only to demo flows an extra cases)
- `./website` - documentation for the CLI

## Website

To view the website for this repository, run `make docs/open`. The website will open at [`http://localhost:3000`](http://localhost:3000).

## Contributing

Check out the [Contributing Guide](./CONTRIBUTING.md) to learn more about the repository and how you can contribute.