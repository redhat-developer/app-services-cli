package svcstatus

type ServiceStatus = string

// accepted, preparing, provisioning, ready, failed, deprovision, deleting
const (
	StatusAccepted     ServiceStatus = "accepted"
	StatusPreparing    ServiceStatus = "preparing"
	StatusProvisioning ServiceStatus = "provisioning"
	StatusReady        ServiceStatus = "ready"
	StatusFailed       ServiceStatus = "failed"
	StatusDeprovision  ServiceStatus = "deprovision"
	StatusDeleting     ServiceStatus = "deleting"
)

func IsCreating(status string) bool {
	return status == StatusAccepted || status == StatusPreparing || status == StatusProvisioning
}
