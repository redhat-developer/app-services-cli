## rhoas connector type describe

Get the details of a connector type

### Synopsis

Get the full list of details of a connector type by using its type ID

```
rhoas connector type describe [flags]
```

### Examples

```
# Describe connector type with ID of "slack_source_0.1"
rhoas connector type describe --type=slack_source_0.1 

# Describe the connector type with ID of "slack_source_0.1" and use YAML for the output format
rhoas connector type describe --type=slack_source_0.1 -o yaml

```

### Options

```
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
      --type string     The ID of the connector type that you want to get details about
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector type](rhoas_connector_type.md)	 - List and get details of the different connector types

