## rhoas context unset

Unset services in context

### Synopsis

Unset services in context

When you unset a service in context, it will no longer point to any instance of that service.


```
rhoas context unset [flags]
```

### Examples

```
# Unset services for current context
$ rhoas context unset --services kafka,service-registry

# Unset service-registry for a specific context
$ rhoas context unset --name dev --services service-registry

```

### Options

```
      --name string        Name of the context
      --services strings   context.unset.flag.services.description
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

