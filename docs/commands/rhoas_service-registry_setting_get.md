## rhoas service-registry setting get

Get value of the setting for a Service Registry instance

### Synopsis

Get the formatted value of the setting and other information for a Service Registry instance

```
rhoas service-registry setting get [flags]
```

### Examples

```
## Get the setting for a Service Registry instance by name
$ rhoas service-registry setting get --name registry.ccompat.legacy-id-mode.enabled

## Get the setting for a Service Registry instance in YAML format by name
$ rhoas service-registry setting get --name registry.ccompat.legacy-id-mode.enabled --output yaml

```

### Options

```
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -n, --name string          Name of the setting for a Service Registry instance
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings for a Service Registry instance

