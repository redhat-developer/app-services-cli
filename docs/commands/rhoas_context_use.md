## rhoas context use

Set the current context

### Synopsis

Select a service context to be used as the current context.

When you set the context to be used, it is set as the current context for all service-based rhoas commands.


```
rhoas context use [flags]
```

### Examples

```
# Set the current context
$ rhoas context use --name dev

```

### Options

```
      --name string   Name of the context
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

