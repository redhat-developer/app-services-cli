package connectorcmdutil

import (
	"fmt"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type TypeSearchQuery struct {
	filters []string
	input   string
}

func NewSearchQuery(input string) *TypeSearchQuery {
	return &TypeSearchQuery{
		input:   input,
		filters: make([]string, 0),
	}
}

func (sq *TypeSearchQuery) Filter(filter string) *TypeSearchQuery {
	sq.filters = append(sq.filters, filter)
	return sq
}

func (sq *TypeSearchQuery) Build() string {
	query := ""

	for i := 0; i < len(sq.filters); i++ {
		query += fmt.Sprintf("%[1]s like %[3]s%[2]s%[3]s", sq.filters[i], sq.input, "%")
		if i+1 < len(sq.filters) {
			query += " or "
		}
	}

	return query
}

func FilterValidTypesArgs(f *factory.Factory, toComplete string) ([]string, cobra.ShellCompDirective) {
	validTypes := []string{}
	directive := cobra.ShellCompDirectiveNoSpace

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return validTypes, directive
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return validTypes, directive
	}

	instanceID := currCtx.KafkaID
	if instanceID == "" {
		return validTypes, directive
	}

	conn, err := f.Connection()
	if err != nil {
		return validTypes, directive
	}

	api := conn.API().ConnectorsMgmt().ConnectorTypesApi
	if err != nil {
		return validTypes, directive
	}

	req := api.GetConnectorTypes(f.Context)

	req = req.Size("500")
	req = req.Page("1")

	types, httpRes, err := req.Execute()
	if httpRes != nil {
		defer func() {
			_ = httpRes.Body.Close()
		}()
	}
	if err != nil {
		return validTypes, directive
	}

	items := types.GetItems()
	for _, connector_type := range items {
		if strings.HasPrefix(connector_type.GetId(), toComplete) {
			validTypes = append(validTypes, connector_type.GetId())
		}
	}

	return validTypes, directive
}
