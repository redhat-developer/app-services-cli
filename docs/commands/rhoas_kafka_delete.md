## rhoas kafka delete

Delete a Kafka instance

### Synopsis

Permanently delete a Kafka instance, including all topics.

When this command is run, you will be asked to confirm the name of the instance you want to delete. Otherwise you can use "--yes" to skip confirmation and forcibly delete the instance.


```
rhoas kafka delete [flags]
```

### Examples

```
# Delete the current Kafka instance
$ rhoas kafka delete

# Delete a Kafka instance with a specific ID
$ rhoas kafka delete --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

# Delete a Kafka instance with a specific name
$ rhoas kafka delete --name=my-kafka

```

### Options

```
      --id string     Unique ID of the Kafka instance you want to delete
      --name string   Name of the Kafka instance you want to delete
  -y, --yes           Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

