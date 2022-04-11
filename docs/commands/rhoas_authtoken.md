## rhoas authtoken

Output the current token

### Synopsis

View the authentication token of the current user that can be used to 
make general API requests against api.openshift.com APIs.

This command outputs the token for the user currently logged in.


```
rhoas authtoken [flags]
```

### Examples

```
# Returns header with token used for authorization
$ echo Authorization: BEARER ${rhoas authtoken}

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

