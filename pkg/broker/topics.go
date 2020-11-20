package broker

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
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

	if cfg.Services.Kafka.ClusterHost == "" {
		fmt.Fprint(os.Stderr, "No Kafka selected. Run rhoas kafka use")
		panic("Missing config")
	}
	var clusterURL string
	if strings.HasPrefix(cfg.Services.Kafka.ClusterHost, "localhost") {
		clusterURL = "localhost:9092"
	} else {
		clusterURL = cfg.Services.Kafka.ClusterHost + ":443"
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
		fmt.Println(&partitions[i])
	}

	return nil
}
