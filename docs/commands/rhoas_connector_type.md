## rhoas connector type

List and get details of different connector types

### Synopsis

List, search and get details of connector types that are available to use in the connector catalog.

Use list to list available connector types.

Use describe to get more details about a specific connector type.


### Examples

```
# List all connector types
rhoas connector type list

# List all connector types that start with 'Amazon'
rhoas connector type list --search=Amazon%

# Get more details of connector type with a type id of IEJF87hg2342hsdHFG
rhoas connector type describe --id=IEJF87hg2342hsdHFG

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands
* [rhoas connector type describe](rhoas_connector_type_describe.md)	 - Get details of a connector type
* [rhoas connector type list](rhoas_connector_type_list.md)	 - List connector types

