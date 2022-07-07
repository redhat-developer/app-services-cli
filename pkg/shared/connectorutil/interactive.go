package connectorutil

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	"github.com/AlecAivazis/survey/v2"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

func InteractiveSelect(ctx context.Context, connection connection.Connection, logger logging.Logger, localizer localize.Localizer) (*connectormgmtclient.Connector, error) {
	api := connection.API().ConnectorsMgmt()

	list, _, err := api.ConnectorsApi.ListConnectors(ctx).Execute()
	if err != nil {
		return nil, err
	}

	if len(list.Items) == 0 {
		logger.Info(localizer.MustLocalize("connector.common.log.info.noConnectorInstances"))
		return nil, nil
	}

	connectorNames := make([]string, len(list.Items))
	for index := 0; index < len(list.Items); index++ {
		connectorNames[index] = list.Items[index].Name
	}

	prompt := &survey.Select{
		Message:  localizer.MustLocalize("connector.common.input.instanceName.message"),
		Options:  connectorNames,
		PageSize: 10,
	}

	var selectedIndex int
	err = survey.AskOne(prompt, &selectedIndex)
	if err != nil {
		return nil, err
	}

	return &list.Items[selectedIndex], nil
}
