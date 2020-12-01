# RHOAS CLI

Repository for the RHOAS CLI.

## Installing CLI

Go to [releases](https://github.com/bf2fc6cc711aee1a0c2a/cli/releases) to download the latest release for your operating system.

## Getting Started

1. Login to RHOAS

```shell
rhoas login --insecure
```

This will redirect you to log in to https://sso.redhat.com/realms/redhat-external with your browser. The `--insecure` flag is required as this uses self-signed certs.

```shell
rhoas login
```

2. Use available Kafka commands

```
rhoas kafka
```

## Documentation

[Documentation](./docs) 


## Contributing

Check out the [Contributing Guide](./CONTRIBUTING.md) to learn more about the repository and how you can contribute.
