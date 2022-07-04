# Whats new

Latest changes for the RHOAS CLI.
For information on what was included in latest release please refer to our [changelog](https://github.com/redhat-developer/app-services-cli/blob/main/CHANGELOG.md)

## Latest


## 0.44.0 

### Support for consuming and producing to a topic

`rhoas kafka topic produce` and `rhoas kafka topic consume` for producing and consuming messages to a kafka topic.

New `produce` and `consume` commands work without any configuration and will point to your current Kafka instance by default.
Commands can act as alternative to kcat and kafka bin scripts by providing unified interface and simplicity of use.

> NOTE: commands are released as technology preview. Some of the flags and arguments can change without any notice.

### Generate Config

New `--overwrite` flag for `rhoas generate-config` enables users to overwrite their current configuration

## One SSO provider used by CLI

RHOAS CLI would use only single SSO provided by default. 
Users will not longer see multiple tabs opened when login and communicate with two different authentication servers.

> NOTE: This change should not affect users and it have been documented only for informational reasons

## Windows Powershell Suggestions

`rhoas completion powershell` 

> NOTE: Feature have been provided by our community. We are looking for feedback on usability of Powershell extensions.

### Support for Marketplace Billing model

`rhoas kafka create` supports now `--marketplace` and `--marketplace-account-id` flags that allow users to specify AWS marketplace details for billing purposes. 
