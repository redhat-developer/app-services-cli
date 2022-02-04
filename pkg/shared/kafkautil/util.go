package kafkautil

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

// RegisterNameFlagCompletionFunc adds dynamic completion for the --name flag
func RegisterNameFlagCompletionFunc(cmd *cobra.Command, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var validNames []string
		directive := cobra.ShellCompDirectiveNoSpace

		conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
		if err != nil {
			return validNames, directive
		}

		req := conn.API().KafkaMgmt().GetKafkas(f.Context)
		if toComplete != "" {
			searchQ := "name like " + toComplete + "%"
			req = req.Search(searchQ)
		}
		kafkas, httpRes, err := req.Execute()
		if err != nil {
			return validNames, directive
		}
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		items := kafkas.GetItems()
		for _, kafka := range items {
			validNames = append(validNames, kafka.GetName())
		}

		return validNames, directive
	})
}

// GetCloudProviderCompletionValues returns the list of supported cloud providers for creating a Kafka instance
// This is used in the cmd.RegisterFlagCompletionFunc for dynamic completion of --provider
func GetCloudProviderCompletionValues(f *factory.Factory) (validProviders []string, directive cobra.ShellCompDirective) {
	validProviders = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validProviders, directive
	}

	cloudProviderResponse, _, err := conn.API().KafkaMgmt().GetCloudProviders(f.Context).Execute()
	if err != nil {
		return validProviders, directive
	}

	cloudProviders := cloudProviderResponse.GetItems()
	validProviders = GetEnabledCloudProviderNames(cloudProviders)

	return validProviders, directive
}

// GetCloudProviderCompletionValues returns the list of region IDs for a particular cloud provider
func GetCloudProviderRegionCompletionValues(f *factory.Factory, providerID string) (validRegions []string, directive cobra.ShellCompDirective) {
	validRegions = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	if providerID == "" {
		return
	}

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validRegions, directive
	}

	cloudProviderResponse, _, err := conn.API().
		KafkaMgmt().
		GetCloudProviderRegions(f.Context, providerID).
		Execute()
	if err != nil {
		return validRegions, directive
	}

	cloudProviders := cloudProviderResponse.GetItems()
	validRegions = GetEnabledCloudRegionIDs(cloudProviders)

	return validRegions, directive
}

// GetEnabledCloudProviderNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledCloudProviderNames(cloudProviders []kafkamgmtclient.CloudProvider) []string {
	cloudProviderNames := []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
}

// FindCloudProviderByName finds and returns a cloud provider item from the list by its name
func FindCloudProviderByName(cloudProviders []kafkamgmtclient.CloudProvider, name string) *kafkamgmtclient.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}

// GetEnabledCloudRegionIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledCloudRegionIDs(regions []kafkamgmtclient.CloudRegion) []string {
	var regionIDs []string
	for _, region := range regions {
		if region.GetEnabled() {
			regionIDs = append(regionIDs, region.GetId())
		}
	}
	return regionIDs
}

// FilterValidTopicNameArgs filters topics from the API and returns the names
// This is used in for dynamic completion of topic names
func FilterValidTopicNameArgs(f *factory.Factory, toComplete string) (validNames []string, directive cobra.ShellCompDirective) {
	validNames = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	cfg, err := f.Config.Load()
	if err != nil {
		return validNames, directive
	}

	instanceID, ok := cfg.GetKafkaIdOk()
	if !ok {
		return validNames, directive
	}

	conn, err := f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return validNames, directive
	}

	api, _, err := conn.API().KafkaAdmin(instanceID)
	if err != nil {
		return validNames, directive
	}
	req := api.TopicsApi.GetTopics(f.Context)
	if toComplete != "" {
		req = req.Filter(toComplete)
	}

	topicRes, _, err := req.Execute()
	if err != nil {
		return validNames, directive
	}

	items := topicRes.GetItems()
	for _, topic := range items {
		validNames = append(validNames, topic.GetName())
	}

	return validNames, directive
}

// FilterValidConsumerGroupIDs returns the list of consumer group IDs from the API
func FilterValidConsumerGroupIDs(f *factory.Factory, toComplete string) (validIDs []string, directive cobra.ShellCompDirective) {
	validIDs = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	cfg, err := f.Config.Load()
	if err != nil {
		return validIDs, directive
	}

	instanceID, ok := cfg.GetKafkaIdOk()
	if !ok {
		return validIDs, directive
	}

	conn, err := f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return validIDs, directive
	}

	api, _, err := conn.API().KafkaAdmin(instanceID)
	if err != nil {
		return validIDs, directive
	}
	req := api.GroupsApi.GetConsumerGroups(f.Context)
	if toComplete != "" {
		req = req.GroupIdFilter(toComplete)
	}

	cgRes, _, err := req.Execute()
	if err != nil {
		return validIDs, directive
	}

	items := cgRes.GetItems()
	for _, cg := range items {
		validIDs = append(validIDs, cg.GetGroupId())
	}

	return validIDs, directive
}
