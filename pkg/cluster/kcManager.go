/**
 * Handles specific operations for Kafka Connection resource
 */
package cluster

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"net/http"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
)

var (
	AKCGroup   = "rhoas.redhat.com"
	AKCVersion = "v1alpha1"
)

var AKCRMeta = metav1.TypeMeta{
	Kind:       "KafkaConnection",
	APIVersion: AKCGroup + "/" + AKCVersion,
}

var AKCResource = schema.GroupVersionResource{
	Group:    AKCGroup,
	Version:  AKCVersion,
	Resource: "kafkaconnections",
}

// checks the cluster to see if a KafkaConnection CRD is installed
func IsKCInstalledOnCluster(ctx context.Context, c *KubernetesCluster) (bool, error) {
	namespace, err := c.CurrentNamespace()
	if err != nil {
		return false, err
	}

	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(getKafkaConnectionsAPIURL(namespace)).
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

func CheckIfConnectionsExist(ctx context.Context, c *KubernetesCluster, namespace string, k *kafkamgmtclient.KafkaRequest) error {
	data := c.clientset.
		RESTClient().
		Get().
		AbsPath(getKafkaConnectionsAPIURL(namespace), k.GetName()).
		Do(ctx)

	var status int
	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", c.localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), k.GetName())
	}

	return nil
}

func getKafkaConnectionsAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/kafkaconnections", namespace)
}

func watchForKafkaStatus(c *KubernetesCluster, crName string, namespace string) error {
	c.logger.Info(c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.log.info.wait"))

	w, err := c.dynamicClient.Resource(AKCResource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
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
								c.logger.Info(icon.Success(), c.localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.log.info.success", localize.NewEntry("Name", crName), localize.NewEntry("Namespace", namespace)))

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

func createKCObject(crName string, namespace string, kafkaID string) *KafkaConnection {
	kafkaConnectionCR := &KafkaConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: AKCRMeta,
		Spec: KafkaConnectionSpec{
			KafkaID:               kafkaID,
			AccessTokenSecretName: tokenSecretName,
			Credentials: CredentialsSpec{
				SecretName: serviceAccountSecretName,
			},
		},
	}

	return kafkaConnectionCR
}
