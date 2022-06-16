## rhoas kafka topic produce

Produce a new message to a topic

### Synopsis

Produce a message to a topic in a Kafka instance. Pass a file path to read that file as the message value or use stdin as your message. You can specify the partition, key, timestamp and value.


```
rhoas kafka topic produce [flags]
```

### Examples

```
# Produce to a topic
$ rhoas kafka topic produce --name=topic-1 --file="./value.json" --key=my-value

```

### Options

```
      --file string          Path to file containing message sent
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --key string           The message key. Empty if not set
      --name string          Topic name
      --partition int32      The partition number for the message. Must be positive integer value that represents number of partitions for the specified topic
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

