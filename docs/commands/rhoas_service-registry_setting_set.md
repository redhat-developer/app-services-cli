## rhoas service-registry setting set

Set value of the setting

### Synopsis

Set value of the setting to a specific value or reset to default

```
rhoas service-registry setting set [flags]
```

### Examples

```
## Set value of setting by name
$ rhoas service-registry setting set --setting-name registry.ccompat.legacy-id-mode.enabled --value true

## Reset value of setting by name
$ rhoas service-registry setting set --setting-name registry.ccompat.legacy-id-mode.enabled --default

```

### Options

```
      --default               Reset value of the setting to default
      --instance-id string    ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -n, --setting-name string   Name of the setting
      --value string          New value of the setting
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings of a Service Registry instance

