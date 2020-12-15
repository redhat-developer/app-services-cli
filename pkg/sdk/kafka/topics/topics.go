package topics

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func brokerConnect() (broker *kafka.Conn, ctl *kafka.Conn, err error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, err
	}
	mechanism := plain.Mechanism{
		Username: cfg.ServiceAuth.ClientID,
		Password: cfg.ServiceAuth.ClientSecret,
	}

	dialer := &kafka.Dialer{
		Timeout:       100 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	cfg, err = config.Load()
	if err != nil {
		return nil, nil, err
	}

	if cfg.Services.Kafka.ClusterID == "" {
		return nil, nil, fmt.Errorf("No Kafka selected. Run rhoas kafka use")
	}

	connection, err := cfg.Connection()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not create connection: %w", err)
	}

	managedservices := connection.NewMASClient()
	kafkaInstance, _, err := managedservices.DefaultApi.GetKafkaById(context.TODO(), cfg.Services.Kafka.ClusterID)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not get Kafka instance: %w", err)
	}

	if kafkaInstance.BootstrapServerHost == "" {
		return nil, nil, fmt.Errorf("Kafka instance is missing a Bootstrap Server Host")
	}

	fmt.Println(kafkaInstance.BootstrapServerHost)

	conn, err := dialer.Dial("tcp", kafkaInstance.BootstrapServerHost)
	if err != nil {
		return nil, nil, err
	}

	controller, err := conn.Controller()
	if err != nil {
		return nil, nil, err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return nil, nil, err
	}

	return conn, controllerConn, nil
}

func ValidateCredentials() error {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	if cfg.ServiceAuth.ClientID == "" {
		connection, err := cfg.Connection()
		if err != nil {
			return fmt.Errorf("Can't create connection: %w", err)
		}
		client := connection.NewMASClient()
		fmt.Fprint(os.Stderr, "\nNo Service credentials. \nCreating service account for CLI\n")
		svcAcctPayload := &managedservices.ServiceAccountRequest{Name: "RHOAS-CLI", Description: "RHOAS-CLI Service Account"}
		response, _, err := client.DefaultApi.CreateServiceAccount(context.Background(), *svcAcctPayload)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return err
		}
		cfg.ServiceAuth.ClientID = response.ClientID
		cfg.ServiceAuth.ClientSecret = response.ClientSecret
		if err = config.Save(cfg); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return err
		}
	}
	return nil
}

func CreateKafkaTopic(topicConfigs *[]kafka.TopicConfig) error {
	conn, controllerConn, err := brokerConnect()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.CreateTopics(*topicConfigs...)
}

func DeleteKafkaTopic(topic string) error {
	conn, controllerConn, err := brokerConnect()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.DeleteTopics([]string{topic}...)
}

func ListKafkaTopics() error {
	conn, controllerConn, err := brokerConnect()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer controllerConn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	for i := range partitions {
		topicPartition := &partitions[i]
		replicas := strconv.Itoa(len(topicPartition.Replicas))
		fmt.Fprintf(os.Stderr, "Name: %v (Replicas: %v)\n",
			color.HiGreenString(topicPartition.Topic),
			color.HiRedString(replicas))
	}

	return nil
}
