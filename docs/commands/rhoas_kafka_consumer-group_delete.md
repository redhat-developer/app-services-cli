## rhoas kafka consumer-group delete

Delete a consumer group

### Synopsis

Delete a consumer group from the current Kafka instance.

To select a different Kafka instance, use the “rhoas kafka use” command.


```
rhoas kafka consumer-group delete [flags]
```

### Examples

```
# delete a consumer group
$ rhoas kafka consumer-group delete --id consumer_group_1

```

### Options

```
      --id string   The unique ID of the consumer group to delete
  -y, --yes         Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka consumer-group](rhoas_kafka_consumer-group.md)	 - Describe, list, and delete consumer groups for the current Kafka instance

