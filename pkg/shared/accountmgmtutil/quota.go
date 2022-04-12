package accountmgmtutil

type QuotaType = string

const (
	QuotaTrialType    QuotaType = "trial"
	QuotaStandardType QuotaType = "standard"
)

type CloudProviderValues = string

const (
	DeveloperType CloudProviderValues = "developer"
	// Deprecated by DeveloperType
	TrialType    CloudProviderValues = "eval"
	StandardType CloudProviderValues = "standard"
)
