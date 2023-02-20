## rhoas connector cluster create

Create a Connectors cluster

### Synopsis

Create a Connectors cluster and specify its name. You must have administrator access to run this command.

```
rhoas connector cluster create [flags]
```

### Examples

```
# Create a Connectors cluster that is named "my-connectors-cluster"
rhoas connector cluster create --name=my-connectors-cluster

```

### Options

```
      --name string     Name of the Connectors cluster to create
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters

