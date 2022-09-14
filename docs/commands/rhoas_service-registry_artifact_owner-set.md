## rhoas service-registry artifact owner-set

Set owner of the artifact

### Synopsis

Set owner of the specific artifact


```
rhoas service-registry artifact owner-set [flags]
```

### Examples

```
## Set owner of the artifact with artifact id 'example-name' in group 'example-group' to 'new-owner-name'
$ rhoas service-registry artifact owner-get --artifact-id example-name --group example-group --owner new-owner-name

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
      --owner string         Name of new owner
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

