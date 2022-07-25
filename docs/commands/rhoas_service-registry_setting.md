## rhoas service-registry setting

Configure settings of a Service Registry instance

### Synopsis

Configure settings of a Service Registry instance


### Examples

```
## List all settings of the current Service Registry instance
$ rhoas service-registry setting list

## Set the value of setting
$ rhoas service-registry setting set --name registry.ccompat.legacy-id-mode.enabled --value true

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry setting get](rhoas_service-registry_setting_get.md)	 - Get value of the setting
* [rhoas service-registry setting list](rhoas_service-registry_setting_list.md)	 - List settings
* [rhoas service-registry setting set](rhoas_service-registry_setting_set.md)	 - Set value of the setting

