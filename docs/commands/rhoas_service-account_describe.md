## rhoas service-account describe

View configuration details for a service account

### Synopsis

View configuration details for a service account.

Use the “--id” flag to specify which service account you would like to view.

You can view the output as either JSON or YAML.


```
rhoas service-account describe [flags]
```

### Examples

```
# View a specific service account
$ rhoas service-account describe --id=8a06e685-f827-44bc-b0a7-250bc8abe52e --output yml

```

### Options

```
      --id string       The unique ID of the service account to view
  -o, --output string   Format in which to display the service account (choose from: "json", "yml", "yaml" or "none") (default "json")
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts

