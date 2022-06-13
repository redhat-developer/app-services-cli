package topic

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/consume"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/produce"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/update"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicCommand gives commands that manages Kafka topics.
func NewTopicCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "topic",
		Short:   f.Localizer.MustLocalize("kafka.topic.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.topic.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.topic.cmd.example"),
	}

	cmd.AddCommand(
		create.NewCreateTopicCommand(f),
		list.NewListTopicCommand(f),
		delete.NewDeleteTopicCommand(f),
		describe.NewDescribeTopicCommand(f),
		update.NewUpdateTopicCommand(f),
		produce.NewProduceTopicCommand(f),
		consume.NewConsumeTopicCommand(f),
	)

	return cmd
}
