package register

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	dedicatedcmdutil "github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/dedicatedcmdutil"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	selectedClusterId             string
	clusterManagementApiUrl       string
	accessToken                   string
	clusterList                   []clustersmgmtv1.Cluster
	selectedCluster               clustersmgmtv1.Cluster
	existingMachinePoolList       []clustersmgmtv1.MachinePool
	selectedClusterMachinePool    clustersmgmtv1.MachinePool
	requestedMachinePoolNodeCount int
	accessKafkasViaPrivateNetwork bool
	// newMachinePool                clustersmgmtv1.MachinePool

	f *factory.Factory
}

// list of consts should come from KFM
const (
	machinePoolId                  = "kafka-standard"
	machinePoolTaintKey            = "bf2.org/kafkaInstanceProfileType"
	machinePoolTaintEffect         = "NoExecute"
	machinePoolTaintValue          = "standard"
	machinePoolInstanceType        = "r5.xlarge"
	machinePoolLabelKey            = "bf2.org/kafkaInstanceProfileType"
	machinePoolLabelValue          = "standard"
	clusterReadyState              = "ready"
	defaultClusterManagementAPIURL = "https://api.openshift.com"
	fleetshardAddonId              = "kas-fleetshard-operator"
	strimziAddonId                 = "managed-kafka"
	fleetshardAddonIdQE            = "kas-fleetshard-operator-qe"
	strimziAddonIdQE               = "managed-kafka-qe"
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

	return cmd
}

