## rhoas connector start

Start a Connectors instance

### Synopsis

Start the current Connectors instance or start a Connectors instance by specifying its ID

```
rhoas connector start [flags]
```

### Examples

```
# Start the current Connectors instance
rhoas connector start

# Start a Connectors instance by specifying its ID
rhoas connector start --id=IJD76DUH675234

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

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

