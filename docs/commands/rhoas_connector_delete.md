## rhoas connector delete

Delete a Connectors instance

### Synopsis

Delete a Connectors instance by specifying its ID. Use the "connector list" command to see a list of all Connectors instances and their ID values.

```
rhoas connector delete [flags]
```

### Examples

```
# Delete a Connectors instance with ID myconnector
rhoas connector delete --id=myconnector

[connector.delete.flag.id.description]
one = 'The ID of the Connectors instance to delete'

[connector.delete.info.success]
one = 'Successfully deleted the Connectors instance'

[connector.delete.confirmDialog.message]
one = 'Are you sure that you want to delete the Connectors instance with ID "<no value>"?'

```

### Options

```
      --id string       The ID for the Connectors instance
      --name string     The name for the Connectors instance
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
  -y, --yes             Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

