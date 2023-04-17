package create

import (
	"encoding/json"
	"fmt"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/openshift-cluster/openshiftclustercmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/auth/token"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	"os"
	"os/signal"
	"strconv"
	"time"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
	"github.com/redhat-developer/app-services-cli/pkg/shared/svcstatus"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/error"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

const (
	// FlagProvider is a flag representing an provider ID
	FlagProvider = "provider"
	// FlagRegion is a flag representing an region ID
	FlagRegion = "region"
	// FlagSize is a flag representing an size ID
	FlagSize = "size"
	// FlagMarketPlaceAcctID is a flag representing a marketplace cloud account ID used to purchase the instance
	FlagMarketPlaceAcctID = "marketplace-account-id"
	// FlagMarketPlace is a flag representing marketplace where the instance is purchased on
	FlagMarketPlace = "marketplace"
	// FlagMarketPlace is a flag representing billing model of the instance
	FlagBillingModel  = "billing-model"
	TrialInstanceType = "Trial"
	CloudProvider     = "aws"
)

type options struct {
	name      string
	provider  string
	region    string
	size      string
	clusterId string

	marketplaceAcctId string
	marketplace       string
	billingModel      string

	outputFormat string
	autoUse      bool

	interactive  bool
	wait         bool
	bypassChecks bool
	dryRun       bool

	kfmClusterList          *kafkamgmtclient.EnterpriseClusterList
	selectedCluster         *kafkamgmtclient.EnterpriseCluster
	clusterMap              *map[string]v1.Cluster
	useEnterpriseFlow       bool
	hasLegacyQuota          bool
	useLegacyFlow           bool
	useTrialFlow            bool
	clusterManagementApiUrl string
	accessToken             string

	f *factory.Factory
}

var (
	defaultRegion   = "us-east-1"
	defaultProvider = "aws"
)

// NewCreateCommand creates a new command for creating kafkas.
// nolint: funlen
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("kafka.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name != "" {
				validator := &kafkacmdutil.Validator{
					Localizer:  f.Localizer,
					Connection: f.Connection,
				}
				if err := validator.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if opts.bypassChecks && (opts.marketplace != "" || opts.marketplaceAcctId != "" || opts.billingModel != "") {
				return f.Localizer.MustLocalizeError("kafka.create.error.bypassChecks.marketplace")
			}

			if (opts.marketplace != "") != (opts.marketplaceAcctId != "") {
				return f.Localizer.MustLocalizeError("kafka.create.error.insufficientMarketplaceInfo")
			}

			if opts.billingModel == accountmgmtutil.QuotaStandardType && (opts.marketplaceAcctId != "" || opts.marketplace != "") {
				return f.Localizer.MustLocalizeError("kafka.create.error.standard.invalidFlags")
			}

			if opts.billingModel == accountmgmtutil.QuotaEvalType && (opts.marketplaceAcctId != "" || opts.marketplace != "") {
				return f.Localizer.MustLocalizeError("kafka.create.error.eval.invalidFlags")
			}

			if !f.IOStreams.CanPrompt() && opts.name == "" {
				return f.Localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
			} else if opts.name == "" {
				if opts.provider != "" || opts.region != "" || opts.marketplaceAcctId != "" ||
					opts.marketplace != "" || opts.size != "" || opts.billingModel != "" {
					return f.Localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
				}
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)

	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("kafka.create.flag.name.description"))
	flags.StringVar(&opts.provider, FlagProvider, "", f.Localizer.MustLocalize("kafka.create.flag.cloudProvider.description"))
	flags.StringVar(&opts.region, FlagRegion, "", f.Localizer.MustLocalize("kafka.create.flag.cloudRegion.description"))
	flags.StringVar(&opts.size, FlagSize, "", f.Localizer.MustLocalize("kafka.create.flag.size.description"))
	flags.StringVar(&opts.marketplaceAcctId, FlagMarketPlaceAcctID, "", f.Localizer.MustLocalize("kafka.common.flag.marketplaceId.description"))
	flags.StringVar(&opts.marketplace, FlagMarketPlace, "", f.Localizer.MustLocalize("kafka.common.flag.marketplaceType.description"))
	flags.AddOutput(&opts.outputFormat)
	flags.BoolVar(&opts.autoUse, "use", true, f.Localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	flags.BoolVarP(&opts.wait, "wait", "w", false, f.Localizer.MustLocalize("kafka.create.flag.wait.description"))
	flags.BoolVarP(&opts.dryRun, "dry-run", "", false, f.Localizer.MustLocalize("kafka.create.flag.dryrun.description"))
	flags.StringVar(&opts.billingModel, FlagBillingModel, "", f.Localizer.MustLocalize("kafka.common.flag.billingModel.description"))
	flags.AddBypassTermsCheck(&opts.bypassChecks)
	flags.StringVar(&opts.clusterId, "cluster-id", "", f.Localizer.MustLocalize("kafka.create.flag.clusterId.description"))
	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("kafka.openshiftCluster.registerCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("kafka.openshiftCluster.registercluster.flag.accessToken.description"))

	_ = cmd.RegisterFlagCompletionFunc(FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetCloudProviderCompletionValues(f)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagRegion, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetCloudProviderRegionCompletionValues(f, opts.provider)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagSize, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetKafkaSizeCompletionValues(f, opts.provider, opts.region)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagMarketPlace, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetMarketplaceCompletionValues(f)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagMarketPlaceAcctID, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetMarketplaceAccountCompletionValues(f, opts.marketplace)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagBillingModel, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetBillingModelCompletionValues(f)
	})

	openshiftclustercmdutil.HideClusterMgmtFlags(flags)

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	f := opts.f
	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
	if err != nil {
		return err
	}

	var conn connection.Connection
	if conn, err = f.Connection(); err != nil {
		return err
	}

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return err
	}

	var userQuota *accountmgmtutil.QuotaSpec
	if !opts.bypassChecks {
		f.Logger.Debug("Checking if terms and conditions have been accepted")
		// the user must have accepted the terms and conditions from the provider
		// before they can create a kafka instance
		var termsAccepted bool
		var termsURL string
		termsAccepted, termsURL, err = accountmgmtutil.CheckTermsAccepted(f.Context, &constants.Kafka.Ams, conn)

		if err != nil {
			return err
		}
		if !termsAccepted && termsURL != "" {
			f.Logger.Info(f.Localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
			return nil
		}
	}

	var payload *kafkamgmtclient.KafkaRequestPayload
	if opts.interactive {
		f.Logger.Debug()
		if opts.bypassChecks {
			return f.Localizer.MustLocalizeError("kafka.create.error.noInteractiveMode")
		}

		payload, err = promptKafkaPayload(opts, constants)
		if err != nil {
			return err
		}
	} else {

		if opts.provider == "" {
			opts.provider = defaultProvider
		}

		if opts.region == "" {
			opts.region = defaultRegion
		}

		payload = &kafkamgmtclient.KafkaRequestPayload{
			Name:          opts.name,
			Region:        &opts.region,
			CloudProvider: &opts.provider,
			ClusterId:     *kafkamgmtclient.NewNullableString(&opts.clusterId),
		}

		if !opts.bypassChecks {

			err = ValidateBillingModel(opts.billingModel)
			if err != nil {
				return err
			}

			orgQuotas, newErr := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
			if newErr != nil {
				return newErr
			}

			marketplaceInfo := accountmgmtutil.MarketplaceInfo{
				BillingModel:   opts.billingModel,
				Provider:       opts.marketplace,
				CloudAccountID: opts.marketplaceAcctId,
			}

			if opts.marketplace != "" && opts.marketplace != accountmgmtutil.RedHatMarketPlace {
				if opts.marketplace != opts.provider {
					return opts.f.Localizer.MustLocalizeError("kafka.create.provider.error.unsupportedMarketplace")
				}
			}

			userQuota, err = accountmgmtutil.SelectQuotaForUser(f, orgQuotas, marketplaceInfo, opts.provider)
			if err != nil {
				return err
			}

			userQuotaJSON, marshalErr := json.MarshalIndent(userQuota, "", "  ")
			if marshalErr != nil {
				f.Logger.Debug(marshalErr)
			} else {
				f.Logger.Debug("Selected Quota object:")
				f.Logger.Debug(string(userQuotaJSON))
			}

			if userQuota.BillingModel == accountmgmtutil.QuotaMarketplaceType && userQuota.CloudAccounts != nil {
				payload.Marketplace = kafkamgmtclient.NullableString{}
				payload.Marketplace.Set((*userQuota.CloudAccounts)[0].CloudProviderId)
				payload.BillingCloudAccountId = kafkamgmtclient.NullableString{}
				payload.BillingCloudAccountId.Set((*userQuota.CloudAccounts)[0].CloudAccountId)
			}

			if opts.billingModel != "" {
				payload.BillingModel = kafkamgmtclient.NullableString{}
				payload.BillingModel.Set(&opts.billingModel)
			}

			validator := ValidatorInput{
				provider:            opts.provider,
				region:              opts.region,
				size:                opts.size,
				userAMSInstanceType: userQuota,
				f:                   f,
				constants:           constants,
				conn:                conn,
			}
			err1 := validator.ValidateProviderAndRegion()
			if err1 != nil {
				return err1
			}

			err1 = validator.ValidateSize()
			if err1 != nil {
				return err1
			}
			if opts.size != "" {
				sizes, err1 := FetchValidKafkaSizes(opts.f, opts.provider, opts.region, *userQuota)
				if err1 != nil {
					return err1
				}
				printSizeWarningIfNeeded(opts.f, opts.size, sizes)
				payload.SetPlan(mapAmsTypeToBackendType(userQuota) + "." + opts.size)
			}
		}

	}

	f.Logger.Debug("Creating kafka instance", payload.Name)
	data, marshalErr := json.MarshalIndent(payload, "", "  ")
	if marshalErr != nil {
		f.Logger.Debug(marshalErr)
	} else {
		f.Logger.Debug(string(data))
	}

	if opts.dryRun {
		f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.dryRun.success"))
		return nil
	}

	api := conn.API()

	a := api.KafkaMgmt().CreateKafka(f.Context)
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := kafkamgmtv1errors.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case kafkamgmtv1errors.ERROR_120:
			// For standard instances
			return f.Localizer.MustLocalizeError("kafka.create.error.quota.exceeded")
		case kafkamgmtv1errors.ERROR_24:
			// For dev instances
			return f.Localizer.MustLocalizeError("kafka.create.error.instance.limit")
		case kafkamgmtv1errors.ERROR_36:
			return f.Localizer.MustLocalizeError("kafka.create.error.conflictError", localize.NewEntry("Name", payload.Name))
		case kafkamgmtv1errors.ERROR_41:
			return f.Localizer.MustLocalizeError("kafka.create.error.notsupported", localize.NewEntry("Name", payload.Name))
		case kafkamgmtv1errors.ERROR_42:
			return f.Localizer.MustLocalizeError("kafka.create.error.plan.notsupported", localize.NewEntry("Plan", payload.Plan))
		case kafkamgmtv1errors.ERROR_43:
			return f.Localizer.MustLocalizeError("kafka.create.error.billing.invalid", localize.NewEntry("Billing", payload.BillingCloudAccountId))
		}
	}

	if err != nil {
		return err
	}

	if opts.autoUse {
		f.Logger.Debug("Auto-use is set, updating the current instance")
		currCtx.KafkaID = response.GetId()
		svcContext.Contexts[svcContext.CurrentContext] = *currCtx

		if err = f.ServiceContext.Save(svcContext); err != nil {
			return fmt.Errorf("%v: %w", f.Localizer.MustLocalize("kafka.common.error.couldNotUseKafka"), err)
		}
	} else {
		f.Logger.Debug("Auto-use is not set, skipping updating the current instance")
	}

	nameTemplateEntry := localize.NewEntry("Name", response.GetName())

	if opts.wait {
		f.Logger.Debug("--wait flag is enabled, waiting for Kafka to finish creating")
		s := spinner.New(f.IOStreams.ErrOut, f.Localizer)
		s.SetLocalizedSuffix("kafka.create.log.info.creatingKafka", nameTemplateEntry)
		s.Start()

		// when there is a SIGINT, display a message informing the user that this does not cancel the creation
		// and that it is being created in the background
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				f.Logger.Info()
				f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.creatingKafkaSyncSigint"))
				os.Exit(0)
			}
		}()

		for svcstatus.IsInstanceCreating(response.GetStatus()) {
			time.Sleep(cmdutil.DefaultPollTime)

			response, httpRes, err = api.KafkaMgmt().GetKafkaById(f.Context, response.GetId()).Execute()
			if err != nil {
				return err
			}
			defer httpRes.Body.Close()
			f.Logger.Debug("Checking Kafka status:", response.GetStatus())

			s.SetLocalizedSuffix("kafka.create.log.info.creationInProgress",
				localize.NewEntry("Name", response.GetName()),
				localize.NewEntry("Status", color.Info(response.GetStatus())),
			)

		}
		s.Stop()
		f.Logger.Info()
		f.Logger.Info(icon.SuccessPrefix(), f.Localizer.MustLocalize("kafka.create.info.successSync", nameTemplateEntry))
	}

	if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	if !opts.wait {
		f.Logger.Info()
		f.Logger.Info(f.Localizer.MustLocalize("kafka.create.info.successAsync", nameTemplateEntry))
	}

	return nil
}

