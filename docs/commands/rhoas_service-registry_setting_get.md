## rhoas service-registry setting get

Get value of the setting

### Synopsis

Get formatted setting value and other information

```
rhoas service-registry setting get [flags]
```

### Examples

```
## Get setting by name
$ rhoas service-registry setting get --setting-name registry.ccompat.legacy-id-mode.enabled

## Get setting in yaml format by name
$ rhoas service-registry setting get --setting-name registry.ccompat.legacy-id-mode.enabled --output yaml

```

### Options

```
      --instance-id string    ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -o, --output string         Specify the output format. Choose from: "json", "yaml", "yml"
  -n, --setting-name string   Name of the setting
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings of a Service Registry instance

