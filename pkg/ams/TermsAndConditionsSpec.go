package ams

type TermsAndConditionsSpec struct {
	Kafka           ServiceTermsSpec `json:"kafka"`
	ServiceRegistry ServiceTermsSpec `json:"service-registry"`
}

type ServiceTermsSpec struct {
	EventCode         string `json:"EventCode"`
	SiteCode          string `json:"SiteCode"`
	StopOnTermsChange bool   `json:"StopOnTermsChange"`
}
