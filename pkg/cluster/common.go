package cluster

import (
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

func watchCustomResourceStatus(w watch.Interface, namespace string, crName string, opts Options) error {
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
