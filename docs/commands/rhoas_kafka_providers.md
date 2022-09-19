## rhoas kafka providers

List Kafka Cloud Providers

### Synopsis

List all enabled Cloud Providers and Regions for Kafka deployment


```
rhoas kafka providers [flags]
```

### Examples

```
# List all providers using the default output format
$ rhoas kafka providers 

# List all providersin JSON format
$ rhoas kafka providers  -o json

```

### Options

```
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances

