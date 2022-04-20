## rhoas service-registry describe

Describe a Service Registry instance

### Synopsis

Describe a Service Registry instance. Fetch all required fields including the registry URL.


```
rhoas service-registry describe [flags]
```

### Examples

```
# Describe a Service Registry instance by name
rhoas service-registry describe --name my-service-regisrty


# Describe a Service Registry instance by ID
rhoas service-registry describe --id 1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string       Unique ID of the Service Registry instance (if not provided, the current Service Registry instance will be used)
      --name string     Name of the Service Registry instance to view
  -o, --output string   Format in which to display the Service Registry instance (choose from: "json", "yml", "yaml") (default "json")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands

