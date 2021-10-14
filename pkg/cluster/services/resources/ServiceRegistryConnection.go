package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceRegistryConnectionSpec contains credentials and connection parameters to  Kafka
type ServiceRegistryConnectionSpec struct {
	AccessTokenSecretName string                  `json:"accessTokenSecretName,omitempty"`
	ServiceRegistryId     string                  `json:"serviceRegistryId,omitempty"`
	Credentials           RegistryCredentialsSpec `json:"credentials"`
}

// CredentialsSpec specification containing various formats of credentials
type RegistryCredentialsSpec struct {
	// Reference to secret name that needs to be fetched
	SecretName string `json:"serviceAccountSecretName,omitempty"`
}

// ServiceRegistryConnectionStatus defines the observed state of ServiceRegistryConnection
type ServiceRegistryConnectionStatus struct {
	CreatedBy   string `json:"createdBy,omitempty"`
	Message     string `json:"message,omitempty"`
	Updated     string `json:"updated,omitempty"`
	RegistryUrl string `json:"registryUrl"`
	// Reference to secret name that needs to be fetched
	ServiceAccountSecretName string `json:"ServiceAccountSecretName,omitempty"`
}

// ServiceRegistryConnection schema
type ServiceRegistryConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceRegistryConnectionSpec   `json:"spec,omitempty"`
	Status ServiceRegistryConnectionStatus `json:"status,omitempty"`
}
