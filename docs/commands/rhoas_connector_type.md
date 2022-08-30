## rhoas connector type

View a list of supported connector types

### Synopsis

List and get details of connector types that are available in the connector catalog. 

To see a list of all available connector types, use the "type list" command.
You can optionally use the "--search" flag to filter the requested results by Connector types that start with or contain text that you specify. 

To see a description of a specific connector type, use the "type details" command.


### Examples

```
# List all connector types
rhoas connector type list

# List all connector types that start with "Amazon"
rhoas connector type list --search=Amazon%

# Get all of the details for the connector type by specifying the type ID
rhoas connector type describe --type=aws_kinesis_sink_0.1

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands
* [rhoas connector type describe](rhoas_connector_type_describe.md)	 - Get the details of a connector type
* [rhoas connector type list](rhoas_connector_type_list.md)	 - List connector types

