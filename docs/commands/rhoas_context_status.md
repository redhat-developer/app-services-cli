## rhoas context status

View the status of your application services

### Synopsis

View the status of your application services. This command shows the status of the current instance for each of your application services.

To view the status of a specific application service, use "rhoas status [service]".

Note: You can change the current instance for an application service with the "rhoas [service] use” command.


```
rhoas context status [args] [flags]
```

### Examples

```
# View the status of all application services
$ rhoas status

# View the status of the current Kafka instance
$ rhoas status kafka

# View the status of your services in JSON format
$ rhoas status -o json

```

### Options

```
      --name string     Name of the context
  -o, --output string   Format in which to display the status of your services (choose from: "json", "yml", "yaml")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

