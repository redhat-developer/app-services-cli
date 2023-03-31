## rhoas service-registry artifact create

Create new artifact from file or standard input

### Synopsis

Create a new artifact by posting the artifact content to the registry instance.
Artifact content can be read from a file or downloaded from a URL, either locally by this CLI, or by the Service Registry server.

Artifacts are typically in JSON format for most of the supported types, but might be in another format for a few types (for example, PROTOBUF).

Service Registry attempts to identify what type of artifact is being added from the following supported list:

* Avro (AVRO)
* Protobuf (PROTOBUF)
* JSON Schema (JSON)
* Kafka Connect (KCONNECT)
* OpenAPI (OPENAPI)
* AsyncAPI (ASYNCAPI)
* GraphQL (GRAPHQL)
* Web Services Description Language (WSDL)
* XML Schema (XSD)

An artifact is created using the content provided in the request body.
This content is created with a unique artifact ID that can be provided by user.
If not provided in the request, the registry server generates a unique ID for the artifact.
It is typically recommended that callers provide the ID, because this is a meaningful identifier, and for most use cases should be supplied by the caller.
If an artifact with the provided artifact ID already exists, the command will fail with an error.

When the --group parameter is missing, the command uses the "default" group.
when the --instance-id is missing, the command creates a new artifact for the currently active Service Registry instance (displayed in rhoas service-registry describe)


```
rhoas service-registry artifact create [flags]
```

### Examples

```
# Create an artifact in the default group
rhoas service-registry artifact create my-artifact.json

# Create an artifact with the specified type
rhoas service-registry artifact create --type=JSON my-artifact.json

# Create an artifact from a URL, dowloaded by the CLI:
rhoas service-registry artifact create https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.json

# Create an artifact from a URL, dowloaded by the Service Registry server:
rhoas service-registry artifact create --download-on-server https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.json

```

### Options

```
      --artifact-id string            ID of the artifact
      --description string            Custom description of the artifact
      --download-on-server            If enabled, artifact created by providing a URL will be downloaded on the server
      --file string                   File location of the artifact
  -g, --group string                  Artifact group (default "default")
      --instance-id string            ID of the Service Registry instance to be used (by default, uses the currently selected instance)
      --name string                   Custom name of the artifact
  -o, --output string                 Output format (json, yaml, yml) (default "json")
  -r, --reference stringArray         One or multiple artifact reference records in the format "<reference name>=<group>:<artifactId>:<version>". When referencing an artifact from the default group and/or with the latest version, these parts may be left empty
      --reference-separators string   Specify alternative separator characters when specifying reference records for an artifact. Two distinct characters in order are required (default "=:")
  -t, --type string                   Type of artifact. Run "rhoas service-registry artifact types" to get a list
      --version string                Custom version of the artifact (for example 1.0.0)
```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry artifact](rhoas_service-registry_artifact.md)	 - Manage Service Registry artifacts

