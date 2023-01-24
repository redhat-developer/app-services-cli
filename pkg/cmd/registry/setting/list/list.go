package list

import (
	registryinstanceclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

// settingRow is the details of a Service Registry settings needed to print to a table
type settingRow struct {
	Name string `json:"name" header:"Name"`

	Label string `json:"label,omitempty" header:"Label"`

	Type string `json:"type" header:"Type"`

	Value string `json:"value" header:"Value"`
}

type options struct {
	registryID string

	f *factory.Factory
}

// NewListCommand creates a new command to view a list of settings
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("setting.list.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("setting.list.cmd.description.long"),
		Example: f.Localizer.MustLocalize("setting.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			if opts.registryID != "" {
				return runList(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runList(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	return cmd

}

func runList(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	request := a.AdminApi.ListConfigProperties(opts.f.Context)

	response, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	rows := mapResponseItemsToRows(response)

	opts.f.Logger.Info("")
	dump.Table(opts.f.IOStreams.Out, rows)
	opts.f.Logger.Info("")

	return nil
}

func mapResponseItemsToRows(settings []registryinstanceclient.ConfigurationProperty) []settingRow {
	rows := make([]settingRow, len(settings))

	for i := range settings {
		k := (settings)[i]
		row := settingRow{
			Name:  k.GetName(),
			Value: k.GetValue(),
			Type:  k.GetType(),
			Label: k.GetLabel(),
		}

		rows[i] = row
	}

	return rows
}
