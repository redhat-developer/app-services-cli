## rhoas context set-namespace

Set the current namespace in context

### Synopsis

Set a namespace as the current working namespace in context. The rhoas CLI uses the
current namespace you run any rhoas connector cluster commands.

You can set a namespace in the current context providing its name or ID.


```
rhoas context set-namespace [flags]
```

### Examples

```
# Set the current namespace providing the name of a Connectors namespace
$ rhoas context set-namespace --name=my-namespace

# Set the current namespace providing the ID of a Connectors namespace
$ rhoas context set-namespace --id=1iSSK8RQ3JKI8Q0OTFHF5FRg

```

### Options

```
      --id string     The unique ID of the namespace you want to set as the current namespace
      --name string   The name of the namespace you want to set as the current namespace
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

