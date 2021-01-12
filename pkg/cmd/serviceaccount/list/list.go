package list

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	output string
}

type table struct {
	ID       string `json:"id" header:"ID"`
	Name     string `json:"name" header:"Name"`
	ClientID string `json:"clientID" header:"Client ID"`
}

// NewCreateCommand creates a new command to list service accounts
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all service accounts",
		Example: heredoc.Doc(`
			$ rhoas serviceaccount list
			$ rhoas serviceaccount list -o json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !flagutil.IsValidInput(opts.output, flagutil.AllowedListFormats...) {
				opts.output = "table"
			}
			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "table", "Format to display the service accounts")

	return cmd
}

func runList(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return nil
	}

	client := connection.NewAPIClient()

	a := client.DefaultApi.ListServiceAccounts(context.Background())
	serviceaccounts, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to list service accounts: %w", apiErr)
	}

	if len(*serviceaccounts.Items) == 0 {
		logger.Info("No service accounts")
		return nil
	}

	jsonResponse, _ := json.Marshal(serviceaccounts.Items)
	var tableList []table

	if err = json.Unmarshal(jsonResponse, &tableList); err != nil {
		logger.Infof("Could not unmarshal service accounts into table, defaulting to JSON instead: %v", err)
		return nil
	}

	switch opts.output {
	case "json":
		data, _ := json.MarshalIndent(serviceaccounts, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(os.Stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(serviceaccounts)
		_ = dump.YAML(os.Stdout, data)
	default:
		dump.Table(os.Stdout, tableList)
		logger.Info("")
	}

	return nil
}
