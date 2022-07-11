## rhoas connector

Connectors instance commands

### Synopsis

With Red Hat OpenShift Connectors, you can create and configure connections between Red Hat OpenShift Streams for Apache Kafka and third-party systems. You can configure Connectors that produce data (data source Connectors) and Connectors that specify where to send data (data sink Connectors).

A Connectors instance is an instance of a one of the supported Connectors.
Use the "connector" command to create, delete, and view a list of Connectors instances.


### Examples

```
   
# List of Connectors instances
rhoas connector list

# Create a Connectors instance
rhoas connector create --file=myconnector.json

# Create a Connectors instance from stdin
cat myconnector.json | rhoas connector create

# Update a Connectors instance
rhoas connector update --id=my-connector --file=myconnector.json

# Update a Connectors instance from stdin
cat myconnector.json | rhoas connector update

# Delete a Connectors instance with ID c9b71ucotd37bufoamkg
rhoas connector delete --id=c9b71ucotd37bufoamkg

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters
* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connectors namespace commands
* [rhoas connector start](rhoas_connector_start.md)	 - Start a Connectors instance
* [rhoas connector stop](rhoas_connector_stop.md)	 - Stop a Connectors instance
* [rhoas connector use](rhoas_connector_use.md)	 - Set the current connector instance

