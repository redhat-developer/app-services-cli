## rhoas service-registry

Service Registry commands

### Synopsis

Manage and interact with your Service Registry instances directly from the command line.

Create new Service Registry instances and interact with them by adding schema and API artifacts and downloading them to your computer.

Commands are divided into the following categories:

* Instance management commands: create, list, and so on
* Commands executed on selected instance: artifacts
* "use" command that selects the current instance


### Examples

```
## Create Service Registry instance
rhoas service-registry create --name myregistry

## List Service Registry instances
rhoas service-registry list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts
* [rhoas service-registry create](rhoas_service-registry_create.md)	 - Create a Service Registry instance
* [rhoas service-registry delete](rhoas_service-registry_delete.md)	 - Delete a Service Registry instance
* [rhoas service-registry describe](rhoas_service-registry_describe.md)	 - Describe a Service Registry instance
* [rhoas service-registry list](rhoas_service-registry_list.md)	 - List Service Registry instances
* [rhoas service-registry role](rhoas_service-registry_role.md)	 - Service Registry role management
* [rhoas service-registry rule](rhoas_service-registry_rule.md)	 - Manage artifact rules in a Service Registry instance
* [rhoas service-registry setting](rhoas_service-registry_setting.md)	 - Configure settings of a Service Registry instance
* [rhoas service-registry use](rhoas_service-registry_use.md)	 - Use a Service Registry instance

