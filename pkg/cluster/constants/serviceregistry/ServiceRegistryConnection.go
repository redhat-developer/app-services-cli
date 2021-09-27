package serviceregistry

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceRegsitryConnectionSpec contains credentials and connection parameters to  Kafka
type ServiceRegsitryConnectionSpec struct {
	AccessTokenSecretName string          `json:"accessTokenSecretName,omitempty"`
	ServiceRegistryId     string          `json:"serviceRegistryId,omitempty"`
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

// ServiceRegsitryConnectionStatus defines the observed state of ServiceRegsitryConnection
type ServiceRegsitryConnectionStatus struct {
	CreatedBy       string              `json:"createdBy,omitempty"`
	Message         string              `json:"message,omitempty"`
	Updated         string              `json:"updated,omitempty"`
	BootstrapServer BootstrapServerSpec `json:"bootstrapServer"`
	RegistryUrl     string              `json:"registryUrl"`
	// Reference to secret name that needs to be fetched
	ServiceAccountSecretName string `json:"ServiceAccountSecretName,omitempty"`
}

// ServiceRegsitryConnection schema
type ServiceRegsitryConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceRegsitryConnectionSpec   `json:"spec,omitempty"`
	Status ServiceRegsitryConnectionStatus `json:"status,omitempty"`
}

// ServiceRegsitryConnectionList contains a list of ServiceRegsitryConnection
type ServiceRegsitryConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRegsitryConnection `json:"items"`
}
