package register

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	dedicatedcmdutil "github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/dedicatedcmdutil"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
	"strings"
)

type options struct {
	selectedClusterId string
	clusterList       []v1.Cluster
	//interactive                   bool
	selectedCluster         v1.Cluster
	clusterMachinePoolList  v1.MachinePoolList
	existingMachinePoolList []v1.MachinePool
	//validatedMachinePoolList      []v1.MachinePool
	selectedClusterMachinePool    v1.MachinePool
	requestedMachinePoolNodeCount int
	accessKafkasViaPrivateNetwork bool
	newMachinePool                v1.MachinePool
	fleetshardParams              map[string]string

	f *factory.Factory
}

const (
	machinePoolId          = "kafka-standard"
	machinePoolTaintKey    = "bf2.org/kafkaInstanceProfileType"
	machinePoolTaintEffect = "NoExecute"
	machinePoolTaintValue  = "standard"
	//machinePoolInstanceType = "m5.2xlarge"
	machinePoolInstanceType = "r5.xlarge"
	machinePoolLabelKey     = "bf2.org/kafkaInstanceProfileType"
	machinePoolLabelValue   = "standard"
)

func NewRegisterClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	//TODO add Localizer
	cmd := &cobra.Command{
		Use: "register-cluster",
		// Short: f.Localizer.MustLocalize("registerCluster.cmd.shortDescription"),
		Short:   "registerCluster.cmd.shortDescription",
		Long:    "registerCluster.cmd.longDescription",
		Example: "registerCluster.cmd.example",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runRegisterClusterCmd(opts)
		},
	}

	// TODO add Localizer and flags
	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	// this flag will allow the user to pass the cluster id as a flag
	opts.selectedClusterId = "123"
	flags.StringVar(&opts.selectedClusterId, "cluster-id", "", "cluster id")
	//flags.StringVar(&opts.selectedClusterId, "cluster-id", "", f.Localizer.MustLocalize("registerCluster.flag.clusterId"))

	return cmd
}

func runRegisterClusterCmd(opts *options) error {
	getListClusters(opts)
	runClusterSelectionInteractivePrompt(opts)
	getMachinePoolList(opts)
	selectAccessPrivateNetworkInteractivePrompt(opts)
	registerClusterWithKasFleetManager(opts)

	return nil
}

func getListClusters(opts *options) error {
	// ocm client connection
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt()
	if err != nil {
		return err
	}

	resource := client.Clusters().List()
	response, err := resource.Send()
	if err != nil {
		return err
	}
	clusters := response.Items()
	var cls = []v1.Cluster{}
	for _, cluster := range clusters.Slice() {
		// TODO the cluster must be multiAZ
		//if cluster.State() == "ready" && cluster.MultiAZ() == true {
		if cluster.State() == "ready" {
			cls = append(cls, *cluster)
		}
	}
	opts.clusterList = cls
	defer cc()
	//defer conn.Close()
	return nil
}

func runClusterSelectionInteractivePrompt(opts *options) error {
	// get rid of this
	clusterListString := make([]string, 0)
	for _, cluster := range opts.clusterList {
		clusterListString = append(clusterListString, cluster.Name())
	}

	// TODO add page size and Localizer
	prompt := &survey.Select{
		Message: "Select the ready cluster to register",
		Options: clusterListString,
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
		return "", fmt.Errorf("url is empty")
	}
	return strings.SplitAfter(clusterIngressDNSName, ".apps.")[1], nil
}

// TODO this function should be split the ocm call and the response flow logic
func getMachinePoolList(opts *options) error {
	// ocm client connection
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt()
	if err != nil {
		return err
	}
	resource := client.Clusters().Cluster(opts.selectedCluster.ID()).MachinePools().List()
	response, err := resource.Send()
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
	defer cc()
	return nil
}

