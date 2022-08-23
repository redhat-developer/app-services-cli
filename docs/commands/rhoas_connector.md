## rhoas connector

Connectors commands

### Synopsis

With Red Hat OpenShift Connectors, you can create and configure connections between Red Hat OpenShift Streams for Apache Kafka and third-party systems. You can configure Connectors that retrieve data (data source Connectors) and Connectors that specify where to send data (data sink Connectors).

To get started:
1. Decide which type of connector you want to create.
   Use the "rhoas connector type list" command to see a list of connector types
2. Build a configuration file based on one of the connector types.
   Use the "rhoas connector build" command to build a configuration file.
3. Optionally, edit the configuration file.
   Use a text editor of your choice to edit the configuration file.
4. Create a Connectors instance by specfiying the configuration file.
   Use the "rhoas connector create" command to create a Connectors instance.
5. Start the Connectors instance by using the "rhoas connector start" command.
6. Stop the Connectors instance by using the "rhoas connector stop" command.


### Examples

```
   
# List all connector types
rhoas connector type list

# Build a Connectors configuration file named "my_aws_lambda_connector.json" that is based on the "aws_lambda_sink_0.1" connector type
rhoas connector build --name=my_aws_lambda_connector --type=--type=aws_lambda_sink_0.1

# Create a Connectors instance by specifying a configuration file
rhoas connector create --file=myconnector.json

# Update an existing Connectors instance by specifying a configuration file
rhoas connector update --id=my-connector --file=myconnector.json

# List of Connectors instances
rhoas connector list

# Start the Connectors instance with ID my-connector
rhoas connector start --id=my-connector

# Stop the current Connectors instance
rhoas connector stop

# Delete a Connectors instance with ID my-connector
rhoas connector delete --id=my-connector

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas connector build](rhoas_connector_build.md)	 - Build a Connectors instance
* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters
* [rhoas connector create](rhoas_connector_create.md)	 - Create a Connectors instance
* [rhoas connector delete](rhoas_connector_delete.md)	 - Delete a Connectors instance
* [rhoas connector describe](rhoas_connector_describe.md)	 - Get the details for the Connectors instance
* [rhoas connector list](rhoas_connector_list.md)	 - List of Connectors instances
* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connectors namespace commands
* [rhoas connector start](rhoas_connector_start.md)	 - Start a Connectors instance
* [rhoas connector stop](rhoas_connector_stop.md)	 - Stop a Connectors instance
* [rhoas connector type](rhoas_connector_type.md)	 - List and get details of the different connector types
* [rhoas connector update](rhoas_connector_update.md)	 - Update a Connectors instance
* [rhoas connector use](rhoas_connector_use.md)	 - Set the current Connectors instance

