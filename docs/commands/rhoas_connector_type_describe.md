## rhoas connector type describe

Get details of a connector type

### Synopsis

Get full list of details of a connector type using its type id

```
rhoas connector type describe [flags]
```

### Examples

```
# Desribe connector type with id of slack_source_0.1
rhoas connector type describe --type=slack_source_0.1 

# Desribe connector type with id of slack_source_0.1 and give output as yaml
rhoas connector type describe --type=slack_source_0.1 -o yaml

```

### Options

```
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
      --type string     The type id of the connector you want to get details about
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector type](rhoas_connector_type.md)	 - List and get details of different connector types

