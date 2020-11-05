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

#### Get the selected Kafka cluster

```shell
rhmas kafka status
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

### Topics

```bash
rhmas kafka topics --help       
Manage Kafka topics for the current selected Managed Kafka Cluster

Usage:
  rhmas kafka topics [command]

Available Commands:
  create      Create topic
  delete      Delete topic
  list        List topics
  update      Update topic
```

#### Create command
```bash
rhmas kafka topics create --help
Create topic in the current selected Managed Kafka cluster

Usage:
  rhmas kafka topics create [flags]

Flags:
  -f, --config-file string   A path to a file containing extra configuration variables. If this option is not supplied, default configurations will be used
  -h, --help                 help for create
  -n, --name string          Topic name (required)
  -p, --partitions int32     Set number of partitions (default 3)
  -r, --replicas int32       Set number of replicas (default 2)
```

#### List command

```bash
rhmas kafka topics list --help  
List all topics in the current selected Managed Kafka cluster

Usage:
  rhmas kafka topics list [flags]

Flags:
  -h, --help            help for list
  -o, --output string   The output format as 'plain-text', 'json', or 'yaml' (default "plain-text")
```

#### Delete command

```bash
rhmas kafka topics delete --help
Delete topic from the current selected Managed Kafka cluster

Usage:
  rhmas kafka topics delete [flags]

Flags:
  -h, --help          help for delete
  -n, --name string   Topic name (required)
```

#### Update command

```bash
rhmas kafka topics update --help
Update topic in the current selected Managed Kafka cluster

Usage:
  rhmas kafka topics update [flags]

Flags:
  -c, --config string   A comma-separated list of configuration to override e.g 'key1=value1,key2=value2'. (required)
  -h, --help            help for update
  -n, --name string     Topic name (required)
```
