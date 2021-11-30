package status

type ServiceStatus = string

// accepted, preparing, provisioning, ready, failed, deprovision, deleting
const (
	StatusAccepted     ServiceStatus = "accepted"
	StatusPreparing    ServiceStatus = "preparing"
	StatusProvisioning ServiceStatus = "provisioning"
	StatusFailed       ServiceStatus = "failed"
	StatusDeprovision  ServiceStatus = "deprovision"
	StatusDeleting     ServiceStatus = "deleting"
)

// IsInstanceCreating returns whether the Kafka instance is still being created
func IsInstanceCreating(status string) bool {
	return status == StatusAccepted || status == StatusPreparing || status == StatusProvisioning
}
