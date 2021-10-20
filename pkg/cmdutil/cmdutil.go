package cmdutil

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
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

// FilterValidConsumerGroups returns the list of consumer group IDs from the API
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

// StringSliceToListStringWithQuotes converts a string slice to a
// comma-separated list with each value in quotes.
// Example: "a", "b", "c"
func StringSliceToListStringWithQuotes(validOptions []string) string {
	var listF string
	for i, val := range validOptions {
		listF += fmt.Sprintf("\"%v\"", val)
		if i < len(validOptions)-1 {
			listF += ", "
		}
	}
	return listF
}
