## rhoas context

Group, share and manage your rhoas services

### Synopsis

When working with RHOAS CLI developers interact with the commands that are connecting directly to the service instance. At the time CLI commands can only connect to a single instance of the service. 

rhoas context commands allow developers to:

Group services into context that can be used with a number of CLI commands.
Manage different contexts by switching, listing and removing users contexts 
Share context with others to use the same set of services
Generating configuration for connecting to the services from various platforms and tools


### Examples

```
## Select a context
$ rhoas context use --name my-context

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas context create](rhoas_context_create.md)	 - Create context
* [rhoas context list](rhoas_context_list.md)	 - List contexts
* [rhoas context status](rhoas_context_status.md)	 - Show status of the context
* [rhoas context use](rhoas_context_use.md)	 - Set the current context

