package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	ms "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
	msapi "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices/client"
)

// Templates
const (
	templateProperties = `
## Credentials for Kafka cluster: '%v' 
## Generated by rhoas cli
kafka_user=%v
kafka_password=%v
`

	templateEnv = `
## Credentials for Kafka cluster: '%v' 
## Generated by rhoas cli
KAFKA_USER=%v
KAFKA_PASSWORD=%v
`

	templateKafkaPlain = `
## Credentials for Kafka cluster: '%v' 
## Generated by rhoas cli
kafka.sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="%v" password="%v";
kafka.sasl.mechanism=PLAIN
kafka.security.protocol=SASL_SSL
kafka.ssl.protocol=TLSv1.2
`

	templateJSON = `
{ 
	"name":"%v",
	"user":"%v", 
	"password":"%v" 
}
`
)

func RunCredentials(outputFlagValue string) {
	var propertyFormat string
	var fileName string

	switch outputFlagValue {
	case "env":
		propertyFormat = templateEnv
		fileName = ".env"
	case "properties":
		propertyFormat = templateProperties
		fileName = "kafka.properties"
	case "kafka":
		propertyFormat = templateKafkaPlain
		fileName = "kafka.properties"
	case "json":
		propertyFormat = templateJSON
		fileName = "credentials.json"
	}

	client := ms.BuildClient()
	response, _, err := client.DefaultApi.CreateServiceAccount(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Kafka Credentials: %v\n", err)
		os.Exit(1)
	}

	jsonResponse, _ := json.Marshal(response)
	var credentials msapi.TokenResponse
	err = json.Unmarshal(jsonResponse, &credentials)

	fmt.Fprintln(os.Stderr, `Writing credentials to`, fileName)
	// TODO use https://github.com/manifoldco/promptui

	dataToWrite := []byte(propertyFormat)
	err = ioutil.WriteFile(fileName, dataToWrite, 0600)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when saving file:", err)
	} else {
		fmt.Fprintln(os.Stderr, "Successfully saved credentials to", fileName)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
