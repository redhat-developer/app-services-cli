## Usage data

If the user has consented to `rhoas` collecting usage data, the following data will be collected when a command is executed -

* command's ID
* command's pseudonymized error message and error type (in case of failure)
* OS type
* `rhoas` version in use

Note that these commands do not include `--help` commands. We do not collect data about help commands.

### Disabling Usage data

To disable collection of the usage data please set environment variable in your terminal:

```
RED_HAT_DISABLE_TELEMETRY=true
```