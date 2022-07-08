## rhoas context set-connector

Set the current connector instance

### Synopsis

Select a connector instance to be the current instance. When you set the connector instance to be used, it is set as the current instance for all 
“rhoas connector cluster” commands.

You can select a  connector instance by name or ID.


```
rhoas context set-connector [flags]
```

### Examples

```
# Select a connector instance by name to be set in the current context
$ rhoas context set-connector --name=my-connector

# Select a connector instance by ID to be set in the current context
$ rhoas context set-connector --id=1iSSK8RQ3JKI8Q0OTFHF5FRg

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

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