func checkForValidMachinePoolLabels(machinePool v1.MachinePool) bool {
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

func checkForValidMachinePoolTaints(machinePool v1.MachinePool) bool {
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

func createNewMachinePoolTaintsDedicated() *v1.TaintBuilder {
	return v1.NewTaint().
		Key(machinePoolTaintKey).
		Effect(machinePoolTaintEffect).
		Value(machinePoolTaintValue)
}

func createNewMachinePoolLabelsDedicated() map[string]string {
	return map[string]string{
		machinePoolLabelKey: machinePoolLabelValue,
	}
}

func createMachinePoolRequestForDedicated(machinePoolNodeCount int) (*v1.MachinePool, error) {
	mp := v1.NewMachinePool()
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
func createMachinePool(opts *options, mprequest *v1.MachinePool) error {
	// create a new machine pool via ocm
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, cc, err := conn.API().OCMClustermgmt()
	if err != nil {
		return err
	}
	response, err := client.Clusters().Cluster(opts.selectedCluster.ID()).MachinePools().Add().Body(mprequest).Send()
	if err != nil {
		return err
	}
	opts.selectedClusterMachinePool = *response.Body()
	defer cc()
	return nil
}

func createMachinePoolInteractivePrompt(opts *options) error {
	validator := &dedicatedcmdutil.Validator{
		Localizer:  opts.f.Localizer,
		Connection: opts.f.Connection,
	}
	// TODO add page size and Localizer
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
// refactor to take in list -- select or create
func validateMachinePoolNodes(opts *options) error {
	for _, machinePool := range opts.existingMachinePoolList {

		// manyandas code here doesn't work.
		//var nodeCount int
		//replicas, ok :=	machinePool.GetReplicas()
		//if ok {
		//	nodeCount = replicas
		//} else {
		//	autoscaledReplicas, ok := machinePool.GetAutoscaling()
		//	if ok {
		//		nodeCount = autoscaledReplicas.MaxReplicas()
		//	}
		//}

		// does this do as expected??
		mp, ok := machinePool.GetReplicas()
		if !ok {
			mp = 0
		} else {
			autoScaledReplicas, ok := machinePool.GetAutoscaling()
			if ok {
				mp = autoScaledReplicas.MaxReplicas()
			}
		}
		if validateMachinePoolNodeCount(mp) &&
			checkForValidMachinePoolLabels(machinePool) &&
			checkForValidMachinePoolTaints(machinePool) {
			opts.f.Logger.Info("Found a valid machine pool: %s", machinePool.ID())
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

func selectAccessPrivateNetworkInteractivePrompt(opts *options) error {
	options := []string{"Yes", "No"}
	prompt := &survey.Select{
		Message: "Do you want to access Kafkas via a private network?",
		Options: options,
	}
	err := survey.AskOne(prompt, &options)
	if err != nil {
		return err
	}
	if options[0] == "Yes" {
		opts.accessKafkasViaPrivateNetwork = true
	} else {
		opts.accessKafkasViaPrivateNetwork = false
	}

	return nil
}

func registerClusterWithKasFleetManager(opts *options) error {
	clusterIngressDNSName, err := parseDNSURL(opts)
	if err != nil {
		return err
	}
	kfmPayload := kafkamgmtclient.EnterpriseOsdClusterPayload{
		AccessKafkasViaPrivateNetwork: opts.accessKafkasViaPrivateNetwork,
		ClusterId:                     opts.selectedCluster.ID(),
		ClusterExternalId:             opts.selectedCluster.ExternalID(),
		ClusterIngressDnsName:         clusterIngressDNSName,
		KafkaMachinePoolNodeCount:     int32(opts.selectedClusterMachinePool.Replicas()),
	}
	opts.f.Logger.Info("kfmPayload: ", kfmPayload)
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
	fsoParams := []kafkamgmtclient.FleetshardParameter{}

	for _, param := range response.GetFleetshardParameters() {
		fsoParams = append(fsoParams, kafkamgmtclient.FleetshardParameter{
			Id:    param.Id,
			Value: param.Value,
		})
	}

	opts.f.Logger.Info("fsoParams: ", fsoParams)
	// add strimzi first
	//opts.f.Logger.Info("response strimzi params: ", response.FleetshardParameters)
	// id of the param and the value is the list
	opts.f.Logger.Info("response fleetshard params: ", response.FleetshardParameters)
	opts.f.Logger.Info("r: ", r)
	return nil
}
