## rhoas connector delete

Delete a Connectors instance

### Synopsis

Delete a Connectors instance by specifing its ID. Use the "connector list" command to see a list of all Connectors instances and their ID values.

```
rhoas connector delete [flags]
```

### Examples

```
# Delete a Connectors instance with ID c9b71ucotd37bufoamkg
rhoas connector delete --id=c9b71ucotd37bufoamkg

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

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands
