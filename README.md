# RHOAS CLI

Repository for the RHOAS CLI.

## Installing CLI

Go to [releases](https://github.com/bf2fc6cc711aee1a0c2a/cli/releases) to download the latest release for your operating system.

## Getting Started

### Login to RHOAS

```shell
rhoas login --insecure
```

This will redirect you to log in to https://sso.redhat.com/realms/redhat-external with your browser. The `--insecure` flag is required as this uses self-signed certs.

```shell
rhoas login
```

> NOTE: Work is ongoing to get a rhoas-cli client on Red Hat SSO. Until then you will not be able to interact with the control plane using this login flow. To workaround this, please use token-based login, which will be removed as soon as a client is available.

### Login with offline token

This login flow will not be available in the official release of the RHOAS CLI, but should be used to login to https://sso.redhat.com for now if you want to interact with the control plane API.

```shell
rhoas login --token $TOKEN
```

> NOTE: You can obtain an offline token from [cloud.redhat.com](https://cloud.redhat.com/openshift/token)

### Use available Kafka commands

```
rhoas kafka
```

## Documentation

[Documentation](./docs) 


## Contributing

Check out the [Contributing Guide](./CONTRIBUTING.md) to learn more about the repository and how you can contribute.
