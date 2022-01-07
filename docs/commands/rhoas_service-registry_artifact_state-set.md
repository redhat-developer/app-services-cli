## rhoas service-registry artifact state-set

Set artifact state

### Synopsis

Set artifact state by setting one of possible values
- ENABLED (Enable artifact)
- DISABLED (Disable artifact usage)
- DEPRECATED (Deprecate artifact)


```
rhoas service-registry artifact state-set [flags]
```

### Examples

```
## Set artifact state to DISABLED
rhoas service-registry artifact state-set --artifact-id=my-artifact --state=DISABLED

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used. By default, uses the currently selected instance
      --state string         new artifact state
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

