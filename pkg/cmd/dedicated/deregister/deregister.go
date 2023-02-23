package deregister

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	ocmUtils "github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
	"strconv"
)

type options struct {
	selectedClusterId             string
	clusterManagementApiUrl       string
	accessToken                   string
	selectedCluster               clustersmgmtv1.Cluster
	existingMachinePoolList       []clustersmgmtv1.MachinePool
	selectedClusterMachinePool    clustersmgmtv1.MachinePool
	requestedMachinePoolNodeCount int
	accessKafkasViaPrivateNetwork bool

	f *factory.Factory
}

// list of consts should come from KFM
const (
	clusterReadyState      = "ready"
	fleetshardAddonId      = "kas-fleetshard-operator"
	strimziAddonId         = "managed-kafka"
	fleetshardAddonIdQE    = "kas-fleetshard-operator-qe"
	strimziAddonIdQE       = "managed-kafka-qe"
	billingModelEnterprise = "enterprise"
)

func NewDeRegisterClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "deregister-cluster",
		Short:   f.Localizer.MustLocalize("dedicated.registerCluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("dedicated.registerCluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("dedicated.registerCluster.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDeRegisterClusterCmd(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("dedicated.registerCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("dedicated.registercluster.flag.accessToken.description"))
	flags.StringVar(&opts.selectedClusterId, "cluster-id", "", f.Localizer.MustLocalize("dedicated.registerCluster.flag.clusterId.description"))

	return cmd
}

func runDeRegisterClusterCmd(opts *options) (err error) {
	// Set the base URL for the cluster management API
	clusterList, err := getListOfClusters(opts)
	if err != nil {
		return err
	}

	if len(clusterList) == 0 {
		return opts.f.Localizer.MustLocalizeError("dedicated.registerCluster.run.noClusterFound")
	}

	// TO-DO if client has supplied a cluster id, validate it and set it as the selected cluster without listing getting all clusters
	if opts.selectedClusterId == "" {
		err = runClusterSelectionInteractivePrompt(opts, &clusterList)
		if err != nil {
			return err
		}
	} else {
		for i := range clusterList {
			cluster := clusterList[i]
			if cluster.ID() == opts.selectedClusterId {
				opts.selectedCluster = cluster
			}
		}
	}

	kafkas, err := getkafkasInCluster(opts)
	if err != nil {
		return err
	}

	err = deleteKafkasPrompt(opts, &kafkas)
	if err != nil {
		return err
	}

	err = deregisterClusterFromKasFleetManager(opts)
	if err != nil {
		return err
	}

	addonIdsToDelete := []string{fleetshardAddonIdQE, fleetshardAddonId, strimziAddonId, strimziAddonIdQE}
	err = ocmUtils.RemoveAddonsFromCluster(opts.f, opts.clusterManagementApiUrl, opts.accessToken, &opts.selectedCluster, addonIdsToDelete)
	if err != nil {
		return err
	}

	return nil
}

func getListOfClusters(opts *options) ([]clustersmgmtv1.Cluster, error) {
	clusters, err := ocmUtils.GetClusterList(opts.f, opts.accessToken, opts.clusterManagementApiUrl, 1, 99)
	if err != nil {
		return nil, err
	}

	var list []clustersmgmtv1.Cluster
	for _, cluster := range clusters.Slice() {
		if cluster.State() == clusterReadyState && cluster.MultiAZ() == true {
			list = append(list, *cluster)
		}
	}

	return list, nil
}

func getkafkasInCluster(opts *options) ([]kafkamgmtclient.KafkaRequest, error) {

	conn, err := opts.f.Connection()
	if err != nil {
		return nil, err
	}

	api := conn.API()

	// pagination is cringe
	a := api.KafkaMgmt().GetKafkas(opts.f.Context).Page(strconv.Itoa(1)).Size(strconv.Itoa(99)).Search(fmt.Sprintf("cluster_id = %v", opts.selectedCluster.ID()))

	// deal with response errors at some point
	kafkaList, _, err := a.Execute()
	if err != nil {
		return nil, err
	}

	return kafkaList.Items, nil
}

func deleteKafkasPrompt(opts *options, kafkas *[]kafkamgmtclient.KafkaRequest) error {

	for _, kafka := range *kafkas {
		promptConfirmName := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("kafka.delete.input.confirmName.message", localize.NewEntry("Name", kafka.GetName())),
		}

		var confirmedKafkaName string
		err := survey.AskOne(promptConfirmName, &confirmedKafkaName)
		if err != nil {
			return err
		}

		if confirmedKafkaName != kafka.GetName() {
			opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.delete.log.info.incorrectNameConfirmation"))
			return nil
		}

		conn, err := opts.f.Connection()
		if err != nil {
			return err
		}

		api := conn.API()

		// delete the Kafka
		opts.f.Logger.Debug(opts.f.Localizer.MustLocalize("kafka.delete.log.debug.deletingKafka"), fmt.Sprintf("\"%s\"", kafka.GetName()))
		a := api.KafkaMgmt().DeleteKafkaById(opts.f.Context, kafka.GetId())
		a = a.Async(true)
		_, _, err = a.Execute()

		if err != nil {
			return err
		}
	}

	return nil
}

func runClusterSelectionInteractivePrompt(opts *options, clusterList *[]clustersmgmtv1.Cluster) error {
	// TO-DO handle in case of empty cluster list, must be cleared up with UX etc.
	clusterStringList := make([]string, 0)
	for _, cluster := range *clusterList {
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
	for _, cluster := range *clusterList {
		if cluster.Name() == selectedClusterName {
			opts.selectedCluster = cluster
			opts.selectedClusterId = cluster.ID()
		}
	}
	return nil
}

func deregisterClusterFromKasFleetManager(opts *options) error {
	opts.f.Logger.Info("De-registering cluster with name ", opts.selectedCluster.Name())

	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	client := conn.API()

	_, _, err = client.KafkaMgmtEnterprise().DeleteEnterpriseClusterById(context.Background(), opts.selectedCluster.ID()).Async(true).Execute()
	if err != nil {
		return err
	}

	return nil
}
