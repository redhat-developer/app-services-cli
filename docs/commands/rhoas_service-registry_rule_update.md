## rhoas service-registry rule update

Update the configuration of rules

### Synopsis

Update the configuration of validity and compatibility rules for the specified Service Registry instance or artifact.

```
rhoas service-registry rule update [flags]
```

### Examples

```
## Update the global compatibility rule for all artifacts in the current Service Registry instance
$ rhoas service-registry rule update --rule-type=compatibility --config=full

## Update the global compatibility rule for all artifacts in a specific Service Registry instance
$ rhoas service-registry rule update --rule-type=compatibility --config=full --instance-id=8ecff228-1ffe-4cf5-b38b-55223885ee00

## Update the validity rule for a specific artifact
$ rhoas service-registry rule update --rule-type=validity --config=full --artifact-id=my-artifact

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

