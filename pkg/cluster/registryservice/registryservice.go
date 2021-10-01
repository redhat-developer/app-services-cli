package registryservice

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/constants/serviceregistry"
	"github.com/redhat-developer/app-services-cli/pkg/cluster/utils"
	registryPkg "github.com/redhat-developer/app-services-cli/pkg/serviceregistry"
	"github.com/redhat-developer/service-binding-operator/apis/binding/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ServiceName = "service-registry"

// RegistryService contains methods to connect and bind Service registry instance to cluster
type RegistryService struct {
	Opts cluster.Options
}

// CustomResourceExists checks if the given ServiceRegistryConnection already exists in cluster
func (r *RegistryService) CustomResourceExists(ctx context.Context, c *cluster.KubernetesCluster, serviceName string) (status int, err error) {

	ns, err := c.CurrentNamespace()
	if err != nil {
		return status, err
	}

	path := serviceregistry.GetServiceRegistryAPIURL(ns)

	status, err = c.ResourceExists(ctx, path, serviceName, r.Opts)

	if status == http.StatusNotFound {
		return status, fmt.Errorf("%v: %s", r.Opts.Localizer.MustLocalize("cluster.kubernetes.checkIfConnectionExist.existError"), serviceName)
	}

	return status, err
}

// CreateCustomResource creates a ServiceRegistryConnection in cluster
func (r *RegistryService) CreateCustomResource(ctx context.Context, c *cluster.KubernetesCluster, serviceID string) error {

	ns, err := c.CurrentNamespace()
	if err != nil {
		return err
	}

	api := r.Opts.Connection.API()

	registryInstance, httpRes, err := registryPkg.GetServiceRegistryByID(ctx, api.ServiceRegistryMgmt(), serviceID)
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return err
	}

	serviceName := registryInstance.GetName()

	serviceRegistryCR := createSRObject(serviceName, ns, serviceID)

	crJSON, err := json.Marshal(serviceRegistryCR)
	if err != nil {
		return fmt.Errorf("%v: %w", r.Opts.Localizer.MustLocalize("cluster.kubernetes.createKafkaCR.error.marshalError"), err)
	}

	resourceOpts := &cluster.CustomResourceOptions{
		CRName:      serviceregistry.RegistryResourceMeta.Kind,
		Resource:    serviceregistry.SRCResource,
		CRJSON:      crJSON,
		ServiceName: serviceName,
		Path:        serviceregistry.GetServiceRegistryAPIURL(ns),
	}

	err = c.CreateResource(ctx, resourceOpts, r.Opts)

	return err
}

// BindCustomConnection binds a ServiceRegistryConnection to specified project
func (r *RegistryService) BindCustomConnection(ctx context.Context, serviceName string, options cluster.ServiceBindingOptions, clients *cluster.KubernetesClients) error {

	serviceRef := createSRCServiceRef(serviceName)

	appRef := constants.CreateAppRef(options.AppName)

	if options.BindingName == "" {
		randomValue := make([]byte, 2)
		_, err := rand.Read(randomValue)
		if err != nil {
			return err
		}
		options.BindingName = fmt.Sprintf("%v-%x", serviceName, randomValue)
	}

	sb := constants.CreateSBObject(options.BindingName, options.Namespace, &serviceRef, &appRef)

	err := utils.CheckIfOperatorIsInstalled(ctx, clients.DynamicClient, options.Namespace)
	if err != nil {
		return fmt.Errorf("%s: %w", r.Opts.Localizer.MustLocalizeError("cluster.serviceBinding.operatorMissing"), err)
	}

	return utils.UseOperatorForBinding(ctx, r.Opts, sb, clients.DynamicClient, options.Namespace)
}

func createSRObject(crName string, namespace string, registryID string) *serviceregistry.ServiceRegistryConnection {
	serviceRegistryCR := &serviceregistry.ServiceRegistryConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: serviceregistry.RegistryResourceMeta,
		Spec: serviceregistry.ServiceRegistryConnectionSpec{
			ServiceRegistryId:     registryID,
			AccessTokenSecretName: constants.TokenSecretName,
			Credentials: serviceregistry.CredentialsSpec{
				SecretName: constants.ServiceAccountSecretName,
			},
		},
	}

	return serviceRegistryCR
}

func createSRCServiceRef(serviceName string) v1alpha1.Service {
	serviceRef := v1alpha1.Service{
		NamespacedRef: v1alpha1.NamespacedRef{
			Ref: v1alpha1.Ref{
				Group:    serviceregistry.SRCResource.Group,
				Version:  serviceregistry.SRCResource.Version,
				Resource: serviceregistry.SRCResource.Resource,
				Name:     serviceName,
			},
		},
	}
	return serviceRef
}
