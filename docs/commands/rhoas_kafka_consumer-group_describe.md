## rhoas kafka consumer-group describe

Describe a consumer group

### Synopsis

View detailed information for a consumer group and its members.


```
rhoas kafka consumer-group describe [flags]
```

### Examples

```
# describe a consumer group
$ rhoas kafka consumer-group describe --id consumer_group_1 -o json

```

### Options

```
      --id string       The unique ID of the consumer group to view
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml" or "none"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka consumer-group](rhoas_kafka_consumer-group.md)	 - Describe, list, and delete consumer groups for the current Kafka instance

