package consumergroup

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup/resetoffset"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/spf13/cobra"
)

// NewConsumerGroupCommand creates a new command sub-group for consumer group operations
func NewConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "consumer-group",
		Short:   f.Localizer.MustLocalize("kafka.consumerGroup.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.consumerGroup.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.consumerGroup.cmd.example"),
		Args:    cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		list.NewListConsumerGroupCommand(f),
		delete.NewDeleteConsumerGroupCommand(f),
		describe.NewDescribeConsumerGroupCommand(f),
		resetoffset.NewResetOffsetConsumerGroupCommand(f),
	)

	return cmd
}