// set type to store the answers from the prompt with defaults
type promptAnswers struct {
	Name              string
	Size              string
	Region            string
	CloudProvider     string
	BillingModel      string
	MarketplaceAcctID string
	Marketplace       string
}

func setEnterpriseClusterList(opts *options) (*kafkamgmtclient.EnterpriseClusterList, *map[string]v1.Cluster, error) {
	// Get the list of enterprise clusters in the users organization
	kfmClusterList, response, err := kafkautil.ListEnterpriseClusters(opts.f)
	if err != nil {
		if response.StatusCode == 403 {
			emptyClusterMap := make(map[string]v1.Cluster)
			return &kafkamgmtclient.EnterpriseClusterList{}, &emptyClusterMap, nil
		}

		return nil, nil, fmt.Errorf("%v, %w", response.Status, err)
	}

	clusterMap, err := getClusterNameMap(opts, kfmClusterList)
	if err != nil {
		return nil, nil, err
	}

	return kfmClusterList, clusterMap, nil
}

func selectEnterpriseOrRHInfraPrompt(opts *options) error {
	listOfOptions := []string{
		opts.f.Localizer.MustLocalize("kafka.create.input.cluster.option.enterprise"),
		opts.f.Localizer.MustLocalize("kafka.create.input.cluster.option.rhinfr"),
	}

	promptForCluster := &survey.Select{
		Message: opts.f.Localizer.MustLocalize("kafka.create.input.cluster.message"),
		Help:    opts.f.Localizer.MustLocalize("kafka.create.input.cluster.help"),
		Options: listOfOptions,
	}

	var index int
	err := survey.AskOne(promptForCluster, &index)
	if err != nil {
		return err
	}

	opts.useEnterpriseFlow = index == 0
	opts.useLegacyFlow = index == 1

	return nil
}

