## rhoas

RHOAS CLI

### Synopsis

Red Hat OpenShift Application Services

Manage your application services from the command line. You can manage service accounts, Kafka instances, and Service Registry instances, and connect them to your OpenShift clusters and applications.


### Examples

```
# Authenticate securely through your web browser
$ rhoas login

# Create a Kafka instance
$ rhoas kafka create --name my-kafka-instance

# Create a service account and save credentials to a JSON file
$ rhoas service-account create -o json

# Connect your Kubernetes/OpenShift cluster to a service
$ rhoas cluster connect

```

### Options

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
      --version   Show rhoas version
```

### SEE ALSO

* [rhoas authtoken](rhoas_authtoken.md)	 - Output the current token
* [rhoas cluster](rhoas_cluster.md)	 - View and perform operations on your Kubernetes or OpenShift cluster
* [rhoas completion](rhoas_completion.md)	 - Install command completion for your shell (bash, zsh, fish or powershell)
* [rhoas connector](rhoas_connector.md)	 - Connectors commands
* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services
* [rhoas dedicated](rhoas_dedicated.md)	 - Manage your Hybrid OpenShift clusters which host your Kafka instances.
* [rhoas generate-config](rhoas_generate-config.md)	 - Generate configurations for the service context
* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances
* [rhoas login](rhoas_login.md)	 - Log in to RHOAS
* [rhoas logout](rhoas_logout.md)	 - Log out from RHOAS
* [rhoas request](rhoas_request.md)	 - Allows users to perform API requests against the API server
* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts
* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas status](rhoas_status.md)	 - View the status of application services in a service context
* [rhoas whoami](rhoas_whoami.md)	 - Output the current username

