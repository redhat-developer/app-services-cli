## rhoas service-registry artifact get

Get artifact by ID, group, and version

### Synopsis

Get an artifact from the registry based on the artifact ID, group, and version.

NOTE: To fetch artifacts using global identifiers, use the "service-registry artifact download" command.

When --version is specified, the command fetches the specified artifact version. When --version is not specified, the command fetches the latest artifact version.



```
rhoas service-registry artifact get [flags]
```

### Examples

```
## Get latest artifact with name "my-artifact" and print it out to standard out
rhoas service-registry artifact get --artifact-id=my-artifact

## Get latest artifact with name "my-artifact" from group "my-group" and save it to artifact.json file
rhoas service-registry artifact get --artifact-id=my-artifact --group=my-group --output-file=artifact.json

## Get latest artifact and pipe it to another command
rhoas service-registry artifact get --artifact-id=my-artifact | grep -i 'user'

## Get artifact with custom version and print it out to standard out
rhoas service-registry artifact get --artifact-id=myartifact --version=4

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -o, --output string        Output format (table, json, yaml, yml) (default "table")
      --output-file string   Location of the output file
      --references           Get a list of reference records for the given artifact. Is returned instead of the artifact content
      --version string       Version of the artifact
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

