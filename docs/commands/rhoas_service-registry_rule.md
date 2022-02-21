## rhoas service-registry rule

Manage artifact rules in a Service Registry instance

### Synopsis

Set of commands for managing validity and compatibility rules for an entire Service Registry instance or individual for each artifact.
Service Registry rules govern how artifact content can evolve over time by validating for any artifact updates

Rules can be configured as global rules for entire Service Registry instances or specific artifacts.


### Examples

```
## Enable global compatibility rule for artifacts of the current Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full

## Enable global compatibility rule for artifacts of a specific Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full-transitive --instance-id 8ecff228-1ffe-4cf5-b38b-55223885ee00

## display configuration details of global validity rule for current Service Registry instance
$ rhoas service-registry rule describe --rule-type=validity

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry rule enable](rhoas_service-registry_rule_enable.md)	 - Enable validity and compatibility rules

