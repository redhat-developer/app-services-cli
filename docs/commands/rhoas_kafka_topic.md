## rhoas kafka topic

Create, describe, update, list, and delete topics

### Synopsis

Create, describe, update, list and delete topics for the current Kafka instance.

Commands are executed on the currently selected Kafka instance.


### Examples

```
# Create a topic
rhoas kafka topic create --name mytopic

# List all topics
rhoas kafka topic list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances
* [rhoas kafka topic create](rhoas_kafka_topic_create.md)	 - Create a topic
* [rhoas kafka topic delete](rhoas_kafka_topic_delete.md)	 - Delete a topic
* [rhoas kafka topic describe](rhoas_kafka_topic_describe.md)	 - Describe a topic
* [rhoas kafka topic list](rhoas_kafka_topic_list.md)	 - List all topics
* [rhoas kafka topic update](rhoas_kafka_topic_update.md)	 - Update configuration details for a Kafka topic