func checkForLegacyQuota(opts *options, orgQuotas *accountmgmtutil.OrgQuotas) {
	// Check if the user has enterprise quota
	for _, quota := range orgQuotas.EnterpriseQuotas {
		if quota.Quota > 0 {
			opts.useEnterpriseFlow = true
		}
	}

	// Check if the user has a legacy quota
	for _, quota := range orgQuotas.StandardQuotas {
		if quota.Quota > 0 {
			opts.hasLegacyQuota = true
			return
		}
	}
	for _, quota := range orgQuotas.MarketplaceQuotas {
		if quota.Quota > 0 {
			opts.hasLegacyQuota = true
			return
		}
	}
	for _, quota := range orgQuotas.EvalQuotas {
		if quota.Quota > 0 {
			opts.hasLegacyQuota = true
			return
		}
	}

	// use trial flow if user has no other quota
	if !opts.useEnterpriseFlow {
		opts.useTrialFlow = true
	}
	return
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
// nolint:funlen
func promptKafkaPayload(opts *options, constants *remote.DynamicServiceConstants) (*kafkamgmtclient.KafkaRequestPayload, error) {
	f := opts.f

	// getting org quotas
	orgQuota, err := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
	if err != nil {
		return nil, err
	}

	// check if org has legacy (non-hybrid) quota
	checkForLegacyQuota(opts, orgQuota)

	var enterpriseQuota accountmgmtutil.QuotaSpec

	answers := &promptAnswers{}

	// Message the user with a link to get enterprise quota if they don't have any
	if !opts.useEnterpriseFlow {
		f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.create.info.enterpriseQuota"))
	}

	if opts.useTrialFlow {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.create.usingTrialInstance"))
	}

	answers, err = promptForKafkaName(f, answers)
	if err != nil {
		return nil, err
	}

	opts.kfmClusterList = &kafkamgmtclient.EnterpriseClusterList{}

	if opts.useEnterpriseFlow {
		// Get the list of enterprise clusters in the users organization if there are any, creates a map of cluster ids
		// to names to include names in the prompt
		kfmClusterList, clusterMap, err2 := setEnterpriseClusterList(opts)
		if err2 != nil {
			return nil, err2
		}
		opts.kfmClusterList = kfmClusterList
		opts.clusterMap = clusterMap
	}

	// If there are enterprise clusters in the user's organization, prompt them to select a flow (enterprise or legacy)
	// using the interactive prompt. If there are no enterprise clusters, the user must use the legacy flow.
	// Default to enterprise flow
	err = determineFlowFromQuota(opts)
	if err != nil {
		return nil, err
	}

	if opts.useEnterpriseFlow {
		if len(orgQuota.EnterpriseQuotas) < 1 {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.error.noStandardQuota")
		}

		enterpriseQuota = orgQuota.EnterpriseQuotas[0]

		// there is no quota left to use enterprise
		if enterpriseQuota.Quota == 0 {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.error.noQuotaLeft")
		}

		index, err2 := selectClusterPrompt(opts)
		if err2 != nil {
			return nil, err2
		}

		cluster, err2 := getClusterDetails(opts.f, opts.kfmClusterList.GetItems()[index].GetId())
		if err2 != nil {
			return nil, err2
		}

		if cluster.GetCapacityInformation().RemainingKafkaStreamingUnits == 0 {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.error.noCapacityInSelectedCluster")
		}

		opts.selectedCluster = cluster
	}

	if opts.useTrialFlow {
		hasOtherTrialKafka, err2 := doesUserHaveTrialInstances(opts)
		if err2 != nil {
			return nil, err2
		}

		if hasOtherTrialKafka {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.error.trialInstanceAlreadyExists")
		}
	}

	// If the user is not using an enterprise cluster, we set the cloud provider to aws as this is currently the only
	// cloud provider that we support
	if opts.useLegacyFlow {
		answers.CloudProvider = CloudProvider
	}

	// gets the billing model for the kafka, if the user has more than one billing model, a prompt is shown to select one
	answers, err = getBillingModel(opts, orgQuota, answers)
	if err != nil {
		return nil, err
	}

	// if billing model is marketplace, prompt for marketplace provider and account id
	if answers.BillingModel == accountmgmtutil.QuotaMarketplaceType {
		answers, err = marketplaceQuotaPrompt(orgQuota, answers, f)
		if err != nil {
			return nil, err
		}
	}

	var sizes []kafkamgmtclient.SupportedKafkaSize
	// nolint:staticcheck
	payload := &kafkamgmtclient.KafkaRequestPayload{}

	if opts.useEnterpriseFlow {
		/*
			if using dedicated cluster option then get the supported sizes for this cluster
			based on the sizes it supports the user to select the one they want to use
			then check if the size uses <= the remaining streaming units it takes up
			on the cluster
		*/
		supportedInstanceTypes := opts.selectedCluster.GetSupportedInstanceTypes()
		supportedKafkaTypes := supportedInstanceTypes.GetInstanceTypes()

		// right now enterprise clusters only support one type of kafka, so we are hard
		// checking this, this assumption may change in the future
		if len(supportedKafkaTypes) != 1 {
			return nil, fmt.Errorf("expected one supported kafka instance type")
		}

		sizes = supportedKafkaTypes[0].Sizes

		index, err := promptUserForSizes(opts.f, &sizes)
		if err != nil {
			return nil, err
		}

		selectedSize := sizes[index]

		if selectedSize.GetCapacityConsumed() > opts.selectedCluster.CapacityInformation.RemainingKafkaStreamingUnits {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.cluster.error.noEnoughCapacity", localize.NewEntry("Size", selectedSize.GetId()))
		}

		answers.Size = selectedSize.GetId()

		if selectedSize.GetQuotaConsumed() > int32(enterpriseQuota.Quota) {
			return nil, opts.f.Localizer.MustLocalizeError("kafka.create.error.notEnoughQuota", localize.NewEntry("Size", selectedSize.GetId()))
		}

		enterprise := "enterprise"
		billingModelEnterprise := CreateNullableString(&enterprise)
		clusterIdSNS := CreateNullableString(opts.selectedCluster.ClusterId)

		payload = &kafkamgmtclient.KafkaRequestPayload{
			Name:                  answers.Name,
			BillingModel:          billingModelEnterprise,
			BillingCloudAccountId: CreateNullableString(nil),
			Marketplace:           CreateNullableString(nil),
			ClusterId:             clusterIdSNS,
		}

		// enterprise quota spec, kfm will use the default, this should be returned by the `select quota` call
		// we know someone has enterprise quota here because they have been able to select the
		// cluster to add a kafka instance too
		payload.SetPlan(mapAmsTypeToBackendType(&orgQuota.EnterpriseQuotas[0]) + "." + answers.Size)
	} else {

		// This flow is for the provisioning of a standard kafka instance
		marketplaceInfo := accountmgmtutil.MarketplaceInfo{
			BillingModel:   answers.BillingModel,
			Provider:       answers.Marketplace,
			CloudAccountID: answers.MarketplaceAcctID,
		}

		userQuota, err := accountmgmtutil.SelectQuotaForUser(f, orgQuota, marketplaceInfo, answers.CloudProvider)
		if err != nil {
			return nil, err
		}
		userQuotaJSON, marshalErr := json.MarshalIndent(userQuota, "", "  ")
		if marshalErr != nil {
			f.Logger.Debug(marshalErr)
		} else {
			f.Logger.Debug("Selected Quota object:")
			f.Logger.Debug(string(userQuotaJSON))
		}

		regionIDs, err := GetEnabledCloudRegionIDs(opts.f, answers.CloudProvider, userQuota)
		if err != nil {
			return nil, err
		}

		// selecting the region, if there is only one region available, we skip the prompt
		switch {
		case len(regionIDs) == 0:
			return nil, f.Localizer.MustLocalizeError("kafka.create.error.noRegionSupported")
		case len(regionIDs) == 1:
			answers.Region = regionIDs[0]
		default:
			err = promptForCloudRegion(f, regionIDs, answers)
			if err != nil {
				return nil, err
			}
		}
		sizes, err = FetchValidKafkaSizes(opts.f, answers.CloudProvider, answers.Region, *userQuota)
		if err != nil {
			return nil, err
		}

		index, err := promptUserForSizes(opts.f, &sizes)
		if err != nil {
			return nil, err
		}

		answers.Size = sizes[index].GetId()

		accountIDNullable := kafkamgmtclient.NullableString{}
		marketplaceProviderNullable := kafkamgmtclient.NullableString{}
		billingNullable := kafkamgmtclient.NullableString{}

		if answers.BillingModel != "" {
			billingNullable.Set(&answers.BillingModel)
		}

		if answers.Marketplace != "" {
			marketplaceProviderNullable.Set(&answers.Marketplace)
		}

		if answers.MarketplaceAcctID != "" {
			accountIDNullable.Set(&answers.MarketplaceAcctID)
		}
		payload = &kafkamgmtclient.KafkaRequestPayload{
			Name:                  answers.Name,
			Region:                &answers.Region,
			CloudProvider:         &answers.CloudProvider,
			BillingModel:          billingNullable,
			BillingCloudAccountId: accountIDNullable,
			Marketplace:           marketplaceProviderNullable,
		}
		printSizeWarningIfNeeded(opts.f, answers.Size, sizes)
		payload.SetPlan(mapAmsTypeToBackendType(userQuota) + "." + answers.Size)

	}
	return payload, nil
}

func promptForCloudRegion(f *factory.Factory, regionIDs []string, answers *promptAnswers) error {
	regionPrompt := &survey.Select{
		Message: f.Localizer.MustLocalize("kafka.create.input.cloudRegion.message"),
		Options: regionIDs,
		Help:    f.Localizer.MustLocalize("kafka.create.input.cloudRegion.help"),
	}

	err := survey.AskOne(regionPrompt, &answers.Region)
	if err != nil {
		return err
	}
	return nil
}

func getBillingModel(opts *options, orgQuota *accountmgmtutil.OrgQuotas, answers *promptAnswers) (*promptAnswers, error) {
	availableBillingModels := FetchSupportedBillingModels(orgQuota, answers.CloudProvider)

	if len(availableBillingModels) == 0 && len(orgQuota.MarketplaceQuotas) > 0 {
		return nil, opts.f.Localizer.MustLocalizeError("kafka.create.provider.error.noStandardInstancesAvailable")
	}

	// prompting for billing model if there are more than one available, otherwise using the only one available
	if len(availableBillingModels) > 0 {
		if len(availableBillingModels) == 1 {
			answers.BillingModel = availableBillingModels[0]
		} else {
			billingModelPrompt := &survey.Select{
				Message: opts.f.Localizer.MustLocalize("kafka.create.input.billingModel.message"),
				Options: availableBillingModels,
			}
			err := survey.AskOne(billingModelPrompt, &answers.BillingModel)
			if err != nil {
				return nil, err
			}
		}
	}
	return answers, nil
}

func determineFlowFromQuota(opts *options) error {
	switch {
	case len(opts.kfmClusterList.Items) > 0 && opts.hasLegacyQuota:
		err := selectEnterpriseOrRHInfraPrompt(opts)
		if err != nil {
			return err
		}
		return nil
	case opts.hasLegacyQuota:
		opts.useLegacyFlow = true
		return nil
	case opts.useTrialFlow:
		opts.useLegacyFlow = true
		return nil
	default:
		opts.useEnterpriseFlow = true
		return nil
	}
}

func doesUserHaveTrialInstances(opts *options) (bool, error) {
	conn, err := opts.f.Connection()
	if err != nil {
		return false, err
	}

	api := conn.API()

	cfg, err := opts.f.Config.Load()
	if err != nil {
		return false, err
	}

	userName, ok := token.GetUsername(cfg.AccessToken)
	if !ok {
		return false, opts.f.Localizer.MustLocalizeError("kafka.create.error.cannotFindUserDetails")
	}

	searchQuery := fmt.Sprintf("owner = %v", userName)

	list, _, err := api.KafkaMgmt().GetKafkas(opts.f.Context).
		Page(strconv.Itoa(1)).
		Size(strconv.Itoa(1000)).
		Search(searchQuery).
		Execute()

	if err != nil {
		return false, err
	}

	for i := int32(0); i < list.GetSize(); i++ {
		kafka := list.Items[i]
		if kafka.GetInstanceTypeName() == TrialInstanceType {
			return true, nil
		}
	}

	return false, nil
}

func promptForKafkaName(f *factory.Factory, answers *promptAnswers) (*promptAnswers, error) {
	validator := &kafkacmdutil.Validator{
		Localizer:  f.Localizer,
		Connection: f.Connection,
	}

	// prompt for kafka name
	promptName := &survey.Input{
		Message: f.Localizer.MustLocalize("kafka.create.input.name.message"),
		Help:    f.Localizer.MustLocalize("kafka.create.input.name.help"),
	}

	err := survey.AskOne(promptName, &answers.Name, survey.WithValidator(validator.ValidateName), survey.WithValidator(validator.ValidateNameIsAvailable))
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func marketplaceQuotaPrompt(orgQuota *accountmgmtutil.OrgQuotas, answers *promptAnswers, f *factory.Factory) (*promptAnswers, error) {
	validMarketPlaces := FetchValidMarketplaces(orgQuota.MarketplaceQuotas, answers.CloudProvider)
	if len(validMarketPlaces) == 1 {
		answers.Marketplace = validMarketPlaces[0]
	} else {
		marketplacePrompt := &survey.Select{
			Message: f.Localizer.MustLocalize("kafka.create.input.marketplace.message"),
			Options: validMarketPlaces,
		}
		err := survey.AskOne(marketplacePrompt, &answers.Marketplace)
		if err != nil {
			return nil, err
		}
	}

	if len(validMarketPlaces) > 0 {

		validMarketplaceAcctIDs := FetchValidMarketplaceAccounts(orgQuota.MarketplaceQuotas, answers.Marketplace)

		if len(validMarketplaceAcctIDs) == 1 {
			answers.MarketplaceAcctID = validMarketplaceAcctIDs[0]
		} else {
			marketplaceAccountPrompt := &survey.Select{
				Message: f.Localizer.MustLocalize("kafka.create.input.marketplaceAccountID.message"),
				Options: validMarketplaceAcctIDs,
			}
			err := survey.AskOne(marketplaceAccountPrompt, &answers.MarketplaceAcctID)
			if err != nil {
				return nil, err
			}
		}
	}
	return answers, nil
}

// This may need to altered as the `name` are mutable on ocm side
func getClusterNameMap(opts *options, clusterList *kafkamgmtclient.EnterpriseClusterList) (*map[string]v1.Cluster, error) {
	//	for each cluster in the list, get the name from ocm and add it to the cluster list
	str := kafkautil.CreateClusterSearchStringFromKafkaList(clusterList)
	ocmClusterList, err := clustermgmt.GetClusterListWithSearchParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, str, int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), len(clusterList.Items))
	if err != nil {
		return nil, err
	}
	clusterMap := make(map[string]v1.Cluster)
	for _, cluster := range ocmClusterList.Slice() {
		clusterMap[cluster.ID()] = *cluster
	}
	return &clusterMap, nil

}

