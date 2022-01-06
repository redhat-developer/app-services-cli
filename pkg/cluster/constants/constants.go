package constants

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	DeploymentResource       = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	DeploymentConfigResource = schema.GroupVersionResource{Group: "apps.openshift.io", Version: "v1", Resource: "deploymentconfigs"}
)

// TokenSecretName - name of the access token sercret used for RHOAS operator authentication
const TokenSecretName = "rh-cloud-services-accesstoken"

// ServiceAccountSecretName - name of the service account secret used as service credentials
const ServiceAccountSecretName = "rh-cloud-services-service-account"
