## rhoas kafka topic consume

Consume messages from a topic

### Synopsis

Consume messages from a given topic. By default, all messages on the topic are consumed and printed in the format that you specify. You can spcecify filters
such as a starting offset or a starting time for message production.

If you add the --wait flag, the CLI waits for messages to be produced starting from when you run the command and ignores any limit or offset specified.


```
rhoas kafka topic consume [flags]
```

### Examples

```
# Consume from a topic
$ rhoas kafka topic consume --name=topic-1

# Consume from a topic and output to YAML format
$ rhoas kafka topic consume --name=topic-1 --format=yaml

# Consume from a topic by continually polling the topic for new messages
$ rhoas kafka topic consume --name=topic-1 --wait

# Consume from a topic starting from a time specified in default ISO time format
$ rhoas kafka topic consume --name=topic-1 --from-date=2022-06-17T07:05:34.0000Z

# Consume from a topic starting from a time specified in Unix epoch time format
$ rhoas kafka topic consume --name=topic-1 --wait --from-timestamp=1656346796

# Consume from a topic starting from an offset value
$ rhoas kafka topic consume --name=topic-1 --offset=15

# Consume from a topic starting from an offset value and with a specified message limit
$ rhoas kafka topic consume --name=topic-1 --offset=15 --limit=30

# Consume from topic and output to JSON format and using jq to read values
$ rhoas kafka topic consume --name=topic-1 --format=json | jq -rc .value

```

### Options

```
      --format string           Format for printing produced messages (possible values are "json" and "yaml") (default "key-value")
      --from-date string        Consume only messages with a date later than the specified value (requied format is YYYY-MM-DDThh:mm:ss.ssssZ)
      --from-timestamp string   Consume only messages with a timestamp later than the specified value (required format is Unix epoch timestamp value)
      --instance-id string      Kafka instance ID. Uses the current instance if not set 
      --limit int32             Maximum number of messages to consume from topic (default 20)
      --name string             Topic name
      --offset string           Consume messages from an offset equal to or greater than the specified value
      --partition string        Consume messages from specified partition (value must be a positive integer)
      --wait                    Wait for messages to consume from topic
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics

