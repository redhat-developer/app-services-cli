package kafka

// We should use dependency once repo is public

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KafkaConnectionSpec contains credentials and connection parameters to  Kafka
type KafkaConnectionSpec struct {
	AccessTokenSecretName string          `json:"accessTokenSecretName,omitempty"`
	KafkaID               string          `json:"kafkaId,omitempty"`
	Credentials           CredentialsSpec `json:"credentials"`
}

// BootstrapServerSpec contains server host information that can be used to connecto the  Kafka
type BootstrapServerSpec struct {
	// Host full host to  Kafka Service including port
	Host string `json:"host,omitempty"`
}

// CredentialsSpec specification containing various formats of credentials
type CredentialsSpec struct {
	// Reference to secret name that needs to be fetched
	SecretName string `json:"serviceAccountSecretName,omitempty"`
}

// KafkaConnectionStatus defines the observed state of KafkaConnection
type KafkaConnectionStatus struct {
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

// KafkaConnection schema
type KafkaConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaConnectionSpec   `json:"spec,omitempty"`
	Status KafkaConnectionStatus `json:"status,omitempty"`
}
