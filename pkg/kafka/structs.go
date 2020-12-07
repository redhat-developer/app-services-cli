package kafka

// Cluster is the details of a Kafka instance
type Cluster struct {
	ID                  string `json:"id" header:"ID"`
	Name                string `json:"name" header:"Name"`
	Owner               string `json:"owner" header:"Owner"`
	Kind                string `json:"kind"`
	Href                string `json:"href"`
	Status              string `json:"status" header:"Status"`
	CloudProvider       string `json:"cloud_provider" header:"Cloud Provider"`
	Region              string `json:"region" header:"Region"`
	BootstrapServerHost string `json:"bootstrapServerHost"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

// ClusterList contains a list of Kafka instances
type ClusterList struct {
	Kind  string    `json:"kind"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Total int       `json:"total"`
	Items []Cluster `json:"items"`
}
