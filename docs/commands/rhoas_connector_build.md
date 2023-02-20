## rhoas connector build

Build a configuration file based on a connector type

### Synopsis

Build a configuration file by specifying a connector type.

For a list of available connector types, use the "connector type list" command.

The "connector build" command prompts you to provide values for the connector configuration properties.

After you build a configuration file, you can optionally edit it in a text editor.

You can then create a Connectors instance by using the "connector create" command and providing the name of the configuration file that you built.


```
rhoas connector build [flags]
```

### Examples

```
# Build a Connectors configuration file based on the "aws_lambda_sink_0.1" connector type. The default configuration file name is "connector.json"
rhoas connector build --type=aws_lambda_sink_0.1

# Build a Connectors configuration file named "my_aws_lambda_connector.json" that is based on the "aws_lambda_sink_0.1" connector type. Use "my-aws-lamda-sink.json" for the configuration file name.
rhoas connector build --name=my_aws_lambda_connector --type=aws_lambda_sink_0.1 --output-file=my-aws-lamda-sink.json

```

### Options

```
      --name string          The name of the connector type that was used to build a configuration file
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
      --output-file string   The file name of the connector configuration file
      --overwrite            Overwrite the file if it aready exists
      --type string          The type of the connector in the catalog - this value is the same as the ID value for the connector in the catalog
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

