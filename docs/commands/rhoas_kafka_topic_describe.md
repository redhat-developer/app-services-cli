## rhoas kafka topic describe

Describe a topic

### Synopsis

View configuration details for a Kafka topic.


```
rhoas kafka topic describe [flags]
```

### Examples

```
# Describe a topic
$ rhoas kafka topic describe --name topic-1

```

### Options

```
      --instance-id string   Kafka instance ID. Uses the current instance if not set 
      --name string          Format in which to display the Kafka topic (choose from: "json", "yml", "yaml")
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

