package topic

import (
	"github.com/redhat-developer/app-services-cli/internal/localizer"
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
		Use:   localizer.MustLocalizeFromID("kafka.topic.cmd.use"),
		Short: localizer.MustLocalizeFromID("kafka.topic.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("kafka.topic.cmd.longDescription"),
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
