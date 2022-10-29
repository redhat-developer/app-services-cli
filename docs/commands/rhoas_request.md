## rhoas request

Allows users to perform API requests against the API server

### Synopsis

Command allows users to perform API requests against the API server.


```
rhoas request [flags]
```

### Examples

```
# Perform a GET request to the specified path
rhoas request --path /api/kafkas_mgmt/v1/kafkas

# Perform a POST request to the specified path
# cat request.json | rhoas request --path "/api/kafkas_mgmt/v1/kafkas?async=true" --method post

```

### Options

```
      --method string   HTTP method to use. (get, post) (default "GET")
      --path string     Path to send request. For example /api/kafkas_mgmt/v1/kafkas?async=true
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

