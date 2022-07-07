## rhoas connector use

Set the current connector instance

### Synopsis

Select a connector instance to be the current instance. When you set the connector instance to be used, it is set as the current instance for all 
“rhoas connector cluster” commands.

You can select a  connector instance by name or ID.


```
rhoas connector use [flags]
```

### Examples

```
# Select a connector instance by name to be set in the current context
$ rhoas connector use --name=my-connector

# Select a connector instance by ID to be set in the current context
$ rhoas connector use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string     Unique ID of the connector instance you want to set as the current instance
      --name string   Name of the connector instance you want to set as the current instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector](rhoas_connector.md)	 - Connectors instance commands

