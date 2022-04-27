## rhoas context create

Create a service context

### Synopsis

Create a service context and assign associated service identifiers.

A service context is a group of application service instances and their configuration details. By creating a service context, you can group together application service instances that you want to use together.

After creating the service context, add application service instances to it by using the "rhoas context set-[service]" commands.


```
rhoas context create [flags]
```

### Examples

```
# Create context
$ rhoas context create --name dev

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

