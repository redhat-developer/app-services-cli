# RHOAS CLI
<p align="center">
  <img alt="Logo" src="https://user-images.githubusercontent.com/11743717/127519981-97c76ae4-f17b-4ac8-8b4d-365bfa4a6374.png">
</p>

`rhoas` is a command-line client for managing Red Hat application services


## Installing RHOAS

See [Installing the rhoas CLI](https://github.com/redhat-developer/app-services-guides/tree/main/docs/rhoas/rhoas-cli-installation#installing-the-rhoas-cli) 
for instructions on how to install CLI from official sources.

To install or update to latest version of CLI use following script:

```shell
curl -o- https://raw.githubusercontent.com/redhat-developer/app-services-cli/main/scripts/install.sh | bash 
```

## RHOAS Container Image

To install or update [image](https://github.com/redhat-developer/app-services-tools) containing RHOAS CLI:

```shell
docker pull quay.io/rhoas/tools
```

Running the image:

```shell
docker run -ti --rm --name rhoas-devsandbox --entrypoint /bin/bash quay.io/rhoas/tools
```

## Guides

See our [Guides](https://github.com/redhat-developer/app-services-guides/tree/main/docs/rhoas/rhoas-cli-installation) for installation and usage instructions.

## Commands

See the [Command-Line Reference](http://appservices.tech/commands/rhoas) section for details of all available commands and options.

