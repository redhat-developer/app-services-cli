/**
 * Handles specific operations for Kafka Connection resource
 */
package cluster

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/serviceregistry"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

// IsKCInstalledOnCluster checks the cluster to see if a KafkaConnection CRD is installed
func IsKCInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(kafka.GetKafkaConnectionsAPIURL(namespace)).
		Do(ctx)

	if data.Error() == nil {
		return true, nil
	}

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return false, nil
	}

	return true, data.Error()
}

// CheckIfKafkaConnectionExists checks if the Kafka connections exist
func CheckIfKafkaConnectionExists(ctx context.Context, c *KubernetesCluster, namespace string, k string) error {
	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(kafka.GetKafkaConnectionsAPIURL(namespace), k).
		Do(ctx)

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", c.localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), k)
	}

	return nil
}

// CheckIfRegistryConnectionExists checks if the Service registry connection exists
func CheckIfRegistryConnectionExists(ctx context.Context, c *KubernetesCluster, namespace string, registry string) error {
	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(serviceregistry.GetServiceRegistryAPIURL(namespace), registry).
		Do(ctx)

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", c.localizer.MustLocalize("cluster.kubernetes.checkIfServiceRegistryConnectionExist.existError"), registry)
	}

	return nil
}

func getKafkaConnectionsAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/kafkaconnections", namespace)
}

// Encapsulate things like these in their respective packages
func getServiceRegistryAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/serviceregistryconnections", namespace)
}

func watchForKafkaStatus(ctx context.Context, c *KubernetesCluster, crName string, namespace string) error {
	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.log.info.wait"))

	w, err := c.dynamicClient.Resource(kafka.AKCResource).Namespace(namespace).Watch(ctx, metav1.ListOptions{

		FieldSelector: fields.OneTermEqualSelector("metadata.name", crName).String(),
	})
	if err != nil {
		return err
	}

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
							return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.status"), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								c.logger.Info(icon.SuccessPrefix(), c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.log.info.success", localize.NewEntry("Name", crName), localize.NewEntry("Namespace", namespace)))

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
			return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.timeout"))
		}
	}
}

func watchForServiceRegistryStatus(c *KubernetesCluster, crName string, namespace string) error {
	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.log.info.wait"))

	w, err := c.dynamicClient.Resource(serviceregistry.SRCResource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", crName).String(),
	})
	if err != nil {
		return err
	}

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
							return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.error.status"), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.log.info.success", localize.NewEntry("Name", crName), localize.NewEntry("Namespace", namespace)))

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
			return fmt.Errorf(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.timeout"))
		}
	}
}
