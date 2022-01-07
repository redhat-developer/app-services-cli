## rhoas kafka consumer-group reset-offset

Reset partition offsets for a consumer group

### Synopsis

Reset partition offsets for a particular topic. A reset changes the offset position from which consumers read from the message log of a topic partition. To reset an offset position, the consumer group must have NO MEMBERS connected to a topic.

You can choose from the following options for the reset:
- Earliest (earliest offset at the start of the message log)
- Latest (latest offset at the end of the message log)
- Absolute (specific offset in the message log)
- Timestamp (specific timestamp in the message log)

You can also reset the offset position for all topics or a single, specified topic.

Warning: By resetting the offset position, you risk clients skipping or duplicating messages.


```
rhoas kafka consumer-group reset-offset [flags]
```

### Examples

```
# Reset partition offsets for a consumer group to latest
$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset latest

# Reset partition offsets for a consumer group to earliest
$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset earliest

# Reset partition offsets for a consumer group to an absolute value
$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset absolute --value 0

# Reset partition offsets for a consumer group to a timestamp
$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset timestamp --value "2016-06-23T09:07:21-07:00"

# Reset specific partition offsets for a consumer group
$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset latest --partitions 0,1

```

### Options

```
      --id string               The unique ID of the consumer group to reset-offset
      --offset string           Offset type (choose from: "earliest", "latest", "absolute", "timestamp")
      --partitions int32Slice   Reset consumer group offsets on specified partitions (comma-separated integers) (default [])
      --topic string            Reset consumer group offsets on a specified topic
      --value string            Custom offset value (required when offset is "absolute" or "timestamp")
  -y, --yes                     Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka consumer-group](rhoas_kafka_consumer-group.md)	 - Describe, list, and delete consumer groups for the current Kafka instance

