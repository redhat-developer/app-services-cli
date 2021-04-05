package consumergroup

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/consumergroup/list"
	"github.com/spf13/cobra"
)

// NewConsumerGroupCommand creates a new command sub-group for consumer group operations
func NewConsumerGroupCommand(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("kafka.consumerGroup.cmd.use"),
		Short: localizer.MustLocalizeFromID("kafka.consumerGroup.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("kafka.consumerGroup.cmd.longDescription"),
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		list.NewListCommand(f),
	)

	return cmd
}
