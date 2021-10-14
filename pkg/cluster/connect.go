package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kubeclient"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/services"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

// Connect connects a remote Kafka instance to the Kubernetes cluster
func (api *KubernetesClusterAPIImpl) ExecuteConnect(connectOpts *v1alpha.ConnectOperationOptions) error {
	var currentNamespace string
	var err error
	if connectOpts.Namespace != "" {
		currentNamespace = connectOpts.Namespace
	} else {
		currentNamespace, err = api.KubernetesClients.CurrentNamespace()
		if err != nil {
			return kubeclient.TranslatedKubernetesErrors(api.CommandEnvironment, err)
		}
	}
	cliOpts := api.CommandEnvironment

	// Creates abstraction of the service
	currentService, err := api.createServiceInstance(connectOpts.ServiceType)
	if err != nil {
		return err
	}

	serviceDetails, err := currentService.BuildServiceDetails(
		connectOpts.ServiceName, currentNamespace, connectOpts.IgnoreContext)
	if err != nil {
		return err
	}

	if serviceDetails == nil {
		return nil
	}

	// print status
	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.log.info.statusMessage"))

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.statusInfo",
		localize.NewEntry("ServiceType", color.Info(serviceDetails.Type)),
		localize.NewEntry("ServiceName", color.Info(serviceDetails.Name)),
		localize.NewEntry("Namespace", color.Info(currentNamespace)),
		localize.NewEntry("ServiceAccountSecretName", color.Info(constants.ServiceAccountSecretName))))

	if connectOpts.ForceCreationWithoutAsk == false {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: cliOpts.Localizer.MustLocalize("cluster.kubernetes.input.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			cliOpts.Logger.Debug(cliOpts.Localizer.MustLocalize("cluster.kubernetes.log.debug.cancellingOperation"))
			return nil
		}
	}

	// Token with auth for operator to pick
	err = api.createTokenSecretIfNeeded(currentNamespace, connectOpts.OfflineAccessToken)
	if err != nil {
		return kubeclient.TranslatedKubernetesErrors(api.CommandEnvironment, err)
	}

	err = api.createServiceAccountSecretIfNeeded(currentNamespace)
	if err != nil {
		return kubeclient.TranslatedKubernetesErrors(api.CommandEnvironment, err)
	}

	err = api.createCustomResource(serviceDetails, currentNamespace)
	if err != nil {
		return kubeclient.TranslatedKubernetesErrors(api.CommandEnvironment, err)
	}
	return nil
}

// createCustomResource
func (api *KubernetesClusterAPIImpl) createCustomResource(serviceDetails *services.ServiceDetails, namespace string) error {
	crJSON, err := json.Marshal(serviceDetails.KubernetesResource)
	if err != nil {
		return fmt.Errorf("%v: %w", api.CommandEnvironment.Localizer.MustLocalize("cluster.kubernetes.createCR.error.marshalError"), err)
	}

	err = api.KubernetesClients.MakeCRPostRequest(&serviceDetails.GroupMetadata, serviceDetails.Name, crJSON)
	if err != nil {
		return kubeclient.TranslatedKubernetesErrors(api.CommandEnvironment, err)
	}

	api.CommandEnvironment.Logger.Info(
		api.CommandEnvironment.Localizer.MustLocalize("cluster.kubernetes.createCR.log.info.customResourceCreated",
			localize.NewEntry("Resource", serviceDetails.Type),
			localize.NewEntry("Name", serviceDetails.Name)))

	return api.watchForServiceStatus(serviceDetails, namespace)
}

