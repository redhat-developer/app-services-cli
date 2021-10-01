package cluster

import (
	"context"
	"fmt"
	"net/http"
	"time"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/kafkaservice"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/registryservice"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/v1alpha"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/golang-jwt/jwt/v4"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

// nolint:funlen
// Connect connects a remote Kafka instance to the Kubernetes cluster
func Connect(cliOpts v1alpha.CommandEnvironment, cmdOptions v1alpha.ConnectOperationOptions) error {

	var currentNamespace string
	var err error
	if cmdOptions.Namespace != "" {
		currentNamespace = cmdOptions.Namespace
	} else {
		currentNamespace, err = c.CurrentNamespace()
		if err != nil {
			return err
		}
	}

	api := cliOpts.Connection.API()
	// TODO replace connection to service
	var connection CustomConnection

	switch cmdOptions.SelectedServiceType {
	case kafkaservice.ServiceName:
		connection = &kafkaservice.KafkaService{
			Opts: cliOpts,
		}
		_, _, err = kafka.GetKafkaByID(context.Background(), api.Kafka(), cmdOptions.SelectedServiceID)
		if err != nil {
			return err
		}
	case registryservice.ServiceName:
		connection = &registryservice.RegistryService{
			Opts: cliOpts,
		}
		_, _, err = serviceregistry.GetServiceRegistryByID(context.Background(), api.ServiceRegistryMgmt(), cmdOptions.SelectedServiceID)
		if err != nil {
			return err
		}
	}

	// print status
	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.log.info.statusMessage"))

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.statusInfo",
		localize.NewEntry("ServiceType", color.Info(cmdOptions.SelectedServiceType)),
		localize.NewEntry("ServiceID", color.Info(cmdOptions.SelectedServiceID)),
		localize.NewEntry("Namespace", color.Info(currentNamespace)),
		localize.NewEntry("ServiceAccountSecretName", color.Info(constants.ServiceAccountSecretName))))

	if cmdOptions.ForceCreationWithoutAsk == false {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: cliOpts.Localizer.MustLocalize("cluster.kubernetes.connect.input.confirm.message"),
		}
		err = survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			cliOpts.Logger.Debug(cliOpts.Localizer.MustLocalize("cluster.kubernetes.connect.log.debug.cancellingConnect"))
			return nil
		}
	}

	// Token with auth for operator to pick
	err = c.createTokenSecretIfNeeded(ctx, currentNamespace, cmdOptions, cliOpts)
	if err != nil {
		return err
	}

	err = c.createServiceAccountSecretIfNeeded(ctx, currentNamespace, cliOpts)
	if err != nil {
		return err
	}

	status, err := connection.CustomResourceExists(ctx, c, cmdOptions.SelectedServiceID)
	if status != http.StatusNotFound {
		return err
	}

	err = connection.CreateCustomResource(ctx, c, cmdOptions.SelectedServiceID)
	if err != nil {
		return err
	}

	return nil
}

// IsRhoasOperatorAvailableOnCluster checks the cluster for availability of RHOAS and Service Binding operator in cluster
func (c *KubernetesCluster) IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error) {
	installed, err := IsKCInstalledOnCluster(ctx, c)
	if !installed {
		return installed, err
	}

	installed, err = IsSRCInstalledOnCluster(ctx, c)
	if !installed {
		return installed, err
	}

	installed, err = IsSBOInstalledOnCluster(ctx, c)

	return installed, err
}

