## rhoas service-registry rule disable

Disable validity and compatibility rules

### Synopsis

Disable validity and compatibility rules for the specified Service Registry instance or artifact

```
rhoas service-registry rule disable [flags]
```

### Examples

```
## Disable global compatibility rule for artifacts of the current Service Registry instance
$ rhoas service-registry rule disable --rule-type=compatibility

## Disable global compatibility rule for artifacts of a specific Service Registry instance
$ rhoas service-registry rule disable --rule-type=compatibility --instance-id 8ecff228-1ffe-4cf5-b38b-55223885ee00

## Disable validity rule for a specific artifact
$ rhoas service-registry rule disable --rule-type=validity --artifact-id=my-artifact

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
      --rule-type string     Rule type determines how the content of an artifact can evolve over time
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry rule](rhoas_service-registry_rule.md)	 - Manage artifact rules in a Service Registry instance

