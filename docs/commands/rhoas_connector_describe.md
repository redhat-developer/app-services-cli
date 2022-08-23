## rhoas connector describe

Get the details for the Connectors instance

### Synopsis

Get the details for the Connectors instance by specifying its ID. Use the "connector list" command to see a list of all Connectors instances and their ID values.

```
rhoas connector describe [flags]
```

### Examples

```
#Get the Connectors instance details
rhoas connector describe --id=c980124otd37bufiemj0

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

