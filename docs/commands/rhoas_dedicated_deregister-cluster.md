## rhoas dedicated deregister-cluster

Deregister a OpenShift cluster from use with Red Hat OpenShift Streams for Apache Kafka

### Synopsis

Removes the ability to provision your own Kafka instances on your OpenShift cluster, this command will deregister your
OpenShift cluster from use with Red Hat OpenShift Streams for Apache Kafka.


```
rhoas dedicated deregister-cluster [flags]
```

### Examples

```
# Deregister an OpenShift cluster from use with Red Hat Streams for Apache Kafka by selecting from a list of available clusters.
rhoas dedicated deregister-cluster

# Deregister an OpenShift cluster from Red Hat Streams for Apache Kafka by specifying the cluster ID
rhoas dedicated deregister-cluster --cluster-id 1234-5678-90ab-cdef

```

### Options

```
      --cluster-id string   The ID of the OpenShift cluster to deregister
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas dedicated](rhoas_dedicated.md)	 - Manage your OpenShift clusters which host your Kafka instances

