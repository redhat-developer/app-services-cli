## rhoas context create

Create context

### Synopsis

Create a context and assign service identifiers

```
rhoas context create [flags]
```

### Examples

```
## Create context with Kafka ID and Service Registry ID
$ rhoas context create --name my-context --kafka-id c8696ncpoj7gdjmaiqog --registry-id 0282d488-52b3-405b-9e30-9f6f9407de57

```

### Options

```
      --kafka-id string      Name of the context
      --name string          Name of the context
      --registry-id string   Name of the context
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas context](rhoas_context.md)	 - Group, share and manage your rhoas services

