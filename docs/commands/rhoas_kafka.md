## rhoas kafka

Create, view, use, and manage your Kafka instances

### Synopsis

Manage and interact with Kafka instances.

A Kafka instance includes an Apache Kafka cluster, bootstrap server, and the configurations needed to connect to producer and consumer services.

You can create, view, select, and delete Kafka instances.

For each Kafka instance, you can manage ACLs, consumer groups, and topics.


### Examples

```
# Create a Kafka instance
rhoas kafka create --name my-kafka-instance

# View configuration details of a Kafka instance
rhoas kafka describe

# List all Kafka instances
rhoas kafka list

# Create a Kafka topic
rhoas kafka topic create --name mytopic

```

### Options inherited from parent commands

```
  -h, --help      Show help for a command
  -v, --verbose   Enable verbose mode
```

### SEE ALSO

* [rhoas](rhoas.md)	 - RHOAS CLI
* [rhoas kafka acl](rhoas_kafka_acl.md)	 - Manage Kafka ACLs for users and service accounts
* [rhoas kafka billing](rhoas_kafka_billing.md)	 - List Kafka Billing Types
* [rhoas kafka consumer-group](rhoas_kafka_consumer-group.md)	 - Describe, list, and delete consumer groups for the current Kafka instance
* [rhoas kafka create](rhoas_kafka_create.md)	 - Create a Kafka instance
* [rhoas kafka delete](rhoas_kafka_delete.md)	 - Delete a Kafka instance
* [rhoas kafka describe](rhoas_kafka_describe.md)	 - View configuration details of a Kafka instance
* [rhoas kafka list](rhoas_kafka_list.md)	 - List all Kafka instances
* [rhoas kafka providers](rhoas_kafka_providers.md)	 - List Kafka Cloud Providers
* [rhoas kafka topic](rhoas_kafka_topic.md)	 - Create, describe, update, list, and delete topics
* [rhoas kafka update](rhoas_kafka_update.md)	 - Update configuration details for a Kafka instance.
* [rhoas kafka use](rhoas_kafka_use.md)	 - Set the current Kafka instance

