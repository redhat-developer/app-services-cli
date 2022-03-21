package accountmgmtutil

type QuotaType = string

const (
	// Deprecated by QuotaDeveloperType
	QuotaTrialType     QuotaType = "eval"
	QuotaDeveloperType QuotaType = "developer"
	QuotaStandardType  QuotaType = "standard"
)
