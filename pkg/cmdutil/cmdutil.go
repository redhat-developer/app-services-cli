package cmdutil

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
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
func FilterValidTopicNameArgs(f *factory.Factory, kafkaID string, toComplete string) ([]string, cobra.ShellCompDirective) {
	validNames := []string{}

	conn, err := f.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return validNames, cobra.ShellCompDirectiveError
	}

	api, _, err := conn.API().TopicAdmin(kafkaID)
	if err != nil {
		return validNames, cobra.ShellCompDirectiveError
	}
	topicRes, _, apiErr := api.GetTopicsList(context.Background()).Filter(toComplete).Execute()
	if apiErr.Error() != "" {
		return validNames, cobra.ShellCompDirectiveError
	}

	if topicRes.GetCount() == 0 {
		return validNames, cobra.ShellCompDirectiveError
	}

	items := topicRes.GetItems()
	for _, topic := range items {
		validNames = append(validNames, topic.GetName())
	}

	return validNames, cobra.ShellCompDirectiveDefault
}

// FilterValidKafkaNames filters Kafkas by name from the API and returns the names
// This is used in the cobra.ValidArgsFunction for dynamic completion of topic names
func FilterValidKafkas(f *factory.Factory, searchName string) ([]string, cobra.ShellCompDirective) {
	validNames := []string{}

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validNames, cobra.ShellCompDirectiveError
	}

	req := conn.API().Kafka().ListKafkas(context.Background())
	if searchName != "" {
		req = req.Search(fmt.Sprintf("name+like %v%%", searchName))
	}
	kafkas, _, err := req.Execute()

	if err.Error() != "" {
		return validNames, cobra.ShellCompDirectiveError
	}

	items := kafkas.GetItems()
	for _, kafka := range items {
		validNames = append(validNames, kafka.GetName())
	}

	return validNames, cobra.ShellCompDirectiveDefault
}
