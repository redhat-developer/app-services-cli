---
id: commands
title: Commands Guide
---

## Command structure

### Log In

The first step to use the tool is to log-in with your
Managed Services offline access token which you can get below:

https://cloud.redhat.com/managed-application-services/token

To do that use the `login` command:

```
$ rhmas login --token=eyJ...
```

Alternatively if we already have account we can login directly.
This will redirect us to the login website:

```
$ rhmas login
```

### OpenShift Streams Management commands

#### Create Streams Instance

Command creates Kafka instance and automatically setting it up as default instance that will be used 
by other subcommands.

```
rhmas streams create --name=test --multi-az="true" --provider=aws --region=eu-west-1
```

**Arguments**:
  --multi-az          Whether Kafka request should be Multi AZ or not
  --name string       Kafka request name
  --provider string   OCM provider ID (default "aws")
  --region string     Region ID (default "eu-west-1")

  -f file:  uses file as input to create instance
```
{
  "region": "us-east-1",
  "cloud_provider": "aws",
  "name": "serviceapitest"
}
```

**Returns:**

```json
{
  "value": {
    "id": "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
    "kind": "kafka",
    "href": "/api/managed-services-api/v1/kafkas/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
    "status": "complete",
    "cloud_provider": "aws",
    "multi_az": "false",
    "region": "eu-west-1",
    "owner": "api_kafka_service",
    "name": "serviceapi",
    "bootstrapServerHost": "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
    "created_at": "2020-10-05T12:51:24.053142Z",
    "updated_at": "2020-10-05T12:56:36.362208Z"
  }
}
```

### Get details of streams instance

Get details 

```
rhmas streams get kafka-id --format=json
```

**Arguments**:
-f --format: format of the data (json/yaml/table)

**Returns:**

```json
{
  "value": {
    "id": "1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
    "kind": "kafka",
    "href": "/api/managed-services-api/v1/kafkas/1iSY6RQ3JKI8Q0OTmjQFd3ocFRg",
    "status": "complete",
    "cloud_provider": "aws",
    "multi_az": "false",
    "region": "eu-west-1",
    "owner": "api_kafka_service",
    "name": "serviceapi",
    "bootstrapServerHost": "serviceapi-1isy6rq3jki8q0otmjqfd3ocfrg.apps.ms-bttg0jn170hp.x5u8.s1.devshift.org",
    "created_at": "2020-10-05T12:51:24.053142Z",
    "updated_at": "2020-10-05T12:56:36.362208Z"
  }
}
```

#### List Streams instances 

List Streams instances in your current account

```
rhmas streams list
```

**Arguments**:
  page: index (default "1")
  size: Number of kafka requests per page (default "100")

#### Switch to use managed kafka

```shell
rhmas streams use kafka-id
```

#### Get credentials for the managed kafka

```
rhmas streams get-credentials --cluster-name=my-cluster --type=TLS
```

### OpenShift Streams Kafka specific commands

```shell
rhmas streams tail kafka-id topic
```

### Community supported commands

```
rhmas streams create-connector --cluster-name=my-cluster --type=debezium-mysql-connector --configuration=...
```

