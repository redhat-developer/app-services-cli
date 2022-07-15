## rhoas connector stop

Stop a Connectors instance

### Synopsis

Stop the current Connectors instance or stop a Connectors instance by providing its ID

```
rhoas connector stop [flags]
```

### Examples

```
# Stop current Connectors instance
rhoas connector stop

# Stop a Connectors instance by specifying its ID
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

