package accountmgmtutil

type QuotaType = string

const (
	// Deprecated by QuotaDeveloperType
	QuotaTrialType    QuotaType = "eval"
	QuotaStandardType QuotaType = "standard"
)
