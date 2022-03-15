## rhoas context

Group, share and manage your rhoas services

### Synopsis

rhoas context commands allow developers to:

  * Group services into contexts that can be used with a number of CLI commands.
  * Manage different service contexts by switching, listing and removing service contexts 
  * Share context with others to use the same set of services
  * Generating configuration for connecting to the services from various platforms and tools


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
* [rhoas context kafka-use](rhoas_context_kafka-use.md)	 - Set the current Kafka instance
* [rhoas context list](rhoas_context_list.md)	 - List contexts
* [rhoas context service-registry-use](rhoas_context_service-registry-use.md)	 - Use a Service Registry instance
* [rhoas context status](rhoas_context_status.md)	 - View the status of your application services
* [rhoas context status](rhoas_context_status.md)	 - View the status of your application services
* [rhoas context use](rhoas_context_use.md)	 - Set the current context

