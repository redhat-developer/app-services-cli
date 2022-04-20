## rhoas kafka list

List all Kafka instances

### Synopsis

List all Kafka instances.

By default, this command lists the Kafka instances in a table, showing the ID, name, owner, status, cloud provider, and region. You can also view the instances in JSON or YAML format.

To view additional details for a particular Kafka instance, use the “rhoas kafka describe” command.


```
rhoas kafka list [flags]
```

### Examples

```
# List all Kafka instances using the default output format
$ rhoas kafka list

# List all Kafka instances in JSON format
$ rhoas kafka list -o json

```

### Options

```
      --limit int       The maximum number of Kafka instances to be returned (default 100)
  -o, --output string   Specify the output format. Choose from: "json", "none", "table", "yaml", "yml" (default "table")
      --page int        Display the Kafka instances from the specified page number (default 1)
      --search string   Text search to filter the Kafka instances by name, owner, cloud_provider, region and status
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

