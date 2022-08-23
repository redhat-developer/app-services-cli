## rhoas connector type list

List connector types

### Synopsis

List the types of connectors that are available in the Connectors catalog.'

Use filter options, such as --limit, --page, and --search


```
rhoas connector type list [flags]
```

### Examples

```
# List all connector types
rhoas connector type list

# List connector types with a limit of 10 connector types and 2 pages
rhoas connector type list --limit=10 --page=2

# List all connector types that start with "Amazon"
rhoas connector type list --search=Amazon%

# List all connector types that contain the word "Amazon"
rhoas connector type list --search=%Amazon%

```

### Options

```
      --limit int       Page of the list based on the limit value (default 10)
  -o, --output string   Specify the output format. Choose from: "json", "yaml", "yml"
      --page int        Page of the list based on the limit value (default 1)
      --search string   Search query for name of connector type
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas connector type](rhoas_connector_type.md)	 - List and get details of the different connector types

