## rhoas service-registry artifact export

Export data from Service Registry instance

### Synopsis

Export all artifacts and metadata from a Service Registry instance to a specified file


```
rhoas service-registry artifact export [flags]
```

### Examples

```
## Export all artifacts and metadata to export file for another Service Registry instance
rhoas service-registry artifact export --file=export.zip

```

### Options

```
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
      --output-file string   File location of the artifact
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

