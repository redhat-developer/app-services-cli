package cmdutil

import (
	"context"
	"errors"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
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

	conn, err := f.Connection()
	if err != nil {
		return validNames, cobra.ShellCompDirectiveError
	}

	api := conn.API()
	topicRes, _, apiErr := api.TopicAdmin(kafkaID).GetTopicsList(context.Background()).Filter(toComplete).Execute()
	if apiErr.Error() != "" {
		return validNames, cobra.ShellCompDirectiveNoFileComp
	}

	if topicRes.GetCount() == 0 {
		return validNames, cobra.ShellCompDirectiveNoFileComp
	}

	items := topicRes.GetTopics()
	for _, topic := range items {
		validNames = append(validNames, topic.GetName())
	}

	return validNames, cobra.ShellCompDirectiveNoFileComp
}
