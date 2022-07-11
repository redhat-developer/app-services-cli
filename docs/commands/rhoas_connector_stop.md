## rhoas connector stop

Stop a connector instance

### Synopsis

Stop a connector instance, pass an id or use the instance in the current context

```
rhoas connector stop [flags]
```

### Examples

```
# Stop a connector instance
rhoas connector stop

# Stop a connector instance
rhoas connector stop --id=IJD76DUH675234

```

### Options

```
      --id string       The ID for the Connectors instance
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

