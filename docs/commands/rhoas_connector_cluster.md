## rhoas connector cluster

Create, delete, and list Connectors clusters

### Synopsis

A Connectors cluster is an OpenShift Dedicated instance for deploying your Connectors instances. Use the "connector cluster" command to create, delete, and view a list of Connectors clusters.


### Examples

```
# Create a Connectors cluster that is named "my-connectors-cluster"
rhoas connector cluster create --name=my-connectors-cluster

# Delete a Connectors cluster that has ID c980124otd37bufiemj0
rhoas connector cluster delete --id=c980124otd37bufiemj0

# List Connectors clusters
rhoas connector cluster list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands
* [rhoas connector cluster addon-parameters](rhoas_connector_cluster_addon-parameters.md)	 - Get Connectors add-on parameters
* [rhoas connector cluster create](rhoas_connector_cluster_create.md)	 - Create a Connectors cluster
* [rhoas connector cluster delete](rhoas_connector_cluster_delete.md)	 - Delete a Connectors cluster
* [rhoas connector cluster list](rhoas_connector_cluster_list.md)	 - List Connectors clusters

