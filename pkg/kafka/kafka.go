package kafka

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafka/mgmt/apiv1"
)

const (
	queryLimit = "1000"
)

func InteractiveSelect(connection connection.Connection, logger logging.Logger) (*kafkamgmtv1.KafkaRequest, error) {
	api := connection.API()

	response, _, err := api.Kafka().ListKafkas(context.Background()).Size(queryLimit).Execute()

	if err != nil {
		return nil, fmt.Errorf("unable to list Kafka instances: %w", err)
	}

	if response.Size == 0 {
		logger.Info("No Kafka instances were found.")
		return nil, nil
	}

	kafkas := []string{}
	for index := 0; index < len(response.Items); index++ {
		kafkas = append(kafkas, *response.Items[index].Name)
	}

	prompt := &survey.Select{
		Message:  "Select Kafka instance to connect:",
		Options:  kafkas,
		PageSize: 10,
	}

	var selectedKafkaIndex int
	err = survey.AskOne(prompt, &selectedKafkaIndex)
	if err != nil {
		return nil, err
	}

	selectedKafka := response.Items[selectedKafkaIndex]

	return &selectedKafka, nil
}
