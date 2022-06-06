## rhoas generate-config

Generate configurations for the service context

### Synopsis

Generate configuration files for the service context to connect with to be used with various tools and platforms

You must specify an output format into which the credentials will be stored:

- env (default): Store configurations in an env file as environment variables
- json: Store configurations in a JSON file
- properties: Store configurations in a properties file, which is typically used in Java-related technologies
- secret: Store configurations in a Kubernetes secret file


```
rhoas generate-config [flags]
```

### Examples

```
## Generate configurations for the current service context in json format
$ rhoas generate-config --type json

## Generate configurations for the current service context in env format and save it in specified path
$ rhoas generate-config --type env --output-file ./configs/.env

## Generate configurations for a specified context as Kubernetes secret
$ rhoas generate-config --name qaprod --type secret

```

### Options

```
      --client-id string       Client ID of the service account
      --client-secret string   Client secret of the service account
      --client-secret-stdin    Take the client secret from stdin
      --generate-auth          Create service account
      --name string            Name of the context
      --output-file string     Sets a custom file location to save the configurations
      --overwrite              Forcibly overwrite a configuration file if it already exists
      --type string            Type of configuration file to be generated
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI

