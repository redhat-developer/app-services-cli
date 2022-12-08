## rhoas service-account reset-credentials

Reset service account credentials

### Synopsis

Reset the credentials for a service account.

This command generates a new password for a service account. After the credentials have been reset, any applications or tools that use the credentials must be updated to use the new password.

You must specify an output format into which the credentials will be stored.

- env (default): Store credentials in an env file as environment variables
- json: Store credentials in a JSON file
- properties: Store credentials in a properties file, which is typically used in Java-related technologies
- java-kafka-properties: Store credentials in a properties file suitable for the Java Kafka client
- secret: Store credentials in a Kubernetes secret file


```
rhoas service-account reset-credentials [flags]
```

### Examples

```
# Start an interactive prompt to reset credentials
$ rhoas service-account reset-credentials

# Reset credentials for the service account specified and save the credentials to a JSON file
$ rhoas service-account reset-credentials --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd -o json

# Reset credentials for the service account specified and save the credentials to a file suitable for the Java Kafka client
$ rhoas service-account reset-credentials --id 173c1ad9-932d-4007-ae0f-4da74f4d2ccd -o java-kafka-properties

```

### Options

```
      --file-format string   Format in which to save the service account credentials (choose from: "env", "json", "properties", "secret")
      --id string            The unique ID of the service account for which you want to reset the credentials
      --output-file string   Sets a custom file location to save the credentials
      --overwrite            Forcibly overwrite a credentials file if it already exists
  -y, --yes                  Skip confirmation to forcibly reset service account credentials
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-account](rhoas_service-account.md)	 - Create, list, describe, delete, and update service accounts

