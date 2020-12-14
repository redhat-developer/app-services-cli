package topics

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	sdkkafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/sdk/kafka"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func brokerConnect(cfg *config.Config, insecure bool) (broker *kafka.Conn, ctl *kafka.Conn, err error) {
	mechanism := plain.Mechanism{
		Username: cfg.ServiceAuth.ClientID,
		Password: cfg.ServiceAuth.ClientSecret,
	}

	dialer := &kafka.Dialer{
		Timeout:       100 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		// #nosec 402
		TLS: &tls.Config{
			InsecureSkipVerify: insecure,
		},
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

	sdkkafka.TransformResponse(&kafkaInstance)

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

func ValidateCredentials(cfg *config.Config) error {
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
			return err
		}
		cfg.ServiceAuth.ClientID = response.ClientID
		cfg.ServiceAuth.ClientSecret = response.ClientSecret
		if err = cfg.Save(); err != nil {
			return err
		}
	}
	return nil
}

func CreateKafkaTopic(topicConfigs []kafka.TopicConfig, cfg *config.Config, insecure bool) error {
	conn, controllerConn, err := brokerConnect(cfg, insecure)
	if err != nil {
		return err
	}

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.CreateTopics(topicConfigs...)
}

func DeleteKafkaTopic(topic string, cfg *config.Config, insecure bool) error {
	conn, controllerConn, err := brokerConnect(cfg, insecure)
	if err != nil {
		return err
	}

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.DeleteTopics([]string{topic}...)
}

func ListKafkaTopics(cfg *config.Config, insecure bool) error {
	conn, controllerConn, err := brokerConnect(cfg, insecure)
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
