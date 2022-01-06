## rhoas service-registry delete

Delete a Service Registry instance

### Synopsis

 
Delete a Service Registry instance along with all of its schema and API artifacts.


```
rhoas service-registry delete [flags]
```

### Examples

```
# Delete a Service Registry instance by name
rhoas service-registry delete --name my-service-registry

# Delete a Service Registry instance by ID
rhoas service-registry delete --id 1iSY6RQ3JKI8Q0OTmjQFd3ocFRg

```

### Options

```
      --id string     Unique ID of the Service Registry instance you want to delete (if not provided, the current Service Registry instance will be deleted)
      --name string   Name of the Service Registry instance to delete
  -y, --yes           Skip confirmation to forcibly delete this Service Registry instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands

