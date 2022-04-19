## rhoas kafka topic list

List all topics

### Synopsis

List all topics in the current Kafka instance.


```
rhoas kafka topic list [flags]
```

### Examples

```
# List all topics
$ rhoas kafka topic list

# List all topics in JSON format
$ rhoas kafka topic list -o json

```

### Options

```
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml" or "none"
      --page int32      Current page number for list of topics (default 1)
      --search string   Text search to filter the Kafka topics by name
      --size int32      Maximum number of items to be returned per page (default 10)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

