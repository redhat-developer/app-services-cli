## rhoas context delete

Permanently delete a service context.

### Synopsis

Delete a service context.

```
rhoas context delete [flags]
```

### Examples

```
# Delete the currently-selected service context
$ rhoas context delete

# Delete a service context by name
$ rhoas context delete --name dev

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

