## rhoas connector create

Create a Connectors instance

### Synopsis

Create a Connectors instance

```
rhoas connector create [flags]
```

### Examples

```
# Create a Connectors instance
rhoas connector create --file=myconnector.json

```

### Options

```
      --create-service-account   If set, the connector will be created with the newly specified service account
  -f, --file string              Location of the Connectors JSON file that describes the connector
      --kafka string             Id of the kafka instance (by default kafka instance from context would be used)
      --name string              Override name of the connector (by default name in the connector spec would be used)
      --namespace string         Id of the namespace for the connector (by default namespace from context would be used)
  -o, --output string            Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

