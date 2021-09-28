package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
)

// ResourceExists checks if the service already exists in the cluster
func ResourceExists(ctx context.Context, c *cluster.KubernetesCluster, path string, serviceName string, opts cluster.Options) error {

	err := c.MakeKubernetesGetRequest(ctx, path, serviceName, opts.Localizer)

	return err

}

// CreateResource creates a new custom resource
func CreateResource(ctx context.Context, c *cluster.KubernetesCluster, path string, serviceName string, namespace string, crJSON []byte, resource schema.GroupVersionResource, opts cluster.Options, errorMessages map[string]string) error {

	err := c.MakeKubernetesPostRequest(ctx, path, serviceName, crJSON)

	if err != nil {
		return err
	}

	// opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.log.info.customResourceCreated", localize.NewEntry("Name", serviceName)))
	opts.Logger.Info(opts.Localizer.MustLocalize(errorMessages["customResourceCreated"], localize.NewEntry("Name", serviceName)))

	w, err := c.DynamicClient.Resource(resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", serviceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, namespace, serviceName, opts, errorMessages)
}

func watchCustomResourceStatus(w watch.Interface, namespace string, crName string, opts cluster.Options, errorMessages map[string]string) error {
	for {
		select {
		case event := <-w.ResultChan():
			if event.Type == watch.Modified {
				unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(event.Object)
				if err != nil {
					return err
				}
				conditions, found, err := unstructured.NestedSlice(unstructuredObj, "status", "conditions")
				if err != nil {
					return err
				}

				if found {
					for _, condition := range conditions {
						typedCondition, ok := condition.(map[string]interface{})
						if !ok {
							return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForConnectionStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(opts.Localizer.MustLocalize(errorMessages["statusError"]), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								opts.Logger.Info(opts.Localizer.MustLocalize(errorMessages["successfullyCreated"], localize.NewEntry("Name", crName), localize.NewEntry("Namespace", namespace)))

								w.Stop()
								return nil
							}
						}
					}
					w.Stop()
				}
			}

		case <-time.After(60 * time.Second):
			w.Stop()
			return fmt.Errorf(opts.Localizer.MustLocalize(errorMessages["timeoutError"]))
		}
	}
}
