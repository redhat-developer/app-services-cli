package register

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/dedicatedcmdutil"
	"strings"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"

	"github.com/AlecAivazis/survey/v2"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	selectedClusterId             string
	clusterManagementApiUrl       string
	accessToken                   string
	clusterList                   *clustersmgmtv1.ClusterList
	selectedCluster               *clustersmgmtv1.Cluster
	existingMachinePoolList       []*clustersmgmtv1.MachinePool
	selectedClusterMachinePool    *clustersmgmtv1.MachinePool
	requestedMachinePoolNodeCount int
	accessKafkasViaPrivateNetwork bool
	pageNumber                    int
	pageSize                      int

	f *factory.Factory
}

// list of consts should come from KFM
const (
	machinePoolId           = "kafka-standard"
	machinePoolTaintKey     = "bf2.org/kafkaInstanceProfileType"
	machinePoolTaintEffect  = "NoExecute"
	machinePoolTaintValue   = "standard"
	machinePoolInstanceType = "r5.xlarge"
	machinePoolLabelKey     = "bf2.org/kafkaInstanceProfileType"
	machinePoolLabelValue   = "standard"
	clusterReadyState       = "ready"
	fleetshardAddonId       = "kas-fleetshard-operator"
	strimziAddonId          = "managed-kafka"
	fleetshardAddonIdQE     = "kas-fleetshard-operator-qe"
	strimziAddonIdQE        = "managed-kafka-qe"
)

func NewRegisterClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "register-cluster",
		Short:   f.Localizer.MustLocalize("dedicated.registerCluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("dedicated.registerCluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("dedicated.registerCluster.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runRegisterClusterCmd(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("dedicated.registerCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("dedicated.registercluster.flag.accessToken.description"))
	flags.StringVar(&opts.selectedClusterId, "cluster-id", "", f.Localizer.MustLocalize("dedicated.registerCluster.flag.clusterId.description"))
	flags.IntVar(&opts.pageNumber, "page-number", int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), f.Localizer.MustLocalize("dedicated.registerCluster.flag.pageNumber.description"))
	flags.IntVar(&opts.pageSize, "page-size", 100, f.Localizer.MustLocalize("dedicated.registerCluster.flag.pageSize.description"))

	dedicatedcmdutil.HideClusterMgmtFlags(flags)

	return cmd
}

func runRegisterClusterCmd(opts *options) error {
	// get all valid clusters in the users org else if a clusterId is passed in via a flag, use that cluster
	var err error
	if opts.selectedClusterId == "" {

		opts.clusterList, err = clustermgmt.GetClusterListWithSearchParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, validClusterString(), opts.pageNumber, opts.pageSize)
		if err != nil {
			return err
		}
		err = runClusterSelectionInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else {
		opts.selectedCluster, err = clustermgmt.GetClusterById(opts.f, opts.accessToken, opts.clusterManagementApiUrl, opts.selectedClusterId)
		if err != nil {
			return err
		}
	}

	err = setOrCreateMachinePoolList(opts)
	if err != nil {
		return err
	}

	err = selectAccessPrivateNetworkInteractivePrompt(opts)
	if err != nil {
		return err
	}

	err = registerClusterWithKasFleetManager(opts)
	if err != nil {
		return err
	}

	return nil
}

func validClusterString() string {
	return "multi_az='true' AND state='ready'"
}

func runClusterSelectionInteractivePrompt(opts *options) error {
	if len(opts.clusterList.Slice()) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.registerCluster.run.noClusterFound")
	}
	clusterStringList := make([]string, 0)
	for i := range opts.clusterList.Slice() {
		cluster := opts.clusterList.Get(i)
		display := fmt.Sprintf("%s (%s)", cluster.Name(), cluster.ID())
		clusterStringList = append(clusterStringList, display)
	}

	prompt := &survey.Select{
		Message: opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectCluster.message"),
		Options: clusterStringList,
	}

	var idx int
	err := survey.AskOne(prompt, &idx)
	if err != nil {
		return err
	}
	opts.selectedCluster = opts.clusterList.Get(idx)
	opts.selectedClusterId = opts.clusterList.Get(idx).ID()

	return nil
}

