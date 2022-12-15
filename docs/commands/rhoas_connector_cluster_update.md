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
      --annotations strings   connector.cluster.update.flag.annotations.description
      --id string             ID of the Connectors cluster to update
      --name string           connector.cluster.update.flag.name.description
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector cluster](rhoas_connector_cluster.md)	 - Create, delete, and list Connectors clusters

