# RHOAS CLI

Repository for the RHOAS CLI.

## Installing

### macOS

#### Binary

Download the binary for your CPU architecture from [Releases](https://github.com/bf2fc6cc711aee1a0c2a/cli/releases).

Extract the binary onto your path:

```shell
sudo tar -C /usr/local/bin -xzf rhoas_${VERSION}_macOS_amd64.tar.gz
```

> Note: Replace `VERSION` with the release version of the CLI.

### Linux

#### Binary

Download the binary for your CPU architecture from [Releases](https://github.com/bf2fc6cc711aee1a0c2a/cli/releases).

Extract the binary onto your path:

```shell
sudo tar -C /usr/local/bin -xzf rhoas_${VERSION}_linux_amd64.tar.gz
```

> Note: Replace `VERSION` with the release version of the CLI.

#### Fedora/CentOS/RHEL

Download the `.rpm` for your CPU architecture from [Releases](https://github.com/bf2fc6cc711aee1a0c2a/cli/releases).

Install the RPM package:

```shell
sudo rpm -i rhoas_${VERSION}_linux_amd64.rpm
```

> Note: Replace `VERSION` with the release version of the CLI.

## Getting Started

### Log in to RHOAS

```shell
rhoas login
```

This will redirect you to log in securely at https://sso.redhat.com/auth/realms/redhat-external with your browser.

### Use available Kafka commands

```
rhoas kafka --help
```

## Documentation

[Documentation](./docs) 


## Contributing

Check out the [Contributing Guide](./CONTRIBUTING.md) to learn more about the repository and how you can contribute.
