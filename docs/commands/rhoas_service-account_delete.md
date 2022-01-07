## rhoas service-account delete

Delete a service account

### Synopsis

Permanently delete a service account.

When you delete a service account, any applications and tools that use the service account credentials to connect to Kafka instances will no longer be able to connect to them.


```
rhoas service-account delete [flags]
```

### Examples

```
# Delete a service account
$ rhoas service-account delete --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd

```

### Options

```
      --id string   The unique ID of the service account to delete
  -y, --yes         Skip confirmation to forcibly delete this service account
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts

