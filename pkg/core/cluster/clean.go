package cluster

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/services/resources"
	"github.com/redhat-developer/app-services-cli/pkg/core/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

// ExecuteClean - removes resources created by commands
func (api *KubernetesClusterAPIImpl) ExecuteClean(opts *v1alpha.CleanOperationOptions) error {
	kubeClients := api.KubernetesClients
	if opts.Namespace == "" {
		namespace, err := kubeClients.CurrentNamespace()
		if err != nil {
			return err
		}
		opts.Namespace = namespace
	}
	api.CommandEnvironment.Logger.Info(api.CommandEnvironment.Localizer.MustLocalize("cluster.clean.confirmation",
		&localize.TemplateEntry{Key: "Namespace", Value: opts.Namespace}))
	if !opts.ForceDeleteWithoutAsk {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: api.CommandEnvironment.Localizer.MustLocalize("cluster.kubernetes.input.confirm.message"),
		}
		err := survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			api.CommandEnvironment.Logger.Debug(api.CommandEnvironment.Localizer.MustLocalize("cluster.kubernetes.log.debug.cancellingOperation"))
			return nil
		}
	}

	gracePeriod := int64(10)

	err := kubeClients.Clientset.CoreV1().
		Secrets(opts.Namespace).Delete(api.CommandEnvironment.Context,
		constants.ServiceAccountSecretName, v1.DeleteOptions{
			GracePeriodSeconds: &gracePeriod,
		})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	}
	err = kubeClients.Clientset.CoreV1().
		Secrets(opts.Namespace).Delete(api.CommandEnvironment.Context,
		constants.TokenSecretName, v1.DeleteOptions{
			GracePeriodSeconds: &gracePeriod,
		})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	}

	for _, resource := range resources.AllResources {
		err := kubeClients.DynamicClient.Resource(resource).Namespace(opts.Namespace).
			DeleteCollection(api.CommandEnvironment.Context, v1.DeleteOptions{
				GracePeriodSeconds: &gracePeriod,
			}, v1.ListOptions{
				FieldSelector: fields.Everything().String(),
			})

		if err != nil {
			if !errors.IsNotFound(err) {
				return err
			}
		}

	}

	api.CommandEnvironment.Logger.Info(icon.SuccessPrefix(), api.CommandEnvironment.Localizer.MustLocalize("cluster.clean.success"))
	return nil
}
