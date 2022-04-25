## rhoas service-registry rule enable

Enable validity and compatibility rules

### Synopsis

Enable validity and compatibility rules for the specified Service Registry instance or artifact.

```
rhoas service-registry rule enable [flags]
```

### Examples

```
## Enable the global compatibility rule for all artifacts in the current Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full

## Enable the global compatibility rule for all artifacts in a specific Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full-transitive --instance-id=8ecff228-1ffe-4cf5-b38b-55223885ee00

## Enable the validity rule for a specific artifact
$ rhoas service-registry rule enable --rule-type=validity --config=syntax-only --artifact-id=my-artifact

```

### Options

```
      --artifact-id string   ID of the artifact
      --config string        Configuration value for a rule
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

