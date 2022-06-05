## rhoas service-account create

Create a service account

### Synopsis

Create a service account with credentials that are saved to a file.

Applications and tools use these service account credentials to authenticate and interact with your application services.

You must specify an output format into which the credentials will be stored.

- env (default): Store credentials in an env file as environment variables
- json: Store credentials in a JSON file
- properties: Store credentials in a properties file, which is typically used in Java-related technologies
- secret: Store credentials in a Kubernetes secret file


```
rhoas service-account create [flags]
```

### Examples

```
# Create a service account through an interactive prompt
$ rhoas service-account create

# Create a service account and save the credentials in a JSON file
$ rhoas service-account create --file-format json

# Create a service account and forcibly overwrite the credentials file if it exists already
$ rhoas service-account create --overwrite

# Create a service account and save credentials to a custom file location
$ rhoas service-account create --output-file=./service-acct-credentials.json

```

### Options

```
      --file-format string         Format in which to save the service account credentials (choose from: "env", "json", "properties", "secret")
      --output-file string         Sets a custom file location to save the credentials
      --overwrite                  Forcibly overwrite a credentials file if it already exists
      --short-description string   Short description of the service account
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts

