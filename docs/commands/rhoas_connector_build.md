## rhoas connector build

Build a Connectors instance

### Synopsis

Command builds specficication for Connector 
instance that can be later created using create command.


```
rhoas connector build [flags]
```

### Examples

```
# Build a Connectors instance
rhoas connector build --type=--type=aws_lambda_sink_0.1

# Build a Connectors instance with a name of my_connector
rhoas connector build --name=my_connector --type=--type=aws_lambda_sink_0.1

cat myconnector.json | rhoas connector create

```

### Options

```
      --name string          name of the connector
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
      --output-file string   filename of the connector specification file
      --overwrite            should overwrite file if exist
      --type string          Connector type (id of the connector)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

