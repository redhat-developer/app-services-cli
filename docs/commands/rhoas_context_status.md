## rhoas context status

View the status of application services in a service context

### Synopsis

View the status of your application services. This command shows the status of each of the application services instances in the service context.

To view the status of a specific application service, use "rhoas context status [service]".

Note: You can change the current instance for an application service with the "rhoas [service] use” command.


```
rhoas context status [args] [flags]
```

### Examples

```
# View the status of all application services in the current service context
$ rhoas context status

# View the status of all application services in a specific service context
$ rhoas context status --name my-context

# View the status of the Kafka instance set in the current service context
$ rhoas context status kafka

# View the status of your services in JSON format
$ rhoas context status -o json

```

### Options

```
      --name string     Name of the context
  -o, --output string   Specify the output format. Choose from: "json", "none", "yaml", "yml" (default "json")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