func (c *KubernetesClusterAPIImpl) createTokenSecretIfNeeded(namespace string, accessToken string) error {
	cliOpts := c.CommandEnvironment
	kClients := c.KubernetesClients
	ctx := cliOpts.Context

	_, err := kClients.Clientset.CoreV1().Secrets(namespace).Get(ctx, constants.TokenSecretName, metav1.GetOptions{})
	if err == nil {
		cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.tokensecret.log.info.found"), constants.TokenSecretName)
		return nil
	}

	if accessToken == "" && !cliOpts.IO.CanPrompt() {
		return cliOpts.Localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "token"))
	}

	if accessToken == "" {
		apiTokenInput := &survey.Input{
			Message: cliOpts.Localizer.MustLocalize("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		}
		surveyErr := survey.AskOne(apiTokenInput, &accessToken)
		if surveyErr != nil {
			return surveyErr
		}
	}
	parser := new(jwt.Parser)
	_, _, err = parser.ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	// Create secret type
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.TokenSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"value": accessToken,
		},
	}

	_, err = kClients.Clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	tokenSecretNameTmplEntry := localize.NewEntry("Name", constants.TokenSecretName)
	if err != nil {
		return fmt.Errorf("%v: %w", cliOpts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createFailed", tokenSecretNameTmplEntry), err)
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createSuccess", tokenSecretNameTmplEntry))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesClusterAPIImpl) createServiceAccountSecretIfNeeded(namespace string) error {
	cliOpts := c.CommandEnvironment
	kClients := c.KubernetesClients
	ctx := cliOpts.Context

	_, err := kClients.Clientset.CoreV1().Secrets(namespace).Get(context.TODO(), constants.ServiceAccountSecretName, metav1.GetOptions{})
	if err == nil {
		cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.log.info.exist"))
		return nil
	}

	serviceAcct, err := c.createServiceAccount(ctx, cliOpts)
	if err != nil {
		return err
	}

	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      constants.ServiceAccountSecretName,
			Namespace: namespace,
		},
		StringData: map[string]string{
			"api-id":     serviceAcct.GetClientId(),
			"api-secret": serviceAcct.GetClientSecret(),
		},
	}

	createdSecret, err := kClients.Clientset.CoreV1().Secrets(namespace).Create(cliOpts.Context, secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", cliOpts.Localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.error.createError"), err)
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createSASecret.log.info.createSuccess", localize.NewEntry("Name", createdSecret.Name)))

	return nil
}

// createServiceAccount creates a service account
func (c *KubernetesClusterAPIImpl) createServiceAccount(ctx context.Context, cliOpts *v1alpha.CommandEnvironment) (*kafkamgmtclient.ServiceAccount, error) {
	t := time.Now()

	api := cliOpts.Connection.API()
	serviceAcct := &kafkamgmtclient.ServiceAccountRequest{Name: fmt.Sprintf("rhoascli-%v", t.Unix())}
	req := api.ServiceAccount().CreateServiceAccount(ctx)
	req = req.ServiceAccountRequest(*serviceAcct)
	serviceAcctRes, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %w", cliOpts.Localizer.MustLocalize("cluster.kubernetes.createServiceAccount.error.createError"), err)
	}

	return &serviceAcctRes, nil
}

func (api *KubernetesClusterAPIImpl) watchForServiceStatus(
	serviceDetails *services.ServiceDetails, namespace string) error {
	localizer := api.CommandEnvironment.Localizer
	logger := api.CommandEnvironment.Logger
	logger.Info(localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.log.info.wait",
		localize.NewEntry("Resource", serviceDetails.Type)))

	w, err := api.KubernetesClients.DynamicClient.
		Resource(serviceDetails.GroupMetadata).
		Namespace(namespace).Watch(api.CommandEnvironment.Context, metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", serviceDetails.Name).String(),
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
							return fmt.Errorf(
								localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.format"), typedCondition)
						}
						if typedCondition["type"].(string) == "Finished" {
							if typedCondition["status"].(string) == "False" {
								w.Stop()
								return fmt.Errorf(
									localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.error.status",
										localize.NewEntry("Resource", serviceDetails.Type)),
									typedCondition["message"])
							}
							if typedCondition["status"].(string) == "True" {
								logger.Info(icon.SuccessPrefix(),
									localizer.MustLocalize("cluster.kubernetes.watchForResourceStatus.log.info.success"))

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
			return fmt.Errorf(localizer.MustLocalize("cluster.kubernetes.watchForKafkaStatus.error.timeout"))
		}
	}
}
