# Whats new

Latest changes for the RHOAS CLI.
For information on what was included in latest release please refer to our [changelog](https://github.com/redhat-developer/app-services-cli/blob/main/CHANGELOG.md)

## Unrelased (main branch)

### Support for configuration of service registry
`rhoas service-registry setting get` for getting the value of setting. Supports `--name` flag to define name of the setting.

Name of the setting can also be prompted by running command in interactive mode:
```
rhoas service-registry setting get
```

`rhoas service-registry setting set` for configuring the value of setting. Supports `--name` flag to define name of the setting and `--value` to define new value of the setting. `--default` flag can be used to restore default value of the setting.

Name of the setting and the value can also be prompted by running command in interactive mode:
```
rhoas service-registry setting set
```

`rhoas service-registry setting list` for listing all settings of a service registry instance.

> NOTE: setting command is only for admins of the service registry instance

## 0.46.0

### Support for Marketplace Billing model

`rhoas kafka create` now supports `--billing-model` flag that allow users to specify the type billing details for your kafka instance.

`--billing-model` along with `--marketplace` and `--marketplace-account-id` flags can be used to specify the
billing details.

Billing details can also be prompted by running the command in interactive mode:

```
rhoas kafka create
```

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

`rhoas completion powershell` will enable developers to get command completions on Windows platforms

> NOTE: Feature have been provided by our community. We are looking for feedback on usability of Powershell extensions.

### Support for Marketplace Billing model

`rhoas kafka create` supports now `--marketplace` and `--marketplace-account-id` flags that allow users to specify AWS marketplace details for billing purposes. 
