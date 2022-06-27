## rhoas connector namespace create

Create a connector namespace

### Synopsis

Create a connector namespace

```
rhoas connector namespace create [flags]
```

### Examples

```
# Create a connector namespace
rhoas connector namespace create --name "my-namespace"

# Create an evaluation namespace
rhoas connector namespace create --name "my-namespace" --eval

```

### Options

```
      --eval            Create an evaluation namespace
      --name string     Name of the namespace
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connector namespace commands

