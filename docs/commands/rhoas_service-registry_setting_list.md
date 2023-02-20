## rhoas service-registry setting list

List settings for a Service Registry instance

### Synopsis

List all settings for a Service Registry instance with their values and types

```
rhoas service-registry setting list [flags]
```

### Examples

```
## List all settings for the current Service Registry instance
$ rhoas service-registry setting list

## List all settings for a specific Service Registry instance
$ rhoas service-registry setting list --instance-id=8ecff228-1ffe-4cf5-b38b-55223885ee00

```

### Options

```
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings for a Service Registry instance

