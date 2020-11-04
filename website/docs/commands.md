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

### Apache Kafka Management commands

#### Create Kafka Cluster

Command creates Kafka cluster and automatically setting it up as default cluster that will be used 
by other subcommands.

```
rhmas kafka create --name=test --multi-az="true" --provider=aws --region=eu-west-1
```

**Arguments**:
  --multi-az          Whether Kafka request should be Multi AZ or not
  --name string       Kafka request name
  --provider string   OCM provider ID (default "aws")
  --region string     Region ID (default "eu-west-1")

  -f file:  uses file as input to create cluster
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

### Get details of Kafka cluster

Get details 

```
rhmas kafka get kafka-id

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

#### List Kakfa clusters 

List Kafka clusters in your current account

```
rhmas kafka list
```

**Arguments**:
  page: index (default "1")
  output: Output format to display the Kafka clusters. Choose from "json" or "table" (default "table")
  size: Number of Kafka requests per page (default "100")

**Returns:**

```shell
  ID                            NAME          OWNER               STATUS     CLOUD PROVIDER   REGION     
 ----------------------------- ------------- ------------------- ---------- ---------------- ----------- 
  1iSY6RQ3JKI8Q0OTmjQFd3ocFRg   serviceapi    api_kafka_service   complete   aws              eu-west-1  
  v5Sg6faQ3JKGas4hFd3og45fd31   serviceapi2   api_kafka_service   complete   aws              eu-west-1
```

#### Switch to use managed kafka

```shell
rhmas kafka use kafka-id
```

#### Get credentials for the managed kafka

```
rhmas kafka get-credentials --cluster-name=my-cluster --type=TLS
```

### Apache Kafka specific commands

```shell
rhmas kafka tail kafka-id topic
```

### Community supported commands

```
rhmas kafka create-connector --cluster-name=my-cluster --type=debezium-mysql-connector --configuration=...
```

