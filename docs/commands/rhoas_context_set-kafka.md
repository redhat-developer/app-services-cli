## rhoas context set-kafka

Set the current Kafka instance

### Synopsis

Select a Kafka instance to be the current instance. When you set the Kafka instance to be used, it is set as the current instance for all “rhoas kafka topic” and “rhoas kafka consumer-group” commands.

You can select a  Kafka instance by name or ID.


```
rhoas context set-kafka [flags]
```

### Examples

```
# Select a Kafka instance by name to be set in the current context
$ rhoas context set-kafka --name=my-kafka

# Select a Kafka instance by ID to be set in the current context
$ rhoas context set-kafka --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string     Unique ID of the Kafka instance you want to set as the current instance
      --name string   Name of the Kafka instance you want to set as the current instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

