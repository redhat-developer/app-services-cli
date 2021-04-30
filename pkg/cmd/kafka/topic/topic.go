package topic

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/update"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicCommand gives commands that manages Kafka topics.
func NewTopicCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   f.Localizer.MustLocalize("kafka.topic.cmd.use"),
		Short: f.Localizer.MustLocalize("kafka.topic.cmd.shortDescription"),
		Long:  f.Localizer.MustLocalize("kafka.topic.cmd.longDescription"),
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(f),
		list.NewListTopicCommand(f),
		delete.NewDeleteTopicCommand(f),
		describe.NewDescribeTopicCommand(f),
		update.NewUpdateTopicCommand(f),
	)

	return cmd
}
