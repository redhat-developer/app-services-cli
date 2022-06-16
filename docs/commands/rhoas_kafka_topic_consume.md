## rhoas kafka topic consume

consume short

### Synopsis

consume long

```
rhoas kafka topic consume [flags]
```

### Examples

```
consume example
```

### Options

```
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --limit int32          Max records to consume from topic (default 20)
      --name string          Topic name
      --offset int           Retrieve messages within an offset equal to or greater than this
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
      --partition int32      The partition number used for consumer. Positive integer
      --timestamp string     Timestamp to start consuming from
      --wait                 Waiting for records to consume
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

