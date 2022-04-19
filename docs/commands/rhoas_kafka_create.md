## rhoas kafka create

Create a Kafka instance

### Synopsis

Create a Kafka instance on a particular cloud provider and region.

After creating the instance you can view it by running "rhoas kafka describe".


```
rhoas kafka create [flags]
```

### Examples

```
# Start an interactive prompt to fill out the configuration values for the instance
$ rhoas kafka create

# Create a Kafka instance
$ rhoas kafka create --name my-kafka-instance

# Create a Kafka instance and output the result in YAML format
$ rhoas kafka create -o yaml

```

### Options

```
      --name string       Unique name of the Kafka instance
  -o, --output string     Specify the output format. Choose from: "json", "yaml", "yml" or "none"
      --provider string   Cloud Provider ID
      --region string     Cloud Provider Region ID
      --use               Set the new Kafka instance to the current instance (default true)
  -w, --wait              Wait until the Kafka instance is created
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

