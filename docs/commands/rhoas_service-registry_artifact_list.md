## rhoas service-registry artifact list

List artifacts

### Synopsis

List all artifacts for the group in the specified output format (by default, "table")

```
rhoas service-registry artifact list [flags]
```

### Examples

```
## List all artifacts for the "default" artifact group
rhoas service-registry artifact list

## List all artifacts with "my-group" group
rhoas service-registry artifact list --group=my-group

## List all artifacts with limit and group
rhoas service-registry artifact list --page=2 --limit=10

## List all artifacts for the "default" artifact group with name containing "sample"
rhoas service-registry artifact list --name sample

## List all artifacts for the "default" artifact group having labels "my-label" and "sample"
rhoas service-registry artifact list --label "my-label" --labels "sample"

## List all artifacts for the "default" artifact group with description containing "sample"
rhoas service-registry artifact list --description sample

```

### Options

```
      --description string     Text search to filter artifacts by description
  -g, --group string           Artifact group (default "default")
      --instance-id string     ID of the Service Registry instance to be used (by default, uses the currently selected instance)
      --label stringArray      Text search to filter artifacts by labels
      --limit int32            Page limit (default 100)
      --name string            Text search to filter artifacts by name
  -o, --output string          Output format (json, yaml, yml)
      --page int32             Page number (default 1)
      --property stringArray   Text search to filter artifacts by properties (separate each name/value pair using a colon)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

