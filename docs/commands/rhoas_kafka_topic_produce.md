## rhoas kafka topic produce

Produce to a topic

### Synopsis

Produce to a topic in the current Kafka instance. You can specify the partition, key, and value.


```
rhoas kafka topic produce [flags]
```

### Examples

```
# Produce to a topic
$ rhoas kafka topic produce --name topic-1 --value "Hello world"

```

### Options

```
      --file string          Path to file containing value
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --key string           The key associated with the value produced. Empty if not set
      --name string          Topic name
      --partition int32      The partition receiving the message
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

