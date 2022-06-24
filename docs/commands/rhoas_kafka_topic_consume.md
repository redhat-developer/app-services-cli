## rhoas kafka topic consume

Consume messages from a topic

### Synopsis

Consume messages from a given topic, by default all messages on the topic will be consumed and printed in the foramat chosen. You can add filters 
to these message like a starting offset or a time that the messages must of been produced by.

Adding the --wait flag will wait for messages to be produced starting from when the command was ran and will ignore any limit or offset given.


```
rhoas kafka topic consume [flags]
```

### Examples

```
# Consume from a topic
$ rhoas kafka topic consume --name=topic-1

# Consume from a topic and wait for messages produced since command was ran
$ rhoas kafka topic consume --name=topic-1 --wait

```

### Options

```
      --format string        Format of the messages printed as they are produced, possible values are json and yaml (default "key-value")
      --from string          Only messages with a timestamp after this time will be consumed, time format required is YYYY-MM-DDThh:mm:ss.ssssZ
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --limit int32          Max records to consume from topic (default 20)
      --name string          Topic name
      --offset string        Retrieve messages within an offset equal to or greater than this
      --partition int32      The partition number used for consumer. Positive integer
      --wait                 Waiting for records to consume
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

