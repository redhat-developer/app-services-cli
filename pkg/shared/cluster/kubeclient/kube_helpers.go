package kubeclient

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/shared/cluster/v1alpha"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// APIURLForResource returns kubernetes API url based on GroupVersionResource
func APIURLForResource(resource *schema.GroupVersionResource, namespace string) string {
	return fmt.Sprintf("/apis/%v/%v/namespaces/%v/%v",
		resource.Group, resource.Version, namespace, resource.Resource)
}

// Translate kubernetes errors to more human redable format
func TranslatedKubernetesErrors(env *v1alpha.CommandEnvironment, err error) error {
	if errors.IsNotFound(err) {
		return env.Localizer.MustLocalizeError("cluster.common.kube.resourcemissing")
	}

	if errors.IsForbidden(err) {
		return env.Localizer.MustLocalizeError("cluster.common.kube.unauthorized")
	}

	if errors.IsServiceUnavailable(err) ||
		errors.IsServerTimeout(err) || errors.IsTimeout(err) {
		return env.Localizer.MustLocalizeError("cluster.common.kube.timeout")
	}

	return err
}
