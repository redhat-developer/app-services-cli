// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/consumergroup"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic"
	"github.com/redhat-developer/app-services-cli/pkg/profile"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/update"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/use"
)

func NewKafkaCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kafka",
		Short:   f.Localizer.MustLocalize("kafka.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
		topic.NewTopicCommand(f),
		consumergroup.NewConsumerGroupCommand(f),
	)

	if profile.DevModeEnabled() {
		cmd.AddCommand(update.NewUpdateCommand(f))
	}

	return cmd
}
