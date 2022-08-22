## rhoas connector build

Build a Connectors instance

### Synopsis

Build a configuration file by specifying a connector type.
For a list of available connector types, use the "connector type list" command.

After you build a configuration file, you can create a Connectors instance by using the "connector create" command and providing the name of the configuration file that you built.


```
rhoas connector build [flags]
```

### Examples

```
# Build a Connectors configuration file based on the "aws_lambda_sink_0.1" connector type. The default configuration file name is "connector.json"
rhoas connector build --type=--type=aws_lambda_sink_0.1

# Build a Connectors configuration file named "my_aws_lambda_connector.json" that is based on the "aws_lambda_sink_0.1" connector type
rhoas connector build --name=my_aws_lambda_connector --type=--type=aws_lambda_sink_0.1

```

### Options

```
      --name string          The name of the connector type that was used to build a configuration file.
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
      --output-file string   The filename of the connector configuration file
      --overwrite            Overwrite the file if it aready exists.
      --type string          The type of the connector in the catalog (the same as the ID value for the connector in the catalog)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

