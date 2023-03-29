package deregister

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/redhat-developer/app-services-cli/internal/build"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/openshift-cluster/openshiftclustercmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	ocmUtils "github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/clustermgmt"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
	"time"
)

type options struct {
	selectedClusterId       string
	clusterManagementApiUrl string
	accessToken             string
	selectedCluster         *clustersmgmtv1.Cluster
	machinePoolId           string

	f *factory.Factory
}

// list of consts should come from KFM
const (
	fleetshardAddonId   = "kas-fleetshard-operator"
	strimziAddonId      = "managed-kafka"
	fleetshardAddonIdQE = "kas-fleetshard-operator-qe"
	strimziAddonIdQE    = "managed-kafka-qe"
	machinePoolTaintKey = "bf2.org/kafkaInstanceProfileType"
)

func NewDeRegisterClusterCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "deregister-cluster",
		Short:   f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runDeRegisterClusterCmd(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.clusterManagementApiUrl, "cluster-mgmt-api-url", "", f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.flag.clusterMgmtApiUrl.description"))
	flags.StringVar(&opts.accessToken, "access-token", "", f.Localizer.MustLocalize("kafka.openshiftCluster.deregistercluster.flag.accessToken.description"))
	flags.StringVar(&opts.selectedClusterId, "cluster-id", "", f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.flag.clusterId.description"))

	openshiftclustercmdutil.HideClusterMgmtFlags(flags)

	return cmd
}

func runDeRegisterClusterCmd(opts *options) error {
	clusterList, err := getListOfClusters(opts)
	if err != nil {
		return err
	}

	if len(clusterList) == 0 {
		return opts.f.Localizer.MustLocalizeError("kafka.openshiftCluster.deregisterCluster.run.noClusterFound")
	}

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

		if opts.selectedCluster == nil {
			return opts.f.Localizer.MustLocalizeError("kafka.openshiftCluster.deregisterCluster.noClusterFoundFromIdFlag", localize.NewEntry("ID", opts.selectedClusterId))
		}
	}

	kafkas, err := getkafkasInCluster(opts)
	if err != nil {
		return err
	}

	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.kafka.delete.warning"))

	err = deleteKafkasPrompt(opts, &kafkas)
	if err != nil {
		return err
	}

	err = deregisterClusterFromKasFleetManager(opts)
	if err != nil {
		return err
	}

	addonIdsToDelete := []string{fleetshardAddonIdQE, fleetshardAddonId, strimziAddonId, strimziAddonIdQE}
	err = ocmUtils.RemoveAddonsFromCluster(opts.f, opts.clusterManagementApiUrl, opts.accessToken, opts.selectedCluster, addonIdsToDelete)
	if err != nil {
		return err
	}

	// Get the kafka machine pool id
	opts.machinePoolId, err = ocmUtils.GetMachinePoolIdByTaintKey(opts.f, opts.clusterManagementApiUrl, opts.accessToken, opts.selectedCluster.ID(), machinePoolTaintKey)
	if err != nil {
		return err
	}

	// remove the kafka machine pool from the cluster
	err = ocmUtils.DeleteMachinePool(opts.f, opts.clusterManagementApiUrl, opts.accessToken, opts.selectedCluster.ID(), opts.machinePoolId)
	if err != nil {
		return err
	}

	return nil
}

func getListOfClusters(opts *options) ([]*clustersmgmtv1.Cluster, error) {
	// nolint: bodyclose
	kfmClusterList, response, err := kafkautil.ListEnterpriseClusters(opts.f)
	if err != nil {
		if response != nil {
			if response.StatusCode == 403 {
				return nil, opts.f.Localizer.MustLocalizeError("kafka.openshiftCluster.deregisterCluster.error.403")
			}
			return nil, fmt.Errorf("%v: %w", response.Status, err)
		}
		return nil, err
	}
	ocmClusterList, err := clustermgmt.GetClusterListWithSearchParams(opts.f, opts.clusterManagementApiUrl, opts.accessToken, kafkautil.CreateClusterSearchStringFromKafkaList(kfmClusterList), int(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), len(kfmClusterList.Items))
	if err != nil {
		return nil, err
	}

	// this currently means there is no clusters based on the search we did, IMO
	// this should not be null and should be an empty list
	if ocmClusterList == nil {
		return make([]*clustersmgmtv1.Cluster, 0), nil
	}

	return ocmClusterList.Slice(), nil
}

