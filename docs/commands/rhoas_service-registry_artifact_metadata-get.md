## rhoas service-registry artifact metadata-get

Get artifact metadata

### Synopsis

Get the metadata for an artifact in a Service Registry instance.

The returned metadata includes both generated (read-only) and editable metadata (such as name and description).


```
rhoas service-registry artifact metadata-get [flags]
```

### Examples

```
## Get latest artifact metadata for default group
rhoas service-registry artifact metadata-get --artifact-id=my-artifact

## Get latest artifact metadata for my-group group
rhoas service-registry artifact metadata-get --artifact-id=my-artifact --group mygroup

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -o, --output string        Output format (json, yaml, yml)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

