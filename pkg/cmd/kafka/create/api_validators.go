package create

import (
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"k8s.io/utils/strings/slices"
)

type ValidatorInput struct {
	provider string
	region   string
	size     string

	userAMSInstanceType *accountmgmtutil.QuotaSpec

	f         *factory.Factory
	constants *remote.DynamicServiceConstants
	conn      connection.Connection
}

var validBillingModels []string = []string{accountmgmtutil.QuotaMarketplaceType, accountmgmtutil.QuotaStandardType, accountmgmtutil.QuotaEvalType}

func (input *ValidatorInput) ValidateProviderAndRegion() error {
	f := input.f
	f.Logger.Debug("Validating provider and region")
	cloudProviders, _, err := input.conn.API().
		KafkaMgmt().
		GetCloudProviders(f.Context).
		Execute()

	if err != nil {
		return err
	}

	var selectedProvider kafkamgmtclient.CloudProvider

	providerNames := make([]string, 0)
	for _, item := range cloudProviders.Items {
		if !item.GetEnabled() {
			continue
		}
		if item.GetId() == input.provider {
			selectedProvider = item
		}
		providerNames = append(providerNames, item.GetId())
	}
	f.Logger.Debug("Validating cloud provider", input.provider, ". Enabled providers: ", providerNames)

	if !selectedProvider.Enabled {
		providers := strings.Join(providerNames, ",")
		providerEntry := localize.NewEntry("Provider", input.provider)
		validProvidersEntry := localize.NewEntry("Providers", providers)
		return f.Localizer.MustLocalizeError("kafka.create.provider.error.invalidProvider", providerEntry, validProvidersEntry)
	}

	return validateProviderRegion(input, selectedProvider)
}

func validateProviderRegion(input *ValidatorInput, selectedProvider kafkamgmtclient.CloudProvider) error {
	f := input.f
	cloudRegion, _, err := input.conn.API().
		KafkaMgmt().
		GetCloudProviderRegions(f.Context, selectedProvider.GetId()).
		Execute()

	if err != nil {
		return err
	}

	var selectedRegion kafkamgmtclient.CloudRegion
	regionNames := make([]string, 0)
	for _, item := range cloudRegion.Items {
		if !item.GetEnabled() {
			continue
		}
		regionNames = append(regionNames, item.GetId())
		if item.GetId() == input.region {
			selectedRegion = item
		}
	}

	if len(regionNames) != 0 {
		f.Logger.Debug("Validating region", input.region, ". Enabled regions: ", regionNames)

		if !selectedRegion.Enabled {
			regionsString := strings.Join(regionNames, ", ")
			regionEntry := localize.NewEntry("Region", input.region)
			validRegionsEntry := localize.NewEntry("Regions", regionsString)
			providerEntry := localize.NewEntry("Provider", input.provider)
			return f.Localizer.MustLocalizeError("kafka.create.region.error.invalidRegion", regionEntry, providerEntry, validRegionsEntry)
		}

		return nil

	}
	f.Logger.Debug("No regions found for provider. Skipping provider validation", input.provider)

	return nil
}

func (input *ValidatorInput) ValidateSize() error {
	// Size is not required
	if input.size == "" {
		return nil
	}

	sizes, err := FetchValidKafkaSizesLabels(input.f, input.provider, input.region, *input.userAMSInstanceType)
	if err != nil {
		return err
	}

	if !slices.Contains(sizes, input.size) {
		return input.f.Localizer.MustLocalizeError("kafka.create.error.invalidSize", localize.NewEntry("ValidSizes", sizes))
	}

	return nil
}

// ValidateBillingModel validates if user provided a supported billing model
func ValidateBillingModel(billingModel string) error {

	if billingModel == "" {
		return nil
	}

	isValid := flagutil.IsValidInput(billingModel, validBillingModels...)

	if isValid {
		return nil
	}

	return flagutil.InvalidValueError("billing-model", billingModel, validBillingModels...)
}