func runRegisterClusterCmd(opts *options) (err error) {
	// Set the base URL for the cluster management API
	if opts.clusterManagementApiUrl == "" {
		opts.clusterManagementApiUrl = defaultClusterManagementAPIURL
	}
	err = setListClusters(opts)
	if err != nil {
		return err
	}
	if len(opts.clusterList) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.registerCluster.run.noClusterFound")
	}
	// TO-DO if client has supplied a cluster id, validate it and set it as the selected cluster without listing getting all clusters
	if opts.selectedClusterId == "" {
		err = runClusterSelectionInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else {
		for i := range opts.clusterList {
			cluster := opts.clusterList[i]
			if cluster.ID() == opts.selectedClusterId {
				opts.selectedCluster = cluster
			}
		}
	}
	err = getOrCreateMachinePoolList(opts)
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

func getClusterList(opts *options) (*clustersmgmtv1.ClusterList, error) {
	conn, err := opts.f.Connection()
	if err != nil {
		return nil, err
	}
	client, cc, err := conn.API().OCMClustermgmt(opts.clusterManagementApiUrl, opts.accessToken)
	if err != nil {
		return nil, err
	}
	defer cc()
	// TO-DO deal with pagination, validate clusters -- must be multi AZ and ready.
	resource := client.Clusters().List()
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	clusters := response.Items()
	return clusters, nil
}

func setListClusters(opts *options) error {
	clusters, err := getClusterList(opts)
	if err != nil {
		return err
	}
	var cls = []clustersmgmtv1.Cluster{}
	cls = validateClusters(clusters, cls)
	opts.clusterList = cls
	return nil
}

func validateClusters(clusters *clustersmgmtv1.ClusterList, cls []clustersmgmtv1.Cluster) []clustersmgmtv1.Cluster {
	for _, cluster := range clusters.Slice() {
		if cluster.State() == clusterReadyState && cluster.MultiAZ() == true {
			cls = append(cls, *cluster)
		}
	}
	return cls
}

func runClusterSelectionInteractivePrompt(opts *options) error {
	// TO-DO handle in case of empty cluster list, must be cleared up with UX etc.
	clusterStringList := make([]string, 0)
	for i := range opts.clusterList {
		cluster := opts.clusterList[i]
		clusterStringList = append(clusterStringList, cluster.Name())
	}

	// TO-DO add page size
	prompt := &survey.Select{
		Message: opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectCluster.message"),
		Options: clusterStringList,
	}

	var selectedClusterName string
	err := survey.AskOne(prompt, &selectedClusterName)
	if err != nil {
		return err
	}

	// get the desired cluster
	for i := range opts.clusterList {
		cluster := opts.clusterList[i]
		if cluster.Name() == selectedClusterName {
			opts.selectedCluster = cluster
		}
	}
	return nil
}

// parses the
func parseDNSURL(opts *options) (string, error) {
	clusterIngressDNSName := opts.selectedCluster.Console().URL()
	if len(clusterIngressDNSName) == 0 {
		return "", fmt.Errorf("DNS url is empty")
	}
	return strings.SplitAfter(clusterIngressDNSName, ".apps.")[1], nil
}

func getOrCreateMachinePoolList(opts *options) error {
	// ocm client connection
	response, err := getMachinePoolList(opts)
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
			opts.existingMachinePoolList = append(opts.existingMachinePoolList, *machinePool)
		}
		err = validateMachinePoolNodes(opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func getMachinePoolList(opts *options) (*clustersmgmtv1.MachinePoolsListResponse, error) {
	conn, err := opts.f.Connection()
	if err != nil {
		return nil, err
	}
	client, cc, err := conn.API().OCMClustermgmt(opts.clusterManagementApiUrl, opts.accessToken)
	if err != nil {
		return nil, err
	}
	defer cc()
	resource := client.Clusters().Cluster(opts.selectedCluster.ID()).MachinePools().List()
	response, err := resource.Send()
	if err != nil {
		return nil, err
	}
	return response, nil
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

// TO-DO this function should be moved to an ocm client / provider area
func createMachinePool(opts *options, mprequest *clustersmgmtv1.MachinePool) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt(opts.clusterManagementApiUrl, opts.accessToken)
	if err != nil {
		return err
	}
	defer cc()
	response, err := client.Clusters().Cluster(opts.selectedCluster.ID()).MachinePools().Add().Body(mprequest).Send()
	if err != nil {
		return err
	}
	opts.selectedClusterMachinePool = *response.Body()
	return nil
}

func createMachinePoolInteractivePrompt(opts *options) error {
	validator := &dedicatedcmdutil.Validator{
		Localizer:  opts.f.Localizer,
		Connection: opts.f.Connection,
	}
	// TO-DO add page size and better help message
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
	err = createMachinePool(opts, dedicatedMachinePool)
	if err != nil {
		return err
	}
	return nil
}

// machine pool replica count must be greater than or equal and a multiple of 3
func validateMachinePoolNodes(opts *options) error {
	for i := range opts.existingMachinePoolList {

		machinePool := opts.existingMachinePoolList[i]

		nodeCount := getMachinePoolNodeCount(&machinePool)

		if validateMachinePoolNodeCount(nodeCount) &&
			checkForValidMachinePoolLabels(&machinePool) &&
			checkForValidMachinePoolTaints(&machinePool) {
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

func getMachinePoolNodeCount(machinePool *clustersmgmtv1.MachinePool) int {
	var nodeCount int
	replicas, ok := machinePool.GetReplicas()
	if ok {
		nodeCount = replicas
	} else {
		autoscaledReplicas, ok := machinePool.GetAutoscaling()
		if ok {
			nodeCount = autoscaledReplicas.MaxReplicas()
		}
	}
	return nodeCount
}

func selectAccessPrivateNetworkInteractivePrompt(opts *options) error {
	prompt := &survey.Confirm{
		Message: opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectPublicNetworkAccess.message"),
		Help:    opts.f.Localizer.MustLocalize("dedicated.registerCluster.prompt.selectPublicNetworkAccess.help"),
		Default: false,
	}
	accessKafkasViaPublicNetwork := false
	err := survey.AskOne(prompt, &accessKafkasViaPublicNetwork)
	if err != nil {
		return err
	}
	if accessKafkasViaPublicNetwork {
		opts.accessKafkasViaPrivateNetwork = false
	} else {
		opts.accessKafkasViaPrivateNetwork = true
	}

	return nil
}

func newAddonParameterListBuilder(params *[]kafkamgmtclient.FleetshardParameter) *clustersmgmtv1.AddOnInstallationParameterListBuilder {
	if params == nil {
		return nil
	}
	var items []*clustersmgmtv1.AddOnInstallationParameterBuilder
	for _, p := range *params {
		pb := clustersmgmtv1.NewAddOnInstallationParameter().ID(*p.Id).Value(*p.Value)
		items = append(items, pb)
	}
	return clustersmgmtv1.NewAddOnInstallationParameterList().Items(items...)
}

func createAddonWithParams(opts *options, addonId string, params *[]kafkamgmtclient.FleetshardParameter) error {
	// create a new addon via ocm
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt(opts.clusterManagementApiUrl, opts.accessToken)
	if err != nil {
		return err
	}
	defer cc()
	addon := clustersmgmtv1.NewAddOn().ID(addonId)
	addonParameters := newAddonParameterListBuilder(params)
	addonInstallationBuilder := clustersmgmtv1.NewAddOnInstallation().Addon(addon)
	if addonParameters != nil {
		addonInstallationBuilder = addonInstallationBuilder.Parameters(addonParameters)
	}
	addonInstallation, err := addonInstallationBuilder.Build()
	if err != nil {
		return err
	}
	_, err = client.Clusters().Cluster(opts.selectedCluster.ID()).Addons().Add().Body(addonInstallation).Send()
	if err != nil {
		return err
	}

	return nil
}

func getStrimziAddonIdByEnv(con *config.Config) string {
	if con.APIUrl == build.ProductionAPIURL {
		return strimziAddonId
	}
	return strimziAddonIdQE
}

func getKafkaFleetManagerAddonIdByEnv(con *config.Config) string {
	if con.APIUrl == build.ProductionAPIURL {
		return fleetshardAddonId
	}
	return fleetshardAddonIdQE
}

// TO-DO go through errs and make them more user friendly with actual error messages.
func registerClusterWithKasFleetManager(opts *options) error {
	clusterIngressDNSName, err := parseDNSURL(opts)
	if err != nil {
		return err
	}

	nodeCount := getMachinePoolNodeCount(&opts.selectedClusterMachinePool)
	kfmPayload := kafkamgmtclient.EnterpriseOsdClusterPayload{
		AccessKafkasViaPrivateNetwork: opts.accessKafkasViaPrivateNetwork,
		ClusterId:                     opts.selectedCluster.ID(),
		ClusterExternalId:             opts.selectedCluster.ExternalID(),
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
	err = createAddonWithParams(opts, getStrimziAddonIdByEnv(con), nil)
	if err != nil {
		return err
	}
	err = createAddonWithParams(opts, getKafkaFleetManagerAddonIdByEnv(con), response.FleetshardParameters)
	if err != nil {
		return err
	}
	opts.f.Logger.Debugf("response fleetshard params: ", response.FleetshardParameters)
	opts.f.Logger.Debugf("r: ", r)
	return nil
}
