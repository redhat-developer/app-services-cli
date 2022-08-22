## rhoas connector use

Set the current Connectors instance

### Synopsis

Set a Connectors instance as the current instance. The rhoas CLI uses the 
current Connectors instance when you run any rhoas connector cluster commands.

You can set a Connectors instance as the current instance by providing its name or ID.


```
rhoas connector use [flags]
```

### Examples

```
# Set the current Connectors instance by providing the name of a Connectors instance
$ rhoas connector use --name=my-connector

# Set the current Connectors instance by providing the ID of a Connectors instance
$ rhoas connector use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

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

* [rhoas connector](rhoas_connector.md)	 - Connectors commands

