
# RHOAS OpenAPI Standards

This document is intended as a collaborative document to agree and define standards for the design of RHOAS OpenAPI specifications.

This document is not intended as a replacement for https://github.com/redhat-developer/app-services-api-guidelines. It is intended to complement and build upon it in the context of Managed Services APIs.


## Versioning



* Set the version number in the _info.version_ field.
* The version should follow [Semantic Versioning](https://semver.org/).
    * A **MAJOR** version change is required when you make incompatible API changes.
        * Removal of an endpoint.
        * Changes to an endpoint URI structure.
        * Modifying a schema object such as a request or response payload object.
        * Removal of a query parameter.
        * Making a non-required parameter required.
        * Modifying a field name.
        * Modifying authorization.
        * Introducing rate-limiting.
        * Removal of a field
        * Addition of a required field in the request object
        * Addition of validation that was not present before
    * A **MINOR **version change is required when an additive change is made to the OpenAPI document.
        * Adding an endpoint.
        * Adding optional query parameters.
        * Adding optional fields to a schema used in a request payload.
        * Adding a response field.
    * A **PATCH **version change is required when a change which fixes broken functionality of the schema.
        * Invalid reference to a schema object
        * Typo in a property such as a tag
        * Errors in the document itself, such as a string that is not closed.
    * A pre-release extension may be added to indicate that the API version is unstable. For example: 2.0.0-alpha1.
* APIs must not expose the minor or patch version in the API URL.
* For releases which have alpha or beta stability the API **must** append the stability level after the major version using _[Channel-based versioning](https://cloud.google.com/apis/design/versioning). _For example: **v1beta**, **v2alpha**.
* All resource endpoints must include the channel version; `/v1`, `/v2`, `v2beta` etc.
    * /api/srs_mgmt/v1/registries
    * /api/kafka_mgmt/v1/kafkas
    * /api/kafka_mgmt/v2beta/kafkas
* The name of the OpenAPI file should include the version using the following format: \
“${service_or_api_name}.{yaml | json}”. Some examples:
    * kas-fleet-manager.yaml
    * srs-fleet-manager.json
    * kafka-admin-api.yaml


## Rules



* Should use OpenAPI v3+
* `info.license` must be Apache 2.0
* Every endpoint must have an _operationId_ specified.
* The `servers` config array is required for control plane APIs and should look as follows: \


```yaml
servers:
  - url: https://api.openshift.com
    description: Main (production) server
  - url: https://api.stage.openshift.com
    description: Staging server
  - url: http://localhost:8000
    description: localhost
  - url: /
    description: current domain
```


* Examples should be provided where possible/appropriate for documentation purposes.


## Naming


### API Paths

API endpoints should use _snake_case _for word separation.

Examples:

* /cloud_providers/..
* /service_accounts/../reset_credentials


### Fields

Fields should use _snake_case_ for word separation.

Examples:

* cloud_provider_region
* bootstrap_server_host


### Operations

The **operationId **will be used to name methods in the API clients. As such, changes to the names should be treated as breaking changes, and so we should be careful and explicit with our naming of them. 


* **get{Resource}By{Argument} **- get a single Resource
    * getKafkaById
    * getRegistryById
    * getServiceAccountById
    * getTopicByName
* **get{Resources}** - get multiple Resources
    * getRegistries
    * getKafkas
    * getServiceAccounts
    * getTopics
* **delete{Resource}?By{Arg}** - Delete a single resource or multiple resources
    * deleteKafkaById
    * deleteTopicByName
    * deleteTopics
    * deleteRegistryById
* **update{Resource}?By{Argument}** - update a resource or multiple resources
    * updateKafkaById
    * updateKafkaByName
    * updateRegistryById
    * updateTopics
* **create{Resource}**- create a resource
    * createKafka
    * createRegistry
    * createTopic


## Security

The following security scheme **must** be added to the OpenAPI document. 

Depending on your API this can be applied globally, or per endpoint.


```yaml
# 1) Define the security scheme type (HTTP bearer)
components:
  securitySchemes:
    bearer:                # arbitrary name for the security scheme
      type: http
      scheme: bearer
      bearerFormat: JWT    # optional, arbitrary value for documentation purposes
# 2) Apply the security globally to all operations
security:
  - bearer: []         # use the same name as above
```



## Tags

Grouping operations with tags is a good practice and is recommended.

If tags are used a root level global “tags” section should be defined matching those used in operations.


## Schemas

There are a number of schemas which should be common across all managed service APIs.

_Note: This list is incomplete._


### Error

[Error](https://github.com/redhat-developer/app-services-api-guidelines/tree/main/spectral#rhoas-error-schema)


### List

[List](https://github.com/redhat-developer/app-services-api-guidelines/tree/main/spectral#rhoas-list-schema)

This would mean enforcing page/total/size pagination on lists. 
