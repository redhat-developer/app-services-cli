## rhoas connector namespace use

Set the current namespace in context

### Synopsis

Set a namespace as the current working namespace in context. The rhoas CLI uses the
current namespace you run any rhoas connector cluster commands.

You can set a namespace in the current context providing its name or ID.


```
rhoas connector namespace use [flags]
```

### Examples

```
# Set the current namespace by providing the name
$ rhoas connector namespace use --name=my-namespace

# Set the current namespace by providing the ID of a Connectors namespace
$ rhoas connector namespace use --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string     The unique ID of the namespace you want to set as the current namespace
      --name string   The name of the namespace you want to set as the namespace
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector namespace](rhoas_connector_namespace.md)	 - Connectors namespace commands

