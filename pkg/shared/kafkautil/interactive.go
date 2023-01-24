package kafkautil

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	"github.com/AlecAivazis/survey/v2"
	kafkamgmtclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
)

const (
	queryLimit = "1000"
)

func InteractiveSelect(ctx context.Context, connection connection.Connection, logger logging.Logger, localizer localize.Localizer) (*kafkamgmtclient.KafkaRequest, error) {
	api := connection.API()

	response, httpRes, err := api.KafkaMgmt().GetKafkas(ctx).Size(queryLimit).Execute()
	if httpRes != nil {
		defer func() {
			_ = httpRes.Body.Close()
		}()
	}
	if err != nil {
		return nil, fmt.Errorf("unable to list Kafka instances: %w", err)
	}

	if response.Size == 0 {
		return nil, localizer.MustLocalizeError("kafka.error.interactive.noKafkas")
	}

	kafkas := make([]string, len(response.Items))
	for index := 0; index < len(response.Items); index++ {
		kafkas[index] = *response.Items[index].Name
	}

	prompt := &survey.Select{
		Message:  localizer.MustLocalize("kafka.common.input.instanceName.message"),
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
