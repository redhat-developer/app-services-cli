## rhoas context set-service-registry

Use a Service Registry instance

### Synopsis

Select a Service Registry instance to use with all instance-specific commands.
You can specify a Service Registry instance by --name or --id.

When you set the Service Registry instance to be used, it is set as the current instance for all rhoas service-registry artifact commands.


```
rhoas context set-service-registry [flags]
```

### Examples

```
# Select a Service Registry instance by name to be set in the current context
rhoas context set-service-registry --name my-service-registry

# Select a Service Registry instance by ID to be set in the current context
rhoas context set-service-registry --id 1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string     Unique ID of the Service Registry instance you want to set as the current instance
      --name string   Name the Service Registry instance you want to set as the current instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

