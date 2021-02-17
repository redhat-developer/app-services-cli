package topic

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topic/update"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewTopicCommand gives commands that manages Kafka topics.
func NewTopicCommand(f *factory.Factory) *cobra.Command {
	localizer.LoadMessageFiles("cmd/kafka/topic")

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
