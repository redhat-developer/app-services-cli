package servicespec

const ServiceRegistryServiceName = "service-registry"

const KafkaServiceName = "kafka"

const NamespaceServiceName = "namespace"

const ConnectorServiceName = "connector"

// All services as labels
var AllServiceLabels = []string{KafkaServiceName, ServiceRegistryServiceName, NamespaceServiceName, ConnectorServiceName}
