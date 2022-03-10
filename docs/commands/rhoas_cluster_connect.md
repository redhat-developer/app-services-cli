## rhoas cluster connect

Connect your services to Kubernetes or OpenShift

### Synopsis

Connect your application services to an Kubernetes or OpenShift cluster.

Note: Before you can connect your application service to OpenShift, you need to install the RHOAS Operator in your OpenShift cluster. For installation instructions, see https://github.com/redhat-developer/app-services-guides/tree/main/docs/kafka/service-binding-kafka#installing-the-rhoas-operator-on-openshift.

When you connect an application service to a Kubernetes or OpenShift cluster, this command uses the kubeconfig file (or the KUBECONFIG environment variable) to connect to the cluster and identify the context.

After identifying the context, the RHOAS Operator creates a service account and mounts it as a secret in your cluster. If your cluster has a service account already, the service account is refreshed.

Finally, the RHOAS Operator creates a Kafka Request Object, which the Service Binding Operator (https://github.com/redhat-developer/service-binding-operator) uses to create a "ServiceBinding" object.

After running this command, you need to grant access for the service account that was created by the RHOAS Operator to access your application service. For the Kafka application service, enter the following command:

  $ rhoas kafka acl grant-access --producer --consumer --service-account your-sa --topic all --group all

For the Service Registry application service, enter this command:

  $ rhoas service-registry role add --role=manager --service-account your-sa


```
rhoas cluster connect [flags]
```

### Examples

```
# Connect a Kafka instance to your cluster
$ rhoas cluster connect --service-type kafka --service-name kafka

# Connect a Service Registry instance to your cluster
$ rhoas cluster connect --service-type service-registry --service-name registry

```

### Options

```
      --ignore-context        Ignore currently-selected services and ask to select each service separately
      --kubeconfig string     Location of the kubeconfig file
  -n, --namespace string      Use a custom Kubernetes namespace (if not set, the current namespace will be used)
      --service-name string   Name of the application service to connect to
      --service-type string   Type of custom resource connection
      --token string          Provide an offline token to be used by the Operator (to get a token, visit https://console.redhat.com/openshift/token)
                              
  -y, --yes                   Forcibly perform operation without confirmation
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas cluster](rhoas_cluster.md)	 - View and perform operations on your Kubernetes or OpenShift cluster

