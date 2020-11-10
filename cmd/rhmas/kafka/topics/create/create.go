package create

import (
	"time"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhmas/kafka/flags"
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var partitions, replicas int32
var configFile string

const (
	Partitions = "partitions"
	Replicas   = "replicas"
	ConfigFile = "config-file"
)

var topicName string

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create topic",
		Long:  "Create topic in the current selected Managed Kafka cluster",
		Run:   createTopic,
	}

	cmd.Flags().StringVarP(&topicName, flags.FlagName, "n", "", "Topic name (required)")
	_ = cmd.MarkFlagRequired(flags.FlagName)
	cmd.Flags().Int32VarP(&partitions, Partitions, "p", 3, "Set number of partitions")
	cmd.Flags().Int32VarP(&replicas, Replicas, "r", 2, "Set number of replicas")
	cmd.Flags().StringVarP(&configFile, ConfigFile, "f", "", "A path to a file containing extra configuration variables. If this option is not supplied, default configurations will be used")

	// TODO define file format etc
	return cmd
}

func createTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "Creating topic " + topicName + " ...")
	// Mimick operation happening by sleeping for a while
	time.Sleep(500 * time.Millisecond)
	fmt.Fprintln(os.Stderr, "Topic " + topicName + " created")
}
