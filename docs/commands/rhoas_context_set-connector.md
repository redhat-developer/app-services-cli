## rhoas context set-connector

Set the current Connectors instance

### Synopsis

Set a Connectors instance as the current instance. The rhoas CLI uses the 
current Connectors instance when you run any "rhoas connector cluster" commands.

You can set a Connectors instance as the current instance by providing its name or ID.


```
rhoas context set-connector [flags]
```

### Examples

```
# Set the current Connectors instance by providing the name of a Connectors instance
$ rhoas context set-connector --name=my-connector

# Set the current Connectors instance by providing the ID of a Connectors instance
$ rhoas context set-connector --id=1iSSK8RQ3JKI8Q0OTFHF5FRg

```

### Options

```
      --id string     The unique ID of the Connectors instance that you want to set as the current instance
      --name string   The name of the Connectors instance that you want to set as the current instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

