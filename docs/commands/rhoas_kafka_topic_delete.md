## rhoas kafka topic delete

Delete a topic

### Synopsis

Delete a topic in the current Kafka instance.


```
rhoas kafka topic delete [flags]
```

### Examples

```
# Delete a topic
$ rhoas kafka topic delete --name topic-1

```

### Options

```
      --name string   Topic name
  -y, --yes           Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

