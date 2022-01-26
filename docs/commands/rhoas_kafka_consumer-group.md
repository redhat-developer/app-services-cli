## rhoas kafka consumer-group

Describe, list, and delete consumer groups for the current Kafka instance

### Synopsis

View and delete consumer groups for the current Kafka instance.

These commands operate on the current Kafka instance. To select the Kafka instance, use the “rhoas kafka use” command.


### Examples

```
# Delete a consumer group
rhoas kafka consumer-group delete --id consumer_group_1

# List all consumer groups
rhoas kafka consumer-group list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances
* [rhoas kafka consumer-group delete](rhoas_kafka_consumer-group_delete.md)	 - Delete a consumer group
* [rhoas kafka consumer-group describe](rhoas_kafka_consumer-group_describe.md)	 - Describe a consumer group
* [rhoas kafka consumer-group list](rhoas_kafka_consumer-group_list.md)	 - List all consumer groups
* [rhoas kafka consumer-group reset-offset](rhoas_kafka_consumer-group_reset-offset.md)	 - Reset partition offsets for a consumer group

