## rhoas dedicated register-cluster

Register an OpenShift cluster with Red Hat OpenShift Streams for Apache Kafka

### Synopsis

You can use your own OpenShift cluster to provision your kafkas which will be managed by Red Hat Streams for Apache Kafka.
This command will register your cluster with Red Hat Streams for Apache Kafka.


```
rhoas dedicated register-cluster [flags]
```

### Examples

```
# Register an OpenShift cluster with Red Hat Streams for Apache Kafka by selecting from a list of available clusters.
rhoas cluster register-cluster

# Register an OpenShift cluster with Red Hat Streams for Apache Kafka by specifying the cluster ID.
rhoas cluster register-cluster --cluster-id 1234-5678-90ab-cdef

```

### Options

```
      --cluster-id string   The ID of the OpenShift cluster to register:
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas dedicated](rhoas_dedicated.md)	 - shortDescription

