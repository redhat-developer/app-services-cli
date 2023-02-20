## rhoas kafka acl

Manage Kafka ACLs for users and service accounts

### Synopsis

Manage Kafka Access Control Lists (ACLs). An ACL maps permitted operations on specified resources for a specified principal (username or service account) or for all accounts in an organization.

You can use these commands to manage how other user accounts and service accounts are permitted to access Kafka resources. You can manage access for only the Kafka instances that you create or for instances that the owner has enabled you to access and alter.


### Examples

```
# Grant access to principal for consuming messages from all topics
$ rhoas kafka acl grant-access --consumer --user foo_user --topic all --group all

# Grant access to principal for producing messages to all topics
$ rhoas kafka acl grant-access --producer --user foo_user --topic all

# List ACL rules for a Kafka instance
$ rhoas kafka acl list

# Give admin rights to user "abc"
$ rhoas kafka acl grant-admin --user abc

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas kafka](rhoas_kafka.md)	 - Create, view, use, and manage your Kafka instances
* [rhoas kafka acl create](rhoas_kafka_acl_create.md)	 - Create a Kafka ACL
* [rhoas kafka acl delete](rhoas_kafka_acl_delete.md)	 - Delete Kafka ACLs matching the provided filters
* [rhoas kafka acl grant-access](rhoas_kafka_acl_grant-access.md)	 - Add ACL rules to grant users access to produce and consume from topics
* [rhoas kafka acl grant-admin](rhoas_kafka_acl_grant-admin.md)	 - Grant an account permissions to create and delete ACLs in the Kafka instance
* [rhoas kafka acl list](rhoas_kafka_acl_list.md)	 - List all Kafka ACL rules

