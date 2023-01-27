package register

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	dedicatedcmdutil "github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/dedicatedcmdutil"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
	"strings"
)

type options struct {
	selectedClusterId             string
	clusterManagementApiUrl       string
	accessToken                   string
	clusterList                   []clustersmgmtv1.Cluster
	selectedCluster               clustersmgmtv1.Cluster
	clusterMachinePoolList        clustersmgmtv1.MachinePoolList
	existingMachinePoolList       []clustersmgmtv1.MachinePool
	selectedClusterMachinePool    clustersmgmtv1.MachinePool
	requestedMachinePoolNodeCount int
	accessKafkasViaPrivateNetwork bool
	newMachinePool                clustersmgmtv1.MachinePool

	f *factory.Factory
}

// list of consts should come from KFM
const (
	machinePoolId          = "kafka-standard"
	machinePoolTaintKey    = "bf2.org/kafkaInstanceProfileType"
	machinePoolTaintEffect = "NoExecute"
	machinePoolTaintValue  = "standard"
	//machinePoolInstanceType = "m5.2xlarge"
	machinePoolInstanceType = "r5.xlarge"
	machinePoolLabelKey     = "bf2.org/kafkaInstanceProfileType"
	machinePoolLabelValue   = "standard"
	clusterReadyState       = "ready"
	fleetshardAddonId       = "kas-fleetshard-operator"
	strimziAddonId          = "managed-kafka"
	clusterManagementAPIURL = "https://api.openshift.com"
)

func NewRegisterClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	//TODO add Localizer
	cmd := &cobra.Command{
		Use: "register-cluster",
		//Short: f.Localizer.MustLocalize("registerCluster.cmd.shortDescription"),
		Short:   "registerCluster.cmd.shortDescription",
		Long:    "registerCluster.cmd.longDescription",
		Example: "registerCluster.cmd.example",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runRegisterClusterCmd(opts)
		},
	}

	// TODO add Localizer and flags
	// add a flag for clustermgmt url, i.e --cluster-management-api-url, make the flag hidden, default to api.openshift.com
	// supply customer mgmt access token via a flag, i.e --access-token, make the flag hidden, default to ""
	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	//flags.StringVar(&opts.clusterManagementApiUrl, "cluster-management-api-url", clusterManagementAPIURL, "cluster management api url")
	//flags.StringVar(&opts.accessToken, "access-token", "", "access token")
	// this flag will allow the user to pass the cluster id as a flag
	flags.StringVar(&opts.selectedClusterId, "cluster-id", "", "cluster id")
	//flags.StringVar(&opts.selectedClusterId, "cluster-id", "", f.Localizer.MustLocalize("registerCluster.flag.clusterId"))

	return cmd
}

func runRegisterClusterCmd(opts *options) error {
	setListClusters(opts)
	if opts.selectedClusterId == "" {
		runClusterSelectionInteractivePrompt(opts)
	} else {
		for _, cluster := range opts.clusterList {
			if cluster.ID() == opts.selectedClusterId {
				opts.selectedCluster = cluster
			}
		}
	}
	getOrCreateMachinePoolList(opts)
	selectAccessPrivateNetworkInteractivePrompt(opts)
	registerClusterWithKasFleetManager(opts)

	return nil
}

func getClusterList(opts *options) (*clustersmgmtv1.ClusterList, error) {
	// ocm client connection
	conn, err := opts.f.Connection()
	if err != nil {
		return nil, err
	}
	client, cc, err := conn.API().OCMClustermgmt()
	if err != nil {
		return nil, err
	}
	defer cc()
	// TODO deal with pagination, validate clusters -- must be multi AZ and ready.
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
	for _, cluster := range clusters.Slice() {
		// TODO the cluster must be multiAZ
		//if cluster.State() == clusterReadyState && cluster.MultiAZ() == true {
		if cluster.State() == clusterReadyState {
			cls = append(cls, *cluster)
		}
	}
	opts.clusterList = cls
	return nil
}

func runClusterSelectionInteractivePrompt(opts *options) error {
	// TODO handle in case of empty cluster list, must be cleared up with UX etc.
	clusterStringList := make([]string, 0)
	for _, cluster := range opts.clusterList {
		clusterStringList = append(clusterStringList, cluster.Name())
	}

	// TODO add page size and Localizer
	prompt := &survey.Select{
		Message: "Select the ready cluster to register",
		Options: clusterStringList,
	}

	var selectedClusterName string
	err := survey.AskOne(prompt, &selectedClusterName)
	if err != nil {
		return err
	}

	// get the desired cluster
	for _, cluster := range opts.clusterList {
		if cluster.Name() == selectedClusterName {
			opts.selectedCluster = cluster
		}
	}
	return nil
}

func parseDNSURL(opts *options) (string, error) {
	clusterIngressDNSName := opts.selectedCluster.Console().URL()
	if len(clusterIngressDNSName) == 0 {
		return "", fmt.Errorf("DNS url is empty")
	}
	return strings.SplitAfter(clusterIngressDNSName, ".apps.")[1], nil
}

// TODO this function should be split the ocm call and the response flow logic
func getOrCreateMachinePoolList(opts *options) error {
	// ocm client connection
	response, err := getMachinePoolList(opts)
	if err != nil {
		return err
	}
	if response.Size() == 0 {
		createMachinePoolInteractivePrompt(opts)
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
	client, cc, err := conn.API().OCMClustermgmt()
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

func checkForValidMachinePoolLabels(machinePool clustersmgmtv1.MachinePool) bool {
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

func checkForValidMachinePoolTaints(machinePool clustersmgmtv1.MachinePool) bool {
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

// TODO create an autoscaling machine pool
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

// TODO this function should be moved to an ocm client / provider area
func createMachinePool(opts *options, mprequest *clustersmgmtv1.MachinePool) error {
	// create a new machine pool via ocm
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt()
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
	// TODO add page size and Localizer, and better help message
	promptNodeCount := &survey.Input{
		Message: "Enter the desired machine pool node count",
		Help:    "The machine pool node count must be greater than or equal to and a multiple of 3",
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
	for _, machinePool := range opts.existingMachinePoolList {

		nodeCount := getMachinePoolNodeCount(machinePool)

		if validateMachinePoolNodeCount(nodeCount) &&
			checkForValidMachinePoolLabels(machinePool) &&
			checkForValidMachinePoolTaints(machinePool) {
			opts.f.Logger.Infof("Found a valid machine pool: %s", machinePool.ID())
			opts.selectedClusterMachinePool = machinePool
			return nil
		} else {
			err := createMachinePoolInteractivePrompt(opts)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getMachinePoolNodeCount(machinePool clustersmgmtv1.MachinePool) int {
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
		Message: "Would you like your Kakfas to be accessible via a public network?",
		Help:    "If you select yes, your Kafka will be accessible via a public network",
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
	client, cc, err := conn.API().OCMClustermgmt()
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

// TODO go through errs and make them more user friendly with actual error messages.
func registerClusterWithKasFleetManager(opts *options) error {
	clusterIngressDNSName, err := parseDNSURL(opts)
	if err != nil {
		return err
	}

	nodeCount := getMachinePoolNodeCount(opts.selectedClusterMachinePool)
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
	err = createAddonWithParams(opts, strimziAddonId, nil)
	if err != nil {
		return err
	}
	err = createAddonWithParams(opts, fleetshardAddonId, response.FleetshardParameters)
	if err != nil {
		return err
	}
	opts.f.Logger.Debugf("response fleetshard params: ", response.FleetshardParameters)
	opts.f.Logger.Debugf("r: ", r)
	return nil
}
