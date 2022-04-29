## rhoas context list

List service contexts

### Synopsis

List all service contexts. This command lists each service context, and indicates the context that is currently being used.

To view the details of a service context, use the "rhoas context status" command.


```
rhoas context list [flags]
```

### Examples

```
# List contexts
$ rhoas context list

```

### Options

```
  -o, --output string   Specify the output format. Choose from: "json", "none", "yaml", "yml" (default "json")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

