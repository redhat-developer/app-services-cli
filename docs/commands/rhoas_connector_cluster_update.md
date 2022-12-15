## rhoas connector cluster update

Update a Connectors cluster

### Synopsis

Update a Connectors cluster

```
rhoas connector cluster update [flags]
```

### Examples

```
# Update name of a connector cluster
rhoas connector cluster update --id cdh0s0bjdpqd9bgomcbg --name my-connector

# Update annotations of a connector cluster
rhoas connector cluster update --id cdh0s0bjdpqd9bgomcbg --annotations h1=head

```

### Options

```
      --annotations key=value   comma-separated list of string annotations in key=value format
      --id string               ID of the Connectors cluster to update
      --name string             Name of the Connectors cluster
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters

