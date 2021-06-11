package cmdutil

import (
	"context"
	"errors"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/redhat-developer/app-services-cli/pkg/cloudprovider/cloudproviderutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/spf13/cobra"
)

// CheckSurveyError checks the error from AlecAivazis/survey
// if the error is from SIGINT, force exit the program quietly
func CheckSurveyError(err error) error {
	if errors.Is(err, terminal.InterruptErr) {
		os.Exit(0)
	} else if err != nil {
		return err
	}

	return nil
}

// FilterValidTopicNameArgs filters topics from the API and returns the names
// This is used in the cobra.ValidArgsFunction for dynamic completion of topic names
func FilterValidTopicNameArgs(f *factory.Factory, toComplete string) (validNames []string, directive cobra.ShellCompDirective) {
	validNames = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	cfg, err := f.Config.Load()
	if err != nil {
		return validNames, directive
	}

	if !cfg.HasKafka() {
		return validNames, directive
	}

	conn, err := f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return validNames, directive
	}

	api, _, err := conn.API().KafkaAdmin(cfg.Services.Kafka.ClusterID)
	if err != nil {
		return validNames, directive
	}
	req := api.GetTopics(context.Background())
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

// FilterValidConsumerGroups returns the list of consumer group IDs from the API
func FilterValidConsumerGroupIDs(f *factory.Factory, toComplete string) (validIDs []string, directive cobra.ShellCompDirective) {
	validIDs = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	cfg, err := f.Config.Load()
	if err != nil {
		return validIDs, directive
	}

	if !cfg.HasKafka() {
		return validIDs, directive
	}

	conn, err := f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return validIDs, directive
	}

	api, _, err := conn.API().KafkaAdmin(cfg.Services.Kafka.ClusterID)
	if err != nil {
		return validIDs, directive
	}
	req := api.GetConsumerGroups(context.Background())

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

// FilterValidKafkaNames filters Kafkas by name from the API and returns the names
// This is used in the cobra.ValidArgsFunction for dynamic completion of Kafka instance names
func FilterValidKafkas(f *factory.Factory, toComplete string) (validNames []string, directive cobra.ShellCompDirective) {
	validNames = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validNames, directive
	}

	req := conn.API().Kafka().GetKafkas(context.Background())
	if toComplete != "" {
		searchQ := "name like " + toComplete + "%"
		req = req.Search(searchQ)
	}
	kafkas, _, err := req.Execute()

	if err != nil {
		return validNames, directive
	}

	items := kafkas.GetItems()
	for _, kafka := range items {
		validNames = append(validNames, kafka.GetName())
	}

	return validNames, directive
}

// FetchCloudProviders returns the list of supported cloud providers for creating a Kafka instance
// This is used in the cmd.RegisterFlagCompletionFunc for dynamic completion of --provider
func FetchCloudProviders(f *factory.Factory) (validProviders []string, directive cobra.ShellCompDirective) {
	validProviders = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validProviders, directive
	}

	cloudProviderResponse, _, err := conn.API().Kafka().GetCloudProviders(context.Background()).Execute()
	if err != nil {
		return validProviders, directive
	}

	cloudProviders := cloudProviderResponse.GetItems()
	validProviders = cloudproviderutil.GetEnabledNames(cloudProviders)

	return validProviders, directive
}
