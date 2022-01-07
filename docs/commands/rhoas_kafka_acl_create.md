## rhoas kafka acl create

Create a Kafka ACL

### Synopsis

Create Kafka Access Control List (ACL) rules. A Kafka ACL defines how other user accounts and service accounts can interact with a Kafka instance and its resources.

```
rhoas kafka acl create [flags]
```

### Examples

```
# Create an ACL for user "dev_user" on all topics
$ rhoas kafka acl create --operation all --permission allow --topic "*" --user dev_user

# Create an ACL for a service account
$ rhoas kafka acl create --operation all --permission allow --topic "rhoas" --prefix --service-account "srvc-acct-11924479-43fe-42b4-9676-cf0c9aca81"

# Create an ACL for all users for the consumer group resource
$ rhoas kafka acl create --operation all --permission allow --group "group-1" --all-accounts

```

### Options

```
      --all-accounts              Set the ACL principal to match all principals (users and service accounts)
      --cluster                   Set the resource type to cluster
      --group string              Set the consumer group resource. When the --prefix option is also passed, this is used as the consumer group prefix
      --instance-id string        Kafka instance ID. Uses the current instance if not set
      --operation string          Set the ACL operation. Choose from: "all", "alter", "alter-configs", "create", "delete", "describe", "describe-configs", "read", "write"
      --permission string         Set the ACL permission. Choose from: "allow", "deny"
      --prefix                    Determine if the resource should be exact match or prefix
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

