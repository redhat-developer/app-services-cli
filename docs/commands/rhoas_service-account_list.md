## rhoas service-account list

List all service accounts

### Synopsis

List all service accounts.

This command provides a high-level view of all service accounts.

The service accounts are displayed by default in a table, but can also be displayed in JSON or YAML format.


```
rhoas service-account list [flags]
```

### Examples

```
# List all service accounts using the default output format
$ rhoas service-account list

# List all service accounts using JSON as the output format
$ rhoas service-account list -o json

```

### Options

```
  -o, --output string   Format in which to display the service accounts (choose from: "json", "yml", "yaml")
      --page int32      Current page number for the list (default 1)
      --size int32      Maximum number of items to be returned per page (default 100)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts

