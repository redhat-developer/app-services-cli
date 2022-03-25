## rhoas kafka acl list

List all Kafka ACL rules

### Synopsis

View the level of access that other user accounts and service accounts have to a Kafka instance. This is displayed as a list of Access Control Lists (ACLs).

An ACL maps permitted operations on specified resources for a selected account or for all accounts in an organization. Each ACL shows a single named account (or all accounts), a permission (such as "Allow"), an operation (such as "Write"), and a Kafka resource type and identifier. The resource type is a Kafka resource (such as "Topic"), and the resource identifier identifies the occurrences of the resource (for example, "Is *" denotes any occurrences of the resource).

By default, new Kafka instances contain the following ACLs:

PRINCIPAL (4)    PERMISSION   OPERATION          DESCRIPTION
---------------- ------------ ------------------ ----------------
All accounts     allow        describe           group is "*"
All accounts     allow        describe           cluster is "*"
All accounts     allow        describe-configs   topic is "*"
All accounts     allow        describe           topic is "*"

These ACLs allow all accounts in the organization to view the Kafka instance permissions and to view topics and consumer groups in the instance, but not to produce or consume messages.

The ACLs are displayed in a table by default. Alternatively, you can display them as JSON or YAML.


```
rhoas kafka acl list [flags]
```

### Examples

```
# Display Kafka ACL rules for the Kafka instance
$ rhoas kafka acl list

# Display Kafka ACL rules for a specific user
$ rhoas kafka acl list --user foo_user

# Display Kafka ACL rules for a specific service account
$ rhoas kafka acl list --service-account srvc-acct-f20a7561-7426-4f5a-b5e7-0ef2db31e15b

# Display Kafka ACL rules for a specific topic
$ rhoas kafka acl list --topic foo_topic_name

# Display Kafka ACL rules for a specific consumer group
$ rhoas kafka acl list --group foo_group_id

# Display Kafka ACL rules for a specific consumer group and user
$ rhoas kafka acl list --group foo_group_id --user foo_user

```

### Options

```
      --all-accounts             Set the ACL principal to match all principals (users and service accounts)
      --cluster                  Set filter to cluster resource
      --group string             Text search to filter ACL rules for consumer groups by ID
      --instance-id string       Kafka instance ID. Uses the current instance if not set 
  -o, --output string            Specify the output format. Choose from: "json", "yaml", "yml"
      --page int32               Current page number for the list  (default 1)
      --service-account string   Service account client ID used as principal for this operation
      --size int32               Maximum number of items to be returned per page  (default 10)
      --topic string             Text search to filter ACL rules for topics by name
      --user string              User ID to be used as principal
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka acl](rhoas_kafka_acl.md)	 - Manage Kafka ACLs for users and service accounts

