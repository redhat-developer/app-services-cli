## rhoas connector namespace list

Get a list of Connectors namespaces

### Synopsis

Get a list of Connectors namespaces for the Connectors cluster. The "connector namespace list" command returns details about the namespaces including their ID values.


```
rhoas connector namespace list [flags]
```

### Examples

```
# Get a list of Connectors namespaces
rhoas connector namespace list

```

### Options

```
      --limit int       Page limit (default 100)
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
      --page int        Page number (default 1)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connectors namespace commands

