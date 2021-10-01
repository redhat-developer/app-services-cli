package constants

import (
	"github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	DeploymentResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
)

const TokenSecretName = "rh-cloud-services-accesstoken-cli"

const ServiceAccountSecretName = "rh-cloud-services-service-account"

// CreateAppRef creates a reference to identify the application connecting to the backing service operator
func CreateAppRef(appName string) v1alpha1.Application {

	appRef := v1alpha1.Application{
		Ref: v1alpha1.Ref{
			Group:    DeploymentResource.Group,
			Version:  DeploymentResource.Version,
			Resource: DeploymentResource.Resource,
			Name:     appName,
		},
	}

	return appRef
}

// CreateSBObject creates a reference of the object to be bound
func CreateSBObject(bindingName string, ns string, serviceRef *v1alpha1.Service, appRef *v1alpha1.Application) *v1alpha1.ServiceBinding {
	sb := &v1alpha1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bindingName,
			Namespace: ns,
		},
		Spec: v1alpha1.ServiceBindingSpec{
			BindAsFiles: true,
			Services:    []v1alpha1.Service{*serviceRef},
			Application: *appRef,
		},
	}
	sb.SetGroupVersionKind(v1alpha1.GroupVersionKind)

	return sb
}
