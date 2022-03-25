## rhoas kafka consumer-group list

List all consumer groups

### Synopsis

List the consumer groups in the current Kafka instance. You can view a list of all consumer groups in the Kafka instance, view a specific consumer group, or view the consumer groups for a particular topic.

```
rhoas kafka consumer-group list [flags]
```

### Examples

```
# List all consumer groups
$ rhoas kafka consumer-group list

# List all consumer groups in JSON format
$ rhoas kafka consumer-group list -o json

```

### Options

```
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
  -o, --output string        Specify the output format. Choose from: "json", "none", "table", "yaml", "yml" (default "table")
      --page int32           View the specified page number in the list of consumer groups (default 1)
      --search string        Text search to filter consumer groups by ID
      --size int32           Maximum number of consumer groups to be returned per page (default 10)
      --topic string         Fetch the consumer groups for a specific Kafka topic
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka consumer-group](rhoas_kafka_consumer-group.md)	 - Describe, list, and delete consumer groups for the current Kafka instance

