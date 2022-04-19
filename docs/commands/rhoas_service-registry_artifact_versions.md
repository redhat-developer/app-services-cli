## rhoas service-registry artifact versions

Get latest artifact versions by artifact-id and group

### Synopsis

Get the latest artifact versions by specifying the group and artifact-id

```
rhoas service-registry artifact versions [flags]
```

### Examples

```
## Get latest artifact versions for default group
rhoas service-registry artifact versions --artifact-id=my-artifact

## Get latest artifact versions for my-group group
rhoas service-registry artifact versions --artifact-id=my-artifact --group mygroup

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -o, --output string        Output format ("json", "yaml", "yml", "none")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

