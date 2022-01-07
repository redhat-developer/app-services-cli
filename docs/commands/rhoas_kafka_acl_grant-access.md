## rhoas kafka acl grant-access

Add ACL rules to grant users access to produce and consume from topics

### Synopsis

Create Access Control List (ACL) rules that grant the specified user access to produce and consume from topics.

```
rhoas kafka acl grant-access [flags]
```

### Examples

```
# Grant access to principal for consuming messages from all topics
$ rhoas kafka acl grant-access --consumer --user user_name --topic all --group all

# Grant access to principal for consuming messages from all topics in a specified instance
$ rhoas kafka acl grant-access --consumer --user user_name --topic all --group all --instance-id c5hv7iru4an1g84pogp0

# Grant access to principal for producing messages to all topics
$ rhoas kafka acl grant-access --producer --user user_name --topic all

# Grant access to principal for consuming messages from topics starting with "abc"
$ rhoas kafka acl grant-access --consumer --user user_name --topic-prefix "abc" --group my-group

# Grant access to principal for producing messages to topics starting with "abc"
$ rhoas kafka acl grant-access --producer --user user_name --topic-prefix "abc"

# Grant access to all users for consuming messages from topic "my-topic"
$ rhoas kafka acl grant-access --consumer --all-accounts --topic my-topic --group my-group

# Grant access to all users for producing messages to topic "my-topic"
$ rhoas kafka acl grant-access --producer --all-accounts --topic my-topic

# Grant access to principal for produce and consume messages from all topics
$ rhoas kafka acl grant-access --producer --consumer --user user_name --topic all --group all

```

### Options

```
      --all-accounts             Set the ACL principal to match all principals (users and service accounts)
      --consumer                 Add ACL rules that grant the specified principal access to consume messages from topics
      --group string             Consumer group ID to define ACL rules for
      --group-prefix string      Prefix name for groups to be selected
      --instance-id string       Kafka instance ID. Uses the current instance if not set
      --producer                 Add ACL rules that grant the specified principal access to produce messages to topics
      --service-account string   Service account client ID used as principal for this operation
      --topic string             Topic name to define ACL rules for
      --topic-prefix string      Prefix name for topics to be selected
      --user string              User ID to be used as principal
  -y, --yes                      Skip confirmation of this action 
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka acl](rhoas_kafka_acl.md)	 - Manage Kafka ACLs for users and service accounts

