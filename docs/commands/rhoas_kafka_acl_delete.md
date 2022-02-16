## rhoas kafka acl delete

Delete Kafka ACLs matching the provided filters

### Synopsis

Delete Access Control List (ACL) rules from a Kafka instance.

```
rhoas kafka acl delete [flags]
```

### Examples

```
# Delete an ACL for user "dev_user" on all topics
$ rhoas kafka acl delete --operation write --permission allow --topic all --user dev_user

# Delete an ACL for a service account
$ rhoas kafka acl delete --operation all --permission any --topic "rhoas" --prefix --service-account "srvc-acct-11924479-43fe-42b4-9676-cf0c9aca81"

# Delete all ACLs for a service account
$ rhoas kafka acl delete --service-account "srvc-acct-11924479-43fe-42b4-9676-cf0c9aca81 --pattern-type=all"

# Delete an ACL for all users on the consumer group resource
$ rhoas kafka acl delete --operation all --permission any --group "group-1" --all-accounts

```

### Options

```
      --all-accounts              Set the ACL principal to match all principals (users and service accounts)
      --cluster                   Set the resource type to cluster
      --group string              Set the consumer group resource. When the --prefix option is also passed, this is used as the consumer group prefix
      --instance-id string        Kafka instance ID. Uses the current instance if not set
      --operation string          Set the ACL operation. Choose from: "all", "alter", "alter-configs", "create", "delete", "describe", "describe-configs", "read", "write"
  -o, --output string             Specify the output format. Choose from: "json", "yaml", "yml"
      --pattern-type string       Determine if the resource should be exact match, prefix or any [any literal prefix] (default "literal")
      --permission string         Set the ACL permission. Choose from: "allow", "any", "deny" (default "any")
      --prefix                    DEPRECATED: Determine if the resource should be exact match or prefix. Use --pattern-type instead
      --service-account string    Service account client ID used as principal for this operation
      --topic string              Set the topic resource. When the --prefix option is also passed, this is used as the topic prefix
      --transactional-id string   Set the transactional ID resource
      --user string               User ID to be used as principal
  -y, --yes                       Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka acl](rhoas_kafka_acl.md)	 - Manage Kafka ACLs for users and service accounts