func getkafkasInCluster(opts *options) ([]kafkamgmtclient.KafkaRequest, error) {

	conn, err := opts.f.Connection()
	if err != nil {
		return nil, err
	}

	api := conn.API()

	a := api.KafkaMgmt().GetKafkas(opts.f.Context).Search(fmt.Sprintf("cluster_id = %v", opts.selectedCluster.ID()))
	// nolint:bodyclose
	kafkaList, response, err := a.Execute()
	if err != nil {
		return nil, err
	}

	opts.f.Logger.Debug("HTTP Response", response)

	return kafkaList.Items, nil
}

func runKafkaNameConfirmPrompt(opts *options, kafka *kafkamgmtclient.KafkaRequest) error {
	promptConfirmName := &survey.Input{
		Message: opts.f.Localizer.MustLocalize("kafka.delete.input.confirmName.message", localize.NewEntry("Name", kafka.GetName())),
	}

	for true {
		var confirmedKafkaName string

		err := survey.AskOne(promptConfirmName, &confirmedKafkaName)
		if err != nil {
			return err
		}

		if confirmedKafkaName == kafka.GetName() {
			break
		}

		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.delete.log.info.incorrectNameConfirmation"))
	}

	return nil
}

func deleteKafkasPrompt(opts *options, kafkas *[]kafkamgmtclient.KafkaRequest) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	checkIfDeletedIdList := make([]string, 0)

	for i := 0; i < len(*kafkas); i++ {
		kafka := (*kafkas)[i]

		err := runKafkaNameConfirmPrompt(opts, &kafka)
		if err != nil {
			return err
		}

		// delete the Kafka
		opts.f.Logger.Debug(opts.f.Localizer.MustLocalize("kafka.delete.log.debug.deletingKafka"), fmt.Sprintf("\"%s\"", kafka.GetName()))
		a := api.KafkaMgmt().DeleteKafkaById(opts.f.Context, kafka.GetId()).Async(true)
		// nolint: bodyclose
		_, _, err = a.Execute()
		if err != nil {
			return err
		}

		checkIfDeletedIdList = append(checkIfDeletedIdList, kafka.GetName())
	}

	for len(checkIfDeletedIdList) > 0 {
		for i, id := range checkIfDeletedIdList {
			a := api.KafkaMgmt().GetKafkaById(opts.f.Context, id)
			_, response, err := a.Execute()

			if err != nil {
				if response == nil {
					return err
				}

				if response.StatusCode == 404 {
					// remove this callback from the callback list as the kafka is deleted
					// break to restart the loop from the beginning as we are modifying the list
					// as we are iterating through it
					checkIfDeletedIdList = append(checkIfDeletedIdList[:i], checkIfDeletedIdList[i+1:]...)
					break
				} else {
					return fmt.Errorf(fmt.Sprintf("%v, %v", opts.f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.kafka.delete.failed"), response.Status))
				}
			}
		}

		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.deletingKafka.message"))
		time.Sleep(5 * time.Second)
	}

	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.openshiftCluster.deregisterCluster.deletingKafka.success"))
	return nil
}

func runClusterSelectionInteractivePrompt(opts *options, clusterList *[]*clustersmgmtv1.Cluster) error {
	clusterStringList := make([]string, 0)
	for _, cluster := range *clusterList {
		display := fmt.Sprintf("%s (%s)", cluster.Name(), cluster.ID())
		clusterStringList = append(clusterStringList, display)
	}

	prompt := &survey.Select{
		Message: opts.f.Localizer.MustLocalize("kafka.openshiftCluster.registerCluster.prompt.selectCluster.message"),
		Options: clusterStringList,
	}

	var idx int
	err := survey.AskOne(prompt, &idx)
	if err != nil {
		return err
	}
	opts.selectedCluster = (*clusterList)[idx]
	opts.selectedClusterId = opts.selectedCluster.ID()
	return nil
}

func deregisterClusterFromKasFleetManager(opts *options) error {
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
