## rhoas service-account

Create, list, describe, delete, and update service accounts

### Synopsis

Manage service accounts. Service accounts enable you to connect your applications to a Kafka instance.

You can create, list, describe, delete, and update service accounts. You can also reset the credentials for a service account.


### Examples

```
# Create a service account
rhoas service-account create

# List all service accounts
rhoas service-account list

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas service-account create](rhoas_service-account_create.md)	 - Create a service account
* [rhoas service-account delete](rhoas_service-account_delete.md)	 - Delete a service account
* [rhoas service-account describe](rhoas_service-account_describe.md)	 - View configuration details for a service account
* [rhoas service-account list](rhoas_service-account_list.md)	 - List all service accounts
* [rhoas service-account reset-credentials](rhoas_service-account_reset-credentials.md)	 - Reset service account credentials

