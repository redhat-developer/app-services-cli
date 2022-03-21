## rhoas context

Group, share and manage your rhoas services

### Synopsis

rhoas context commands allow developers to:

  * Group services into contexts that can be used with a number of CLI commands.
  * Manage different service contexts by switching, listing and removing service contexts 
  * Share context with others to use the same set of services
  * Generating configuration for connecting to the services from various platforms and tools

The loading order of contexts follows these rules:

  1. If $RHOAS_CONTEXT environment variable is set, then contexts are managed using only that file
  2. Otherwise, "<config-dir>/rhoas/contexts.json" is used for contexts



### Examples

```
## Set the current context
$ rhoas context use --name my-context

## List contexts
$ rhoas context list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas context create](rhoas_context_create.md)	 - Create a service context
* [rhoas context delete](rhoas_context_delete.md)	 - Delete a service context
* [rhoas context list](rhoas_context_list.md)	 - List contexts
* [rhoas context status](rhoas_context_status.md)	 - View the status of application services set in service context
* [rhoas context use](rhoas_context_use.md)	 - Set the current context
* [rhoas context use-kafka](rhoas_context_use-kafka.md)	 - Set the current Kafka instance
* [rhoas context use-service-registry](rhoas_context_use-service-registry.md)	 - Use a Service Registry instance

