package register

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"strings"
)

type options struct {
	selectedClusterId         string
	kafkaMachinePoolNodeCount int32
	clusterList               v1.ClusterList
	interactive               bool
	selectedCluster           v1.Cluster
	clusterMachinePoolList    *v1.MachinePoolList
	clusterMachinePool        v1.MachinePool
	requestedMachinePoolNodes int

	f *factory.Factory
}

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

	// TODO add Localizer and logic for pass cluster ID flag
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
	prepareRequestForKFMEndpoint(opts)

	return nil
}

func getListClusters(opts *options) error {
	// ocm client connection
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}
	client, err := conn.API().OCMClustermgmt()
	if err != nil {
		return err
	}

	resource := client.Clusters().List()
	response, err := resource.Send()
	if err != nil {
		return err
	}
	clusters := response.Items()
	opts.clusterList = *clusters
	return nil
}

func runClusterSelectionInteractivePrompt(opts *options) error {
	clusterListString := make([]string, 0)
	for _, cluster := range opts.clusterList.Slice() {
		clusterListString = append(clusterListString, cluster.Name())
	}

	//clusterList := opts.clusterList.
	// TODO add page size and Localizer
	prompt := &survey.Select{
		Message: "Select the cluster to register",
		Options: clusterListString,
	}

	var selectedClusterName string
	err := survey.AskOne(prompt, &selectedClusterName)
	if err != nil {
		return err
	}

	// get the desired cluster
	for _, cluster := range opts.clusterList.Slice() {
		if cluster.Name() == selectedClusterName {
			opts.selectedCluster = *cluster
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

// TODO add machine pool logic
func prepareRequestForKFMEndpoint(opts *options) error {
	clusterIngressDNSName, err := parseDNSURL(opts)
	if err != nil {
		return err
	}
	kfmPayload := kafkamgmtclient.EnterpriseOsdClusterPayload{
		ClusterId:                 opts.selectedCluster.ID(),
		ClusterExternalId:         opts.selectedCluster.ExternalID(),
		ClusterIngressDnsName:     clusterIngressDNSName,
		KafkaMachinePoolNodeCount: opts.kafkaMachinePoolNodeCount,
	}
	opts.f.Logger.Info("kfmPayload: ", kfmPayload)
	// TODO add kfm client and call the endpoint

	return nil
}
