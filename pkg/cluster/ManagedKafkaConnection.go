package cluster

// We should use dependency once repo is public

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedKafkaConnectionSpec contains credentials and connection parameters to Managed Kafka
type ManagedKafkaConnectionSpec struct {
	KafkaID     string          `json:"kafkaId,omitempty"`
	Credentials CredentialsSpec `json:"credentials"`
}

// BootstrapServerSpec contains server host information that can be used to connecto the Managed Kafka
type BootstrapServerSpec struct {
	// Host full host to Managed Kafka Service including port
	Host string `json:"host,omitempty"`
}

// CredentialsSpec specification containing various formats of credentials
type CredentialsSpec struct {
	// Reference to secret name that needs to be fetched
	SecretName string `json:"serviceAccountSecretName,omitempty"`
}

// ManagedKafkaConnectionStatus defines the observed state of ManagedKafkaConnection
type ManagedKafkaConnectionStatus struct {
	CreatedBy       string              `json:"createdBy,omitempty"`
	Message         string              `json:"message,omitempty"`
	Updated         string              `json:"updated,omitempty"`
	BootstrapServer BootstrapServerSpec `json:"bootstrapServer"`
	// Reference to secret name that needs to be fetched
	SecretName string `json:"serviceAccountSecretName,omitempty"`
}

// Not working  // +kubebuilder:printcolumn:name="service.binding/host",type="string",JSONPath=".metadata.annotations",description="status of the kind"

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +k8s:openapi-gen=true

// ManagedKafkaConnection schema
type ManagedKafkaConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedKafkaConnectionSpec   `json:"spec,omitempty"`
	Status ManagedKafkaConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:openapi-gen=true

// ManagedKafkaConnectionList contains a list of ManagedKafkaConnection
type ManagedKafkaConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaConnection `json:"items"`
}
