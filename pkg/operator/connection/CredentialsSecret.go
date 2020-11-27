package connection

// CredentialsSecret Secret file used for credential purposes
type CredentialsSecret struct {
	// ClientID Represents username in client credentials
	ClientID string `json:"clientID,omitempty"`
	// ClientID Represents password in client credentials
	ClientSecret string `json:"clientSecret,omitempty"`
}
