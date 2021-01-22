package kafka

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	serviceapiclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
)

func InteractiveSelect(connection connection.Connection, logger logging.Logger) (*serviceapiclient.KafkaRequest, error) {
	api := connection.API()

	response, _, apiErr := api.Kafka.ListKafkas(context.Background()).Execute()

	if apiErr.Error() != "" {
		return nil, fmt.Errorf("Unable to list Kafka instances: %w", apiErr)
	}

	if response.Size == 0 {
		logger.Info("No Kafka instances")
		return nil, nil
	}

	kafkas := []string{}
	for index := 0; index < len(response.Items); index++ {
		kafkas = append(kafkas, *response.Items[index].Name)
	}

	prompt := &survey.Select{
		Message: "Select Kafka cluster to connect",
		Options: kafkas,
	}

	var selectedKafkaIndex int
	err := survey.AskOne(prompt, &selectedKafkaIndex)
	if err != nil {
		return nil, err
	}

	selectedKafka := response.Items[selectedKafkaIndex]

	return &selectedKafka, nil
}
