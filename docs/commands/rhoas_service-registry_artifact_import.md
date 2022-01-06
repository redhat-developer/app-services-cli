## rhoas service-registry artifact import

Import data into a Service Registry instance

### Synopsis

Import all artifacts and metadata from export file to another Service Segistry instance


```
rhoas service-registry artifact import [flags]
```

### Examples

```
## Import all artifacts and metadata from export file to another Service Registry instance
rhoas service-registry artifact import --file=export.zip

```

### Options

```
      --file string          File location of the artifact
      --instance-id string   ID of the Service Registry instance to be used. By default, uses the currently selected instance
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

