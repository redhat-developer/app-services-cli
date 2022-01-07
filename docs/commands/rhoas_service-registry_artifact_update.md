## rhoas service-registry artifact update

Update artifact

### Synopsis

Update artifact from file or directly standard input

Artifacts are typically in JSON format for most supported types, but may be in another format for a few (for example, PROTOBUF).
The type of the content should be compatible with the current artifact type.

When successful, this creates a new version of the artifact, making it the most recent (and therefore official) version of the artifact.

An artifact is updated using the content provided in the request body.
This content is updated under a unique artifactId provided by user.

Updated artifact content should conform to validity and compatibility rules set for the registry instance.


```
rhoas service-registry artifact update [flags]
```

### Examples

```
## update artifact from group and artifact-id
rhoas service-registry artifact update --artifact-id=my-artifact --group my-group my-artifact.json

## update artifact from group and artifact-id
rhoas service-registry artifact update --artifact-id=my-artifact --group my-group my-artifact.json

```

### Options

```
      --artifact-id string   ID of the artifact
      --description string   Custom description of the artifact
  -f, --file string          File location of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used. By default, uses the currently selected instance
      --name string          Custom name of the artifact
      --version string       Custom version of the artifact (for example 1.0.0)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

