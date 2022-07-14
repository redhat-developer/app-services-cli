## rhoas connector type

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

# Start a Connectors instance with ID c9b71ucotd37bufoamkg
rhoas connector start --id=c9b71ucotd37bufoamkg

# Start the Connectors instance in the current context 
rhoas connector stop

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands
* [rhoas connector type list](rhoas_connector_type_list.md)	 - List connector types

