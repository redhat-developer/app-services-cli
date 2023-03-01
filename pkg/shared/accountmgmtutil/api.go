package accountmgmtutil

// AMS types
type QuotaType = string

const (
	QuotaTrialType       QuotaType = "trial"
	QuotaStandardType    QuotaType = "standard"
	QuotaMarketplaceType QuotaType = "marketplace"
	QuotaEvalType        QuotaType = "eval"
	QuotaEnterpriseType  QuotaType = "enterprise"
)
