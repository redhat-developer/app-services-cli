## rhoas cluster

View and perform operations on your Kubernetes or OpenShift cluster

### Synopsis

Connect and bind your services to Kubernetes or OpenShift applications. You can also check if the Kubernetes or OpenShift cluster has the required Operators to perform this operation.

### Examples

```
# Check the status of the connection to your cluster
$ rhoas cluster status

# Connect to cluster without including currently selected services
$ rhoas cluster connect --ignore-context

# Connect to cluster using the specified token
$ rhoas cluster connect --token=value

# Connect to cluster and save script to create service binding
$ rhoas cluster connect --yes > create_service_binding.sh

# Connect managed service with your application
$ rhoas cluster bind

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas cluster bind](rhoas_cluster_bind.md)	 - Connect your RHOAS services to Kubernetes or OpenShift applications
* [rhoas cluster clean](rhoas_cluster_clean.md)	 - Remove all resources created by cluster extensions
* [rhoas cluster connect](rhoas_cluster_connect.md)	 - Connect your services to Kubernetes or OpenShift
* [rhoas cluster status](rhoas_cluster_status.md)	 - View the status of the current Kubernetes or OpenShift cluster

