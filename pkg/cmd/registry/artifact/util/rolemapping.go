package util

import registryinstanceclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"

func GetRoleLabel(role registryinstanceclient.RoleType) string {
	switch role {
	case registryinstanceclient.ROLETYPE_ADMIN:
		return AdminRole
	case registryinstanceclient.ROLETYPE_DEVELOPER:
		return ManagerRole
	case registryinstanceclient.ROLETYPE_READ_ONLY:
		return ViewerRole
	default:
		return "Unknown"
	}
}

func GetRoleEnum(role string) registryinstanceclient.RoleType {
	switch role {
	case AdminRole:
		return registryinstanceclient.ROLETYPE_ADMIN
	case ManagerRole:
		return registryinstanceclient.ROLETYPE_DEVELOPER
	case ViewerRole:
		return registryinstanceclient.ROLETYPE_READ_ONLY
	default:
		return ""
	}
}
