## rhoas service-registry rule list

List validity and compatibility rules

### Synopsis

View validity and compatibility rules for the selected Service Registry instance or specific artifact

```
rhoas service-registry rule list [flags]
```

### Examples

```
## List global rules for artifacts of the current Service Registry instance
$ rhoas service-registry rule list

## List global rule for artifacts of a specific Service Registry instance
$ rhoas service-registry rule list --instance-id 8ecff228-1ffe-4cf5-b38b-55223885ee00

## List rule for a specific artifact
$ rhoas service-registry rule list --artifact-id=my-artifact

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

* [rhoas service-registry rule](rhoas_service-registry_rule.md)	 - Manage artifact rules in a Service Registry instance

