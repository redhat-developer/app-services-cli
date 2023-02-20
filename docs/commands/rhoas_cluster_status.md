## rhoas cluster status

View the status of the current Kubernetes or OpenShift cluster

### Synopsis

View information about the current Kubernetes or OpenShift cluster. You can use this information to connect your application service to the cluster.

Before using this command, you must be logged in to a Kubernetes or OpenShift cluster. The command uses your kubeconfig file to identify the cluster context.


```
rhoas cluster status [flags]
```

### Examples

```
# print status of the current cluster
$ rhoas cluster status

```

### Options

```
      --kubeconfig string   Location of the kubeconfig file
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas cluster](rhoas_cluster.md)	 - View and perform operations on your Kubernetes or OpenShift cluster

