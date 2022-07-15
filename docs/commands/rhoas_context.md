## rhoas context

Group, share and manage your rhoas services

### Synopsis

Group your service instances into reusable contexts.
Context can be used when running other rhoas commands or to generate service configuration.

A service context is a group of application service instances and their service identifiers. By using service contexts, you can group together application service instances that you want to use together.

After creating a service context, you can share it with other developers so that they can use the same group of application service instances.

You can also use the service context to automatically generate configuration files that you need to use those application service instances in other development platforms and tools. For example, the service context can generate the following types of configurations:

- Standard environment variables for use in local development and tooling
- Java properties files that can be used in Quarkus, Apache Kafka, and so on
- Service binding configuration and service connections for the RHOAS Operator
- Configuration for Helm Charts and Kubernetes

The service context is defined in a JSON file (`contexts.json`), and stored locally on your computer. To find the location of this file, use the "rhoas context status" command.

Note: To specify a custom location for the `contexts.json` file, set the $RHOAS_CONTEXT environment variable to the location you want to use. If you set $RHOAS_CONTEXT to "./rhoas.json", service contexts will be loaded from the current directory.


### Examples

```
# Set the current context
$ rhoas context use --name qa

# List contexts
$ rhoas context list

# Create a context to represent a group of development services
$ rhoas context create --name dev-env

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas context create](rhoas_context_create.md)	 - Create a service context
* [rhoas context delete](rhoas_context_delete.md)	 - Permanently delete a service context.
* [rhoas context list](rhoas_context_list.md)	 - List service contexts
* [rhoas context set-connector](rhoas_context_set-connector.md)	 - Set the current Connectors instance
* [rhoas context set-kafka](rhoas_context_set-kafka.md)	 - Set the current Kafka instance
* [rhoas context set-service-registry](rhoas_context_set-service-registry.md)	 - Use a Service Registry instance
* [rhoas context status](rhoas_context_status.md)	 - View the status of application services in a service context
* [rhoas context unset](rhoas_context_unset.md)	 - Unset services in context
* [rhoas context use](rhoas_context_use.md)	 - Set the current context

