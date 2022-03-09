## rhoas service-registry rule

Manage artifact rules in a Service Registry instance

### Synopsis

Configure the validity and compatibility rules that govern artifact content.

When you add or update an artifact, Service Registry applies rules to check the validity and compatibility of the artifact content. Artifact-specific rules apply to the specified artifact only. Global rules apply to all artifacts in a particular Service Registry instance. Configured artifact-specific rules override any configured global rules. Before a new artifact version can be uploaded to the registry, all configured global rules or artifact-specific rules must pass.

For more information about supported Service Registry content and rules, see https://access.redhat.com/documentation/en-us/red_hat_openshift_service_registry/1/guide/9b0fdf14-f0d6-4d7f-8637-3ac9e2069817.


### Examples

```
## Enable the global compatibility rule for all artifacts in the current Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full

## Enable the global compatibility rule for all artifacts in a specific Service Registry instance
$ rhoas service-registry rule enable --rule-type=compatibility --config=full-transitive --instance-id=8ecff228-1ffe-4cf5-b38b-55223885ee00

## Display the configuration details of the global validity rule for the current Service Registry instance
$ rhoas service-registry rule describe --rule-type=validity

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry rule describe](rhoas_service-registry_rule_describe.md)	 - Display the configuration details of a rule
* [rhoas service-registry rule disable](rhoas_service-registry_rule_disable.md)	 - Disable validity and compatibility rules
* [rhoas service-registry rule enable](rhoas_service-registry_rule_enable.md)	 - Enable validity and compatibility rules
* [rhoas service-registry rule list](rhoas_service-registry_rule_list.md)	 - List the validity and compatibility rules
* [rhoas service-registry rule update](rhoas_service-registry_rule_update.md)	 - Update the configuration of rules

