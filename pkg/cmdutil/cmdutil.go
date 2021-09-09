package cmdutil

import (
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/spf13/cobra"
)

// FilterValidTopicNameArgs filters topics from the API and returns the names
// This is used in for dynamic completion of topic names
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

func ConvertPageValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 1
	}

	return int32(val)
}

func ConvertSizeValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 10
	}

	return int32(val)
}
