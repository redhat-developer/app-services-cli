## rhoas kafka topic create

Create a topic

### Synopsis

Create a topic in the current Kafka instance. You can specify the cleanup policy, number of partitions, retention size, and retention time.

The replicas are preconfigured. The number of partition replicas for the topic is set to 3 and the minimum number of follower replicas that must be in sync with a partition leader is set to 2.


```
rhoas kafka topic create [flags]
```

### Examples

```
# Create a topic
$ rhoas kafka topic create --name topic-1

```

### Options

```
      --cleanup-policy string   Determines whether log messages are deleted, compacted, or both (default "delete")
      --name string             Topic name
  -o, --output string           Specify the output format. Choose from: "json", "yaml", "yml"
      --partitions int32        The number of partitions in the topic (default 1)
      --retention-bytes int     The maximum total size of a partition log segments before old log segments are deleted to free up space.
                                Value of -1 is set by default indicating no retention size limits (default -1)
      --retention-ms int        The period of time in milliseconds the broker will retain a partition log before deleting it (default 604800000)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