// parseDNSURL attempts to parse the cluster ingress dns name from the console url.
func parseDNSURL(opts *options) (string, error) {
	clusterIngressDNSName := opts.selectedCluster.Console().URL()
	if len(clusterIngressDNSName) == 0 {
		return "", fmt.Errorf("DNS url is empty")
	}

	splits := strings.SplitAfterN(clusterIngressDNSName, "console-openshift-console.", 2)
	if len(splits) == 2 {
		return splits[1], nil
	}

	return "", fmt.Errorf("could not construct cluster_ingress_dns_name")
}

func setOrCreateMachinePoolList(opts *options) error {
	// ocm client connection
	response, err := clustermgmt.GetMachinePoolList(opts.f, opts.clusterManagementApiUrl, opts.accessToken, opts.selectedCluster.ID())
	if err != nil {
		return err
	}
	if response.Size() == 0 {
		err = createMachinePoolInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else {
		for _, machinePool := range response.Items().Slice() {
			opts.existingMachinePoolList = append(opts.existingMachinePoolList, machinePool)
		}
		err = validateMachinePoolNodes(opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkForValidMachinePoolLabels(machinePool *clustersmgmtv1.MachinePool) bool {
	labels := machinePool.Labels()
	for key, value := range labels {
		if key == machinePoolLabelKey && value == machinePoolLabelValue {
			return true
		}
	}
	return false
}

func validateMachinePoolNodeCount(nodeCount int) bool {
	if nodeCount <= 2 || nodeCount%3 != 0 {
		return false
	}
	return true
}

func checkForValidMachinePoolTaints(machinePool *clustersmgmtv1.MachinePool) bool {
	taints := machinePool.Taints()
	for _, taint := range taints {
		if taint.Effect() == machinePoolTaintEffect &&
			taint.Key() == machinePoolTaintKey &&
			taint.Value() == machinePoolTaintValue {
			return true
		}
	}
	return false
}

func createNewMachinePoolTaintsDedicated() *clustersmgmtv1.TaintBuilder {
	return clustersmgmtv1.NewTaint().
		Key(machinePoolTaintKey).
		Effect(machinePoolTaintEffect).
		Value(machinePoolTaintValue)
}

func createNewMachinePoolLabelsDedicated() map[string]string {
	return map[string]string{
		machinePoolLabelKey: machinePoolLabelValue,
	}
}

// TO-DO create an autoscaling machine pool
func createMachinePoolRequestForDedicated(machinePoolNodeCount int) (*clustersmgmtv1.MachinePool, error) {
	mp := clustersmgmtv1.NewMachinePool()
	mp.ID(machinePoolId).
		Replicas(machinePoolNodeCount).
		InstanceType(machinePoolInstanceType).
		Labels(createNewMachinePoolLabelsDedicated()).
		Taints(createNewMachinePoolTaintsDedicated())
	machinePool, err := mp.Build()
	if err != nil {
		return nil, err
	}
	return machinePool, nil
}

func createMachinePoolInteractivePrompt(opts *options) error {
	validator := &dedicatedcmdutil.Validator{
		Localizer:  opts.f.Localizer,
		Connection: opts.f.Connection,
	}

	promptNodeCount := &survey.Input{
		Message: opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.createMachinePoolNodeCount.message"),
		Help:    opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.createMachinePoolNodeCount.help"),
	}
	var nodeCount int
	err := survey.AskOne(promptNodeCount, &nodeCount, survey.WithValidator(validator.ValidatorForMachinePoolNodes))
	if err != nil {
		return err
	}
	opts.requestedMachinePoolNodeCount = nodeCount
	dedicatedMachinePool, err := createMachinePoolRequestForDedicated(nodeCount)
	if err != nil {
		return err
	}
	mp, err := clustermgmt.CreateMachinePool(opts.f, opts.clusterManagementApiUrl, opts.accessToken, dedicatedMachinePool, opts.selectedCluster.ID())
	if err != nil {
		return err
	}
	opts.selectedClusterMachinePool = mp
	return nil
}

// machine pool replica count must be greater than or equal and a multiple of 3
func validateMachinePoolNodes(opts *options) error {
	for i := range opts.existingMachinePoolList {

		machinePool := opts.existingMachinePoolList[i]

		nodeCount := clustermgmt.GetMachinePoolNodeCount(machinePool)

		if validateMachinePoolNodeCount(nodeCount) &&
			checkForValidMachinePoolLabels(machinePool) &&
			checkForValidMachinePoolTaints(machinePool) {
			opts.f.Logger.Infof(opts.f.Localizer.MustLocalize(
				"dedicated.registerCluster.info.foundValidMachinePool") + " " + machinePool.ID())
			opts.selectedClusterMachinePool = machinePool
			return nil
		}
		err := createMachinePoolInteractivePrompt(opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func selectAccessPrivateNetworkInteractivePrompt(opts *options) error {
	prompt := &survey.Confirm{
		Message: opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectPublicNetworkAccess.message"),
		Help:    opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectPublicNetworkAccess.help"),
		Default: false,
	}
	accessFromPublicNetwork := true
	err := survey.AskOne(prompt, &accessFromPublicNetwork)
	if err != nil {
		return err
	}
	opts.accessKafkasViaPrivateNetwork = !accessFromPublicNetwork
	return nil
}

func getStrimziAddonIdByEnv(con *config.Config) string {
	if con.APIUrl == build.ProductionAPIURL {
		return strimziAddonId
	}
	return strimziAddonIdQE
}

func getKafkaFleetShardAddonIdByEnv(con *config.Config) string {
	if con.APIUrl == build.ProductionAPIURL {
		return fleetshardAddonId
	}
	return fleetshardAddonIdQE
}

// TO-DO go through errs and make them more user-friendly with actual error messages.
func registerClusterWithKasFleetManager(opts *options) error {
	clusterIngressDNSName, err := parseDNSURL(opts)
	if err != nil {
		return err
	}

	nodeCount := clustermgmt.GetMachinePoolNodeCount(opts.selectedClusterMachinePool)
	kfmPayload := kafkamgmtclient.EnterpriseOsdClusterPayload{
		AccessKafkasViaPrivateNetwork: opts.accessKafkasViaPrivateNetwork,
		ClusterId:                     opts.selectedCluster.ID(),
		ClusterIngressDnsName:         clusterIngressDNSName,
		KafkaMachinePoolNodeCount:     int32(nodeCount),
	}
	opts.f.Logger.Debug("kfmPayload: ", kfmPayload)
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client := conn.API()
	resource := client.KafkaMgmtEnterprise().RegisterEnterpriseOsdCluster(context.Background()).EnterpriseOsdClusterPayload(kfmPayload)
	response, r, err := resource.Execute()
	if err != nil {
		return err
	}
	con, err := opts.f.Config.Load()
	if err != nil {
		return err
	}
	err = clustermgmt.CreateAddonWithParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, getStrimziAddonIdByEnv(con), nil, opts.selectedCluster.ID())
	if err != nil {
		return err
	}
	err = clustermgmt.CreateAddonWithParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, getKafkaFleetShardAddonIdByEnv(con), response.FleetshardParameters, opts.selectedCluster.ID())
	if err != nil {
		return err
	}
	opts.f.Logger.Debugf("response fleetshard params: ", response.FleetshardParameters)
	opts.f.Logger.Debugf("r: ", r)
	opts.f.Logger.Infof(opts.f.Localizer.MustLocalize("dedicated.registerCluster.info.clusterRegisteredWithKasFleetManager"))
	return nil
}
