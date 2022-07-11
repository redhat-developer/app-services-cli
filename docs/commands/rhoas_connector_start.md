## rhoas connector start

Start a connector instance

### Synopsis

Start a Connectors instance, pass an id or use the instance in the current context

```
rhoas connector start [flags]
```

### Examples

```
# Start a connector instance
rhoas connector start

# Start a connector instance
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

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

