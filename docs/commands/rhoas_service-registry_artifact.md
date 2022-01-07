## rhoas service-registry artifact

Manage Service Registry artifacts

### Synopsis

Manage Service Registry schema and API artifacts in the currently selected Service Registry instance.

Commands are executed on the currently selected Service Registry instance, which can be overridden using the --instance-id flag.

Service Registry enables developers to manage and share the structure of their data.
For example, client applications can dynamically push or pull the latest schema or API updates to or from the registry without needing to redeploy.
Service Registry also enables developers to create rules that govern how registry content can evolve over time.
For example, this includes rules for content validation and version compatibility.

Artifact commands enable client applications to manage the schema and API artifacts in the registry instance.
This set of commands provide create, read, update, and delete operations for artifacts, rules, versions, and metadata.


### Examples

```
## Create artifact in my-group from schema.json file
rhoas service-registry artifact create --artifact-id=my-artifact --group=my-group artifact.json

## Get artifact content
rhoas service-registry artifact get --artifact-id=my-artifact --group=my-group --output-file=artifact.json

## Get artifact content by hash
rhoas service-registry artifact download --hash=cab4...al9 --output-file=artifact.json

## Delete artifact
rhoas service-registry artifact delete --artifact-id=my-artifact

## Get artifact metadata
rhoas service-registry artifact metadata --artifact-id=my-artifact --group=my-group

## Update artifact
rhoas service-registry artifact update --artifact-id=my-artifact artifact-new.json

## List artifacts
rhoas service-registry artifact list --group=my-group --limit=10 page=1

## View artifact versions
rhoas service-registry artifact versions --artifact-id=my-artifact --group=my-group

Artifacts file can be instance of any schema like OpenAPI, Avro etc.
For example: https://raw.githubusercontent.com/redhat-developer/app-services-cli/main/docs/resources/avro-userInfo.json

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas service-registry](rhoas_service-registry.md)	 - Service Registry commands
* [rhoas service-registry artifact create](rhoas_service-registry_artifact_create.md)	 - Creates new artifact from file or standard input
* [rhoas service-registry artifact delete](rhoas_service-registry_artifact_delete.md)	 - Deletes single or all artifacts in a given group
* [rhoas service-registry artifact download](rhoas_service-registry_artifact_download.md)	 - Download artifacts from Service Registry using global identifiers
* [rhoas service-registry artifact export](rhoas_service-registry_artifact_export.md)	 - Export data from Service Registry instance
* [rhoas service-registry artifact get](rhoas_service-registry_artifact_get.md)	 - Get artifact by ID and group
* [rhoas service-registry artifact import](rhoas_service-registry_artifact_import.md)	 - Import data into a Service Registry instance
* [rhoas service-registry artifact list](rhoas_service-registry_artifact_list.md)	 - List artifacts
* [rhoas service-registry artifact metadata-get](rhoas_service-registry_artifact_metadata-get.md)	 - Get artifact metadata
* [rhoas service-registry artifact metadata-set](rhoas_service-registry_artifact_metadata-set.md)	 - Update artifact metadata
* [rhoas service-registry artifact state-set](rhoas_service-registry_artifact_state-set.md)	 - Set artifact state
* [rhoas service-registry artifact update](rhoas_service-registry_artifact_update.md)	 - Update artifact
* [rhoas service-registry artifact versions](rhoas_service-registry_artifact_versions.md)	 - Get latest artifact versions by artifact-id and group

