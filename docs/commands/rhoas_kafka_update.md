## rhoas kafka update

Update configuration details of a Kafka instance.

### Synopsis

Update certain configuration details of a Kafka instance.

Currently it is possible to update the "owner" field. The new owner 
will be authorized to manage this instance.


```
rhoas kafka update [flags]
```

### Examples

```
# update the Kafka instance owner
$ rhoas kafka update --name=my-kafka --owner=other-user

# update the owner of the current Kafka instance
$ rhoas kafka update --owner=other-user

# update the reauthentication configuration of the current Kafka instance
$ rhoas kafka update --reauthentication=true

# update the current Kafka instance in interactive mode
$ rhoas kafka update

```

### Options

```
      --id string                  Unique ID of the Kafka instance you want to update
      --name string                Name of the Kafka instance you want to update
      --owner string               ID of the user you want to set as the owner of this Kafka instance
      --reauthentication Tribool   Enable connection reauthentication for the Kafka instance
  -y, --yes                        Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

