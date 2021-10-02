package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

// TODO - THIS methods are WRONG - NEED TO remove all of that and bring original code from main.

// CreateResource creates a CustomResource connection in the cluster
func CreateResource(resourceOpts *CustomResourceOptions) error {
	cliOpts := c.CommandEnvironment
	kClients := c.KubernetesClients

	namespace, err := kClients.CurrentNamespace()
	if err != nil {
		return err
	}

	err = c.makeKubernetesPostRequest(ctx, resourceOpts.Path, resourceOpts.ServiceName, resourceOpts.CRJSON)

	if err != nil {
		return err
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createCR.log.info.customResourceCreated", localize.NewEntry("Resource", resourceOpts.CRName), localize.NewEntry("Name", resourceOpts.ServiceName)))

	w, err := kClients.DynamicClient.Resource(resourceOpts.Resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", resourceOpts.ServiceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, cliOpts, resourceOpts.CRName)
}

// ResourceExists checks if a CustomResource connection already exists in the cluster
func (c *KubernetesClusterAPIImpl) ResourceExists(ctx context.Context, path string, serviceName string, cliOpts Options) (int, error) {

	status, err := c.MakeKubernetesGetRequest(ctx, path, serviceName, cliOpts.Localizer)

	return status, err
}

func watchCustomResourceStatus(w watch.Interface, cliOpts Options, crName string) error {
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
							return fmt.Errorf(cliOpts.Localizer.MustLocalize("cluster.kubernetes.watchForConnectionStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(cliOpts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.status", localize.NewEntry("Resource", crName)), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.log.info.success", localize.NewEntry("Resource", crName)))

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
			return fmt.Errorf(cliOpts.Localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.timeout", localize.NewEntry("Resource", crName)))
		}
	}
}
