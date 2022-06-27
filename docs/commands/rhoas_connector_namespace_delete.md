## rhoas connector namespace delete

Delete a Connectors namespace

### Synopsis

Delete a Connectors namespace by specifing its ID. Use the "connector namespace list" command to see a list of all Connectors namespaces and their ID values.

```
rhoas connector namespace delete [flags]
```

### Examples

```
# Delete the Connectors namespace with ID jdhdhdhmmf
rhoas connector namespace delete --id jdhdhdhmmf

```

### Options

```
      --id string   The ID for the Connectors instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connectors namespace commands

