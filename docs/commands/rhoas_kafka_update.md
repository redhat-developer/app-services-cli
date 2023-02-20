## rhoas kafka update

Update configuration details for a Kafka instance.

### Synopsis

Update configuration details for a Kafka instance. By modifying these settings, you can configure your Kafka instances to suit your particular environment.


```
rhoas kafka update [flags]
```

### Examples

```
# Update the Kafka instance owner
$ rhoas kafka update --name=my-kafka --owner=other-user

# Update the owner of the current Kafka instance
$ rhoas kafka update --owner=other-user

# Update the reauthentication configuration of the current Kafka instance
$ rhoas kafka update --reauthentication=true

# Update the current Kafka instance in interactive mode
$ rhoas kafka update

```

### Options

```
      --id string                  Unique ID of the Kafka instance you want to update
      --name string                Name of the Kafka instance you want to update
      --owner string               ID of the user you want to set as the owner of this Kafka instance
      --reauthentication Tribool   Enable or disable connection reauthentication for the Kafka instance
  -y, --yes                        Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

