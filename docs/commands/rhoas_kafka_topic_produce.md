## rhoas kafka topic produce

Produce a new message to a topic

### Synopsis

Produce a message to a topic in a Kafka instance. Specify a file path for the CLI to read as the message value or use stdin to provide your message. You can specify the partition, key, and value.


```
rhoas kafka topic produce [flags]
```

### Examples

```
# Produce a single message to a topic from a file and provide a custom message key
$ rhoas kafka topic produce --name=users --file="./message.json" --key="{'location': 'us-east-1'}"

# Produce to a topic from standard input
$ rhoas kafka topic produce --name=users

# Produce to a topic from other command output
$ cat yourfile.json | rhoas kafka topic produce --name=users

# Produce to a topic and fetch its offset value
$ rhoas kafka topic produce --name=topic-1 --file="./message.json" | jq .offset

# Produce to a topic from a JSON file and filter using jq
$ cat input.json | jq .data.value | rhoas kafka topic produce --name=topic-1

# Produce to a topic and specified partition
$ rhoas kafka topic produce --name=topic-1 --file="./message.json" --partition=1

```

### Options

```
      --file string          Path to file containing message value
      --format string        Format for printing produced messages (possible values are "json" and "yaml") (default "json")
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --key string           Message key. Empty if not set
      --name string          Topic name
      --partition int32      Partition number for the message. Must be a positive integer value that is valid for the specified topic
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

