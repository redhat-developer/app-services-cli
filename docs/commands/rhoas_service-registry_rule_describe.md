## rhoas service-registry rule describe

Display the configuration details of a rule

### Synopsis

Display the configuration details of a compatibility or validity rule for the specified Service Registry instance or artifact.

```
rhoas service-registry rule describe [flags]
```

### Examples

```
## Display the configuration details of the global validity rule for the current Service Registry instance
$ rhoas service-registry rule describe --rule-type=validity

## Display the configuration details of the compatibility rule for a specific artifact
$ rhoas service-registry rule describe --rule-type=compatibility --artifact-id=my-artifact --group=my-group

```

### Options

```
      --artifact-id string   ID of the artifact
  -g, --group string         Artifact group (default "default")
      --instance-id string   ID of the Service Registry instance to be used (by default, uses the currently selected instance)
  -o, --output string        Specify the output format. Choose from: "json", "yaml", "yml"
      --rule-type string     Rule type determines how the content of an artifact can evolve over time
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry rule](rhoas_service-registry_rule.md)	 - Manage artifact rules in a Service Registry instance