func selectClusterPrompt(opts *options) (int, error) {
	promptOptions, err := openshiftclustercmdutil.CreatePromptOptionsFromClusters(opts.kfmClusterList, opts.clusterMap)
	if err != nil {
		return 0, err
	}
	promptForCluster := &survey.Select{
		Message: opts.f.Localizer.MustLocalize("kafka.create.input.cluster.selectClusterMessage"),
		Options: *promptOptions,
	}

	var index int
	err = survey.AskOne(promptForCluster, &index)
	if err != nil {
		return 0, err
	}

	return index, nil
}

func printSizeWarningIfNeeded(f *factory.Factory, selectedSize string, sizes []kafkamgmtclient.SupportedKafkaSize) {
	for i := range sizes {
		if sizes[i].GetId() == selectedSize {
			f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.sizeUnit",
				localize.NewEntry("DisplaySize", sizes[i].GetDisplayName()), localize.NewEntry("Size", sizes[i].GetId())))
			if sizes[i].GetMaturityStatus() == "preview" {
				f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.sizePreview"))
			}
		}
	}
}

func promptUserForSizes(f *factory.Factory, sizes *[]kafkamgmtclient.SupportedKafkaSize) (int, error) {
	var index int

	if len(*sizes) == 1 {
		return 0, nil
	}

	sizeLabels := GetValidKafkaSizesLabels(*sizes)
	planPrompt := &survey.Select{
		Message: f.Localizer.MustLocalize("kafka.create.input.plan.message"),
		Options: sizeLabels,
	}

	err := survey.AskOne(planPrompt, &index)
	if err != nil {
		return 0, err
	}

	return index, nil
}

func getClusterDetails(f *factory.Factory, clusterId string) (*kafkamgmtclient.EnterpriseCluster, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}
	client := conn.API()
	resource := client.KafkaMgmtEnterprise().GetEnterpriseClusterById(f.Context, clusterId)
	cluster, _, err := resource.Execute()
	if err != nil {
		return nil, err
	}

	return &cluster, nil
}

func CreateNullableString(str *string) kafkamgmtclient.NullableString {
	return *kafkamgmtclient.NewNullableString(str)
}
