## rhoas dedicated register-cluster

Register an OpenShift cluster with Red Hat OpenShift Streams for Apache Kafka

### Synopsis

You can use your own OpenShift cluster to provision your Kafka instances which will be managed by Red Hat Streams for Apache Kafka
This command will register your cluster with Red Hat Streams for Apache Kafka


```
rhoas dedicated register-cluster [flags]
```

### Examples

```
# Register an OpenShift cluster with Red Hat Streams for Apache Kafka by selecting from a list of available clusters
rhoas dedicated register-cluster

# Register an OpenShift cluster with Red Hat Streams for Apache Kafka by specifying the cluster ID
rhoas dedicated register-cluster --cluster-id 1234-5678-90ab-cdef

```

### Options

```
      --access-token string           The access token to use to authenticate with the OpenShift Cluster Management API
      --cluster-id string             The ID of the OpenShift cluster to register:
      --cluster-mgmt-api-url string   The API URL of the OpenShift Cluster Management API
      --page-number int               The page number to use when listing Hybrid OpenShift clusters (default 1)
      --page-size int                 The page size to use when listing Hybrid OpenShift clusters (default 100)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas dedicated](rhoas_dedicated.md)	 - Manage your Hybrid OpenShift clusters which host your Kafka instances

