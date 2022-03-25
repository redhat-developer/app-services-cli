## rhoas kafka acl grant-admin

Grant an account permissions to create and delete ACLs in the Kafka instance

### Synopsis

Grant administrative rights to a user that enable the user to create and delete Access Control Lists (ACLs) in a Kafka instance.

```
rhoas kafka acl grant-admin [flags]
```

### Examples

```
# Give admin rights to user "abc"
$ rhoas kafka acl grant-admin --user abc

# Give admin rights to a service account
$ rhoas kafka acl grant-admin --service-account srvc-acct-0837725a-4e69-44e1-af3b-29da30aa85ce

# Give admin rights to all accounts for a specific kafka instance
$ rhoas kafka acl grant-admin --all-accounts --instance-id c5hv7iru4an1g84pogp0

```

### Options

```
      --all-accounts             Set the ACL principal to match all principals (users and service accounts)
      --instance-id string       Kafka instance ID. Uses the current instance if not set 
      --service-account string   Service account client ID used as principal for this operation
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

