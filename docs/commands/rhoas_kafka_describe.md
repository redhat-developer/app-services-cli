## rhoas kafka describe

View configuration details of a Kafka instance

### Synopsis

View configuration details for a Kafka instance.

Use the "--id" flag to specify which instance you would like to view.

If the "--id" flag is not used then the selected Kafka instance will be used, if available.

You can view the output either as JSON or YAML.

To view a list of all Kafka instances, use the “rhoas kafka list” command.


```
rhoas kafka describe [flags]
```

### Examples

```
# View the current Kafka instance
$ rhoas kafka describe

# View a specific instance by ID
$ rhoas kafka describe --id=1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

# View a specific instance by name
$ rhoas kafka describe --name=my-kafka

# Customize the output format
$ rhoas kafka describe -o yaml

```

### Options

```
      --bootstrap-server   If specified, only the bootstrap server host of the Kafka instance will be displayed
      --id string          Unique ID of the Kafka instance you want to view
      --name string        Name of the Kafka instance you want to view
  -o, --output string      Specify the output format. Choose from: "json", "none", "yaml", "yml" (default "json")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

