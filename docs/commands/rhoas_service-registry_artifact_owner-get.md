## rhoas service-registry artifact owner-get

Get owner of artifact

### Synopsis

Get owner of the artifact


```
rhoas service-registry artifact owner-get [flags]
```

### Examples

```
## Get owner of the artifact with artifact id 'example-name' in group 'example-group'
$ rhoas service-registry artifact owner-get --artifact-id example-name --group example-group

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

