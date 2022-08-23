## rhoas connector cluster list

List the Connectors clusters

### Synopsis

List all Connectors clusters in the OpenShift Dedicated environment. The returned list includes the ID value for each Connectors cluster.


```
rhoas connector cluster list [flags]
```

### Examples

```
# List Connectors clusters
rhoas connector cluster list

```

### Options

```
      --limit int       Page limit (default 10)
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
      --page int        Page number (default 1)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters

