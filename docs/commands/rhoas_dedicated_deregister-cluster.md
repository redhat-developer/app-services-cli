## rhoas dedicated deregister-cluster

Deregister an OpenShift cluster with Red Hat OpenShift Streams for Apache Kafka

### Synopsis

Removes the ability to provision your own Kafka instances on a cluster, this command will deregister your 
Hybrid cluster with Red Hat OpenShift Streams for Apache Kafka.


```
rhoas dedicated deregister-cluster [flags]
```

### Examples

```
# Deregister an OpenShift cluster with Red Hat Streams for Apache Kafka by selecting from a list of available clusters.
rhoas cluster deregister-cluster

# Deregister an OpenShift cluster with Red Hat Streams for Apache Kafka by specifying the cluster ID.
rhoas cluster deregister-cluster --cluster-id 1234-5678-90ab-cdef

```

### Options

```
      --access-token string           The access token to use to authenticate with the OpenShift Cluster Management API.
      --cluster-id string             The ID of the OpenShift cluster to deregister.
      --cluster-mgmt-api-url string   The API URL of the OpenShift Cluster Management API.
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas dedicated](rhoas_dedicated.md)	 - Manage your Hybrid OpenShift clusters which host your Kafka instances.

