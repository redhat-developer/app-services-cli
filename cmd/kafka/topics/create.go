package topics

import (
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

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create topic",
		Long:  "Create topic in the current selected Managed Kafka cluster",
		Run:   createTopic,
	}

	cmd.Flags().StringVarP(&topicName, Name, "n", "", "Topic name (required)")
	cmd.MarkFlagRequired(Name)
	cmd.Flags().Int32VarP(&partitions, Partitions, "p", 3, "Set number of partitions")
	cmd.Flags().Int32VarP(&replicas, Replicas, "r", 2, "Set number of replicas")
	cmd.Flags().StringVarP(&configFile, ConfigFile, "f", "", "A path to a file containing extra configuration variables. If this option is not supplied, default configurations will be used")

	// TODO define file format etc
	return cmd
}

func createTopic(cmd *cobra.Command, _ []string) {
	fmt.Fprintln(os.Stderr, "Creating topic " + topicName + " ...")
	doRemoteOperation()
	fmt.Fprintln(os.Stderr, "Topic " + topicName + " created")
}
