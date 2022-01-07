## rhoas cluster clean

Remove all resources created by cluster extensions

### Synopsis

Remove all resources created by cluster extensions. This command removes all secrets and application service connections in your Kubernetes or OpenShift namespace.

```
rhoas cluster clean [flags]
```

### Examples

```
## Remove all resources
$ rhoas cluster clean

## Remove all resources in custom namespace
$ rhoas cluster clean --namespace=myspace

```

### Options

```
      --kubeconfig string   Location of the kubeconfig file
  -n, --namespace string    Use a custom Kubernetes namespace (if not set, the current namespace will be used)
  -y, --yes                 Forcibly perform operation without confirmation
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas cluster](rhoas_cluster.md)	 - View and perform operations on your Kubernetes or OpenShift cluster

