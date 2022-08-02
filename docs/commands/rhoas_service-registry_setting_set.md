## rhoas service-registry setting set

Set value of the setting for a Service Registry instance

### Synopsis

Set the value of the Service Registry setting to a specific value or reset to default

```
rhoas service-registry setting set [flags]
```

### Examples

```
## Set value of setting by name
$ rhoas service-registry setting set --name registry.ccompat.legacy-id-mode.enabled --value true

## Reset value of setting by name
$ rhoas service-registry setting set --name registry.ccompat.legacy-id-mode.enabled --default

```

### Options

```
      --default              Restore value of the Service Registry setting to default
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -n, --name string          Name of the Service Registry setting
      --value string         New value of the Service Registry setting
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings for a Service Registry instance