func (c *KubernetesCluster) createTokenSecretIfNeeded(ctx context.Context, namespace string, connectOpts *ConnectArguments, cliOpts Options) error {
	_, err := c.Clientset.CoreV1().Secrets(namespace).Get(context.TODO(), constants.TokenSecretName, metav1.GetOptions{})
	if err == nil {
		cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.tokensecret.log.info.found"), constants.TokenSecretName)
		return nil
	}

	if connectOpts.OfflineAccessToken == "" && !cliOpts.IO.CanPrompt() {
		return cliOpts.Localizer.MustLocalizeError("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "token"))
	}

	if connectOpts.OfflineAccessToken == "" {
		apiTokenInput := &survey.Input{
			Message: cliOpts.Localizer.MustLocalize("cluster.common.flag.offline.token.description", localize.NewEntry("OfflineTokenURL", build.OfflineTokenURL)),
		}
		surveyErr := survey.AskOne(apiTokenInput, &connectOpts.OfflineAccessToken)
		if surveyErr != nil {
			return err
		}
	}
	parser := new(jwt.Parser)
	_, _, err = parser.ParseUnverified(connectOpts.OfflineAccessToken, jwt.MapClaims{})
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
			"value": connectOpts.OfflineAccessToken,
		},
	}

	_, err = c.Clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	tokenSecretNameTmplEntry := localize.NewEntry("Name", constants.TokenSecretName)
	if err != nil {
		return fmt.Errorf("%v: %w", cliOpts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createFailed", tokenSecretNameTmplEntry), err)
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createTokenSecret.log.info.createSuccess", tokenSecretNameTmplEntry))

	return nil
}

// createSecret creates a new secret to store the SASL/PLAIN credentials from the service account
func (c *KubernetesCluster) createServiceAccountSecretIfNeeded(ctx context.Context, namespace string, cliOpts Options) error {
	_, err := c.Clientset.CoreV1().Secrets(namespace).Get(context.TODO(), constants.ServiceAccountSecretName, metav1.GetOptions{})
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
			"client-id":     serviceAcct.GetClientId(),
			"client-secret": serviceAcct.GetClientSecret(),
		},
	}

	createdSecret, err := c.Clientset.CoreV1().Secrets(namespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%v: %w", cliOpts.Localizer.MustLocalize("cluster.kubernetes.serviceaccountsecret.error.createError"), err)
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createSASecret.log.info.createSuccess", localize.NewEntry("Name", createdSecret.Name)))

	return nil
}

// createServiceAccount creates a service account
func (c *KubernetesCluster) createServiceAccount(ctx context.Context, cliOpts Options) (*kafkamgmtclient.ServiceAccount, error) {
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

func (c *KubernetesCluster) MakeKubernetesGetRequest(ctx context.Context, path string, serviceName string, localizer localize.Localizer) (int, error) {
	var status int

	data := c.Clientset.
		RESTClient().
		Get().
		AbsPath(path, serviceName).
		Do(ctx)

	data.StatusCode(&status)

	if data.Error() != nil {
		return status, data.Error()
	}

	return status, nil
}

func (c *KubernetesCluster) MakeKubernetesPostRequest(ctx context.Context, path string, serviceName string, crJSON []byte) error {

	data := c.Clientset.
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

// CreateResource creates a CustomResource connection in the cluster
func (c *KubernetesCluster) CreateResource(ctx context.Context, resourceOpts *CustomResourceOptions, cliOpts Options) error {

	namespace, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	err = c.MakeKubernetesPostRequest(ctx, resourceOpts.Path, resourceOpts.ServiceName, resourceOpts.CRJSON)

	if err != nil {
		return err
	}

	cliOpts.Logger.Info(cliOpts.Localizer.MustLocalize("cluster.kubernetes.createCR.log.info.customResourceCreated", localize.NewEntry("Resource", resourceOpts.CRName), localize.NewEntry("Name", resourceOpts.ServiceName)))

	w, err := c.DynamicClient.Resource(resourceOpts.Resource).Namespace(namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", resourceOpts.ServiceName).String(),
	})
	if err != nil {
		return err
	}

	return watchCustomResourceStatus(w, cliOpts, resourceOpts.CRName)
}

// ResourceExists checks if a CustomResource connection already exists in the cluster
func (c *KubernetesCluster) ResourceExists(ctx context.Context, path string, serviceName string, cliOpts Options) (int, error) {

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
