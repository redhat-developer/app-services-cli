package topics

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
)

func brokerConnect() (broker *kafka.Conn, ctl *kafka.Conn) {
	// TODO enable and configure SASL plain
	// mechanism := plain.Mechanism{
	// 	Username: "username",
	// 	Password: "password",
	// }

	dialer := &kafka.Dialer{
		Timeout:   100 * time.Second,
		DualStack: true,
		//SASLMechanism: mechanism,
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	if cfg.Services.Kafka.ClusterID == "" {
		fmt.Fprint(os.Stderr, "No Kafka selected. Run rhoas kafka use")
		panic("Missing config")
	}

	connection, err := cfg.Connection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create connection: %v\n", err)
	}

	managedservices := connection.NewMASClient()
	kafkaInstance, _, err := managedservices.DefaultApi.GetKafkaById(context.TODO(), cfg.Services.Kafka.ClusterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get Kafka instance: %v\n", err)
		return
	}

	var clusterURL string
	if strings.HasPrefix(kafkaInstance.BootstrapServerHost, "localhost") {
		clusterURL = kafkaInstance.BootstrapServerHost
	} else {
		clusterURL = kafkaInstance.BootstrapServerHost + ":443"
	}

	conn, err := dialer.Dial("tcp", clusterURL)
	if err != nil {
		panic(err.Error())
	}

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}

	return conn, controllerConn
}

func CreateKafkaTopic(topicConfigs *[]kafka.TopicConfig) error {
	conn, controllerConn := brokerConnect()

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.CreateTopics(*topicConfigs...)
}

func DeleteKafkaTopic(topic string) error {
	conn, controllerConn := brokerConnect()

	defer conn.Close()
	defer controllerConn.Close()

	return controllerConn.DeleteTopics([]string{topic}...)
}

func ListKafkaTopics() error {
	conn, controllerConn := brokerConnect()

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
