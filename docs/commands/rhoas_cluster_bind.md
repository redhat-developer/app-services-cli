## rhoas cluster bind

Connect your RHOAS services to Kubernetes or OpenShift applications

### Synopsis

Connect an application service to your application running on OpenShift or Kubernetes. This command uses the Service Binding Operator to inject an application running on the cluster with the parameters required to connect to the application service.

When you perform service binding, the Service Binding Operator automatically injects connection parameters as files into the pod for the application (you can see these files in the “/bindings” directory in your application). This means that the application automatically detects and uses the injected connection parameters.

In general, this automatic injection and detection of connection parameters eliminates the need to manually configure an application to connect to an application service. This is a particular advantage if you have many applications in your project that you want to connect to the application service.

If your application uses a “DeploymentConfig”, you must use the “--deployment-config" flag so that your application can be detected.


```
rhoas cluster bind [flags]
```

### Examples

```
# Bind using interactive mode
$ rhoas cluster bind

# Bind to specific namespace and application
$ rhoas cluster bind --namespace=ns --app-name=myapp

```

### Options

```
      --app-name string       Name of the Kubernetes deployment to bind
      --bind-env              Bind service as environment variables
      --binding-name string   Name of the Service Binding object to create when using the Service Binding Operator
      --deployment-config     Use DeploymentConfig to detect your application (you must manually redeploy your application for the binding to take effect)
      --ignore-context        Ignore currently-selected services and ask to select each service separately
      --kubeconfig string     Location of the kubeconfig file
  -n, --namespace string      Use a custom Kubernetes namespace (if not set, the current namespace will be used)
      --service-name string   Name of the application service to connect to
      --service-type string   Type of custom resource connection
  -y, --yes                   Forcibly perform operation without confirmation
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas cluster](rhoas_cluster.md)	 - View and perform operations on your Kubernetes or OpenShift cluster

