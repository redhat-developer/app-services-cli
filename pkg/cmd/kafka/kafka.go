// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	// "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/use"
)

func NewKafkaCommand(f *factory.Factory) *cobra.Command {
	localizer.LoadMessageFiles("cmd/kafka")

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("kafka.cmd.use"),
		Short: localizer.MustLocalizeFromID("kafka.cmd.shortDescription"),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
		// topic.NewTopicCommand(f),
	)

	return cmd
}
