package kafkacmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// FilterValidTopicNameArgs filters topics from the API and returns the names
// This is used in for dynamic completion of topic names
func FilterValidTopicNameArgs(f *factory.Factory, toComplete string) (validNames []string, directive cobra.ShellCompDirective) {
	validNames = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return validNames, directive
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return validNames, directive
	}

	instanceID := currCtx.KafkaID
	if instanceID == "" {
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

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return validIDs, directive
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return validIDs, directive
	}

	instanceID := currCtx.KafkaID
	if instanceID == "" {
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
