## rhoas status

View the status of application services in a service context

### Synopsis

View the status of your application services. This command shows the status of each of your application services instances that belong to a service context.

To view the status of a specific application service, use "rhoas status [service]".

Note: You can change the current instance for an application service with the "rhoas [service] use” command.


```
rhoas status [args] [flags]
```

### Examples

```
# View the status of all application services in the current service context
$ rhoas status

# View the status of all application services in a specific service context
$ rhoas status kafka

# View the status of your services in JSON format
$ rhoas status -o json

```

### Options

```
      --name string     Name of the context
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

