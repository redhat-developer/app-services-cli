## rhoas connector cluster delete

Delete a Connectors cluster

### Synopsis

Delete a Connectors cluster by specifying its cluster ID. Use the "connector cluster list" command to see a list of all Connectors clusters and their ID values.


```
rhoas connector cluster delete [flags]
```

### Examples

```
# Delete a Connectors cluster that has ID c980124otd37bufiemj0
rhoas connector cluster delete --id=c980124otd37bufiemj0

```

### Options

```
      --id string       ID of the Connectors cluster to delete
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
  -y, --yes             Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters

