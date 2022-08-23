## rhoas connector create

Create a Connectors instance

### Synopsis

Create a Connectors instance. 

A Connectors instance is an instance of a one of the supported Connectors.
Use the "connector" command to create, delete, and view a list of Connectors instances.

Before you create a Connectors instance, you must use the "connector build" command to create a configuration file for the type of connector that you want to create. 


```
rhoas connector create [flags]
```

### Examples

```
# Create a Connectors instance by specifying a configuration file
rhoas connector create --file=myconnector.json

```

### Options

```
      --create-service-account   If set, the connector is created with the specified service account
  -f, --file string              Location of the Connectors JSON file that describes the connector
      --kafka string             ID of the Kafka instance (the default is the Kafka instance for the current context)
      --name string              Override the name of the Connectors instance (the default name is the name specified in the connector configuration file)
      --namespace string         ID of the namespace for the Connectors instance (the default is the namespace for the current context)
  -o, --output string            Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

