package cluster

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

func makeRequest(ctx context.Context, clientSet *kubernetes.Clientset, path string, serviceName string, localizer localize.Localizer) error {

	var status int

	data := clientSet.
		RESTClient().
		Get().
		AbsPath(path, serviceName).
		Do(ctx)

	if data.StatusCode(&status); status == http.StatusNotFound {
		return nil
	}

	if data.Error() == nil {
		return fmt.Errorf("%v: %s", localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), serviceName)
	}

	return nil
}

func createResourceRequest(ctx context.Context, clientSet *kubernetes.Clientset, path string, serviceName string, crJSON []byte) error {

	data := clientSet.
		RESTClient().
		Post().
		AbsPath(path, serviceName).
		Body(crJSON).
		Do(ctx)

	if data.Error() != nil {
		return data.Error()
	}

	return nil
}

func watchGeneric(w watch.Interface, namespace string, crName string, opts Options) error {
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
							return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.error.status"), typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								opts.Logger.Info(opts.Localizer.MustLocalize("cluster.kubernetes.watchForRegistryStatus.log.info.success", localize.NewEntry("Name", crName), localize.NewEntry("Namespace", namespace)))

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
			return fmt.Errorf(opts.Localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.timeout"))
		}
	}
}
