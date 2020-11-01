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


### OpenShift Streaming Management commands

#### Create Streaming Instance

Command creates Kafka instance and automatically setting it up as default instance that will be used 
by other subcommands
```
rhmas kafka create my-cluster
```

**Arguments**:

 -f file:  uses file as input to create instance
```
{
  "region": "us-east-1",
  "cloud_provider": "aws",
  "name": "serviceapitest"
}
```

 -r --region: cloud provider region (us-east-1)
 -p --provider: cloud provider
 <name>: name of the instance

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

### Get details of streaming instance

Get details 

```
rhmas kafka get --format <kafka_id>
```

**Arguments**:
-f --format: format of the data (json/table)
<kafka_id>  (optional) id of the service or currently selected cluster 


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

#### List Streaming instances 

List Streaming instances in your current account

```
rhmas kafka list
```

#### Switch to use managed kafka

```
rhmas kafka use <kafka_id>
```

#### Get credentials for the managed kafka

```
rhmas kafka get-credentials --cluster-name=my-cluster --type=TLS
```

### OpenShift Streaming Kafka specific commands

rhmas kafka tail <kafka_id> <topic>

### Community supported commands

```
rhmas kafka create-connector --cluster-name=my-cluster --type=debezium-mysql-connector --configuration=...
```

