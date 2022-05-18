package serviceregistryutil

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

const (
	queryLimit = 100
)

func InteractiveSelect(ctx context.Context, connection connection.Connection, logger logging.Logger) (*srsmgmtv1.Registry, error) {
	api := connection.API()

	response, httpRes, err := api.ServiceRegistryMgmt().GetRegistries(ctx).Size(queryLimit).Execute()
	if httpRes != nil {
		defer func() {
			_ = httpRes.Body.Close()
		}()
	}
	if err != nil {
		return nil, fmt.Errorf("unable to list Service Registry instances: %w", err)
	}

	if response.Total == 0 {
		logger.Info("No Service Registry instances were found.")
		return nil, nil
	}

	regisries := make([]string, len(response.Items))
	for index := 0; index < len(response.Items); index++ {
		regisries[index] = *response.Items[index].Name
	}

	prompt := &survey.Select{
		Message:  "Select Service Registry instance to connect:",
		Options:  regisries,
		PageSize: 10,
	}

	var selectedRegistryIndex int
	err = survey.AskOne(prompt, &selectedRegistryIndex)
	if err != nil {
		return nil, err
	}

	selectedRegistry := response.Items[selectedRegistryIndex]

	return &selectedRegistry, nil
}
