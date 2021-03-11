package describe

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	id           string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
}

func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("serviceAccount.describe.cmd.use"),
		Short:   localizer.MustLocalizeFromID("serviceAccount.describe.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("serviceAccount.describe.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("serviceAccount.describe.cmd.example"),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("serviceAccount.describe.flag.id.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", localizer.MustLocalizeFromID("serviceAccount.common.flag.output.description"))

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka().GetServiceAccountById(context.Background(), opts.id)
	res, httpRes, apiErr := a.Execute()

	if apiErr.Error() != "" {
		switch httpRes.StatusCode {
		case 404:
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.common.error.notFoundError",
				TemplateData: map[string]interface{}{
					"ID": opts.id,
				},
			}))
		case 403:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "view",
				},
			}), apiErr)
		case 500:
			return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("serviceAccount.common.error.internalServerError"), apiErr)
		default:
			return apiErr
		}
	}

	switch opts.outputFormat {
	case "yaml", "yml":
		data, _ := yaml.Marshal(res)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		data, _ := json.Marshal(res)
		_ = dump.JSON(opts.IO.Out, data)
	}

	return nil
}
