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

// table contains the properties used to
// populate the list of service accounts into a table
type table struct {
	ID          string `json:"id" header:"ID"`
	Name        string `json:"name" header:"Name"`
	ClientID    string `json:"clientID" header:"Client ID"`
	Description string `json:"description,omitempty" header:"Description"`
}

// NewListCommand creates a new command to list service accounts
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List service accounts",
		Long:  "List all service accounts belonging to your organization",
		Example: heredoc.Doc(`
			$ rhoas serviceaccount list
			$ rhoas serviceaccount list -o json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger, err := opts.Logger()
			if err != nil {
				return err
			}

			if !flagutil.IsValidInput(opts.output, flagutil.AllowedListFormats...) {
				logger.Infof("Unknown flag value '%v' for --output. Using table format instead", opts.output)
				opts.output = "plain"
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "plain", fmt.Sprintf("Output format of the results. Choose from %q", flagutil.AllowedListFormats))

	return cmd
}

func runList(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka.ListServiceAccounts(context.Background())
	res, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to list service accounts: %w", apiErr)
	}

	serviceaccounts := res.GetItems()
	if len(serviceaccounts) == 0 {
		logger.Info("No service accounts were found.")
		return nil
	}

	var tableList []table
	if opts.output == "plain" {
		jsonResponse, _ := json.Marshal(serviceaccounts)

		if err = json.Unmarshal(jsonResponse, &tableList); err != nil {
			logger.Infof("Could not unmarshal service accounts into table, defaulting to JSON instead: %v", err)
			opts.output = "json"
		}
	}

	switch opts.output {
	case "json":
		logger.Debug("Outputting service accounts to JSON")
		data, _ := json.MarshalIndent(res, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(os.Stdout, data)
	case "yaml", "yml":
		logger.Debug("Outputting service accounts to YAML")
		data, _ := yaml.Marshal(res)
		_ = dump.YAML(os.Stdout, data)
	default:
		logger.Debug("Outputting service accounts to table")
		dump.Table(os.Stdout, tableList)
	}

	return nil
}
