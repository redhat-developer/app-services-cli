## rhoas kafka topic update

Update configuration details for a Kafka topic

### Synopsis

Update a topic in the current Kafka instance. You can update the cleanup policy, number of partitions, retention size, and retention time.


```
rhoas kafka topic update [flags]
```

### Examples

```
# Update the message retention period for a topic
$ rhoas kafka topic update --name topic-1 --retention-ms -1

```

### Options

```
      --cleanup-policy string    Determines whether log messages are deleted, compacted, or both
      --name string              Topic name
      --partitions string        The number of partitions in the topic
      --retention-bytes string   The maximum total size of a partition log segments before old log segments are deleted to free up space
      --retention-ms string      The period of time in milliseconds the broker will retain a partition log before deleting it
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

