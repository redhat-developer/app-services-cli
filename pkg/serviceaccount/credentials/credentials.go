package credentials

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"

	"github.com/MakeNowJust/heredoc"
)

// Templates
var (
	templateProperties = heredoc.Doc(`
	## Generated by rhoas cli
	user=%v
	password=%v
	`)

	templateEnv = heredoc.Doc(`
	## Generated by rhoas cli
	USER=%v
	PASSWORD=%v
	`)

	templateJSON = heredoc.Doc(`
	{ 
		"user":"%v", 
		"password":"%v" 
	}`)

	templateSecret = heredoc.Doc(`
	kind: Secret
	apiVersion: v1
	metadata:
	  name: "rhoas-service-account-secret"
	stringData:
	  clientID: "%v"
	  clientSecret: "%v"
	type: Opaque
	`)
)

// Credentials is a type which represents the SASL/Plain credentials
// for a service account
type Credentials struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// AbsolutePath returns the absolute path for the credentials file
// returning a default location based on the output format if customLocation
// is empty
func AbsolutePath(outputFormat string, customLocation string) (filePath string) {
	filePath = customLocation
	if filePath == "" {
		switch outputFormat {
		case "env":
			filePath = ".env"
		case "properties":
			filePath = "credentials.properties"
		case "json":
			filePath = "credentials.json"
		case "kube":
			filePath = "credentials.yaml"
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "./"
	}

	filePath = filepath.Join(pwd, filePath)

	return filePath
}

// Write saves the credentials to a file
// in the specified output format
func Write(output string, fileName string, credentials *Credentials) error {
	fileTemplate := getFileFormat(output)
	fileBody := fmt.Sprintf(fileTemplate, credentials.ClientID, credentials.ClientSecret)

	fileData := []byte(fileBody)

	return ioutil.WriteFile(fileName, fileData, 0600)
}

func getFileFormat(output string) (format string) {
	switch output {
	case "env":
		format = templateEnv
	case "properties":
		format = templateProperties
	case "json":
		format = templateJSON
	case "kube":
		format = templateSecret
	}

	return format
}

// ChooseFileLocation starts an interactive prompt to get the path to the credentials file
// a while loop will be entered as it can take multiple attempts to find a suitable location
// if the file already exists
func ChooseFileLocation(outputFormat string, filePath string, overwrite bool) (string, bool, error) {
	chooseFileLocation := true

	defaultPath := AbsolutePath(outputFormat, filePath)

	for chooseFileLocation {
		// choose location
		fileNamePrompt := &survey.Input{
			Message: "Credentials file location:",
			Help:    "Enter the path to the file where the service account credentials will be saved to",
			Default: defaultPath,
		}
		if filePath == "" {
			err := survey.AskOne(fileNamePrompt, &filePath, survey.WithValidator(survey.Required))
			if err != nil {
				return "", overwrite, err
			}
		}

		// check if the file selected already exists
		// if so ask the user to confirm if they would like to have it overwritten
		_, err := os.Stat(filePath)
		// file does not exist, we will create it
		if os.IsNotExist(err) {
			return filePath, overwrite, nil
		}
		// another error occurred
		if err != nil {
			return "", overwrite, err
		}

		if overwrite {
			return filePath, overwrite, nil
		}

		overwriteFilePrompt := &survey.Confirm{
			Message: fmt.Sprintf("The file '%v' already exists. Do you want to overwrite it?", filePath),
		}

		err = survey.AskOne(overwriteFilePrompt, &overwrite)
		if err != nil {
			return "", overwrite, err
		}

		if overwrite {
			return filePath, overwrite, nil
		}

		filePath = ""

		diffLocationPrompt := &survey.Confirm{
			Message: "Would you like to specify a different file location?",
		}
		err = survey.AskOne(diffLocationPrompt, &chooseFileLocation)
		if err != nil {
			return "", overwrite, err
		}
		defaultPath = ""
	}

	if filePath == "" {
		return "", overwrite, fmt.Errorf("You must specify a file to save the service account credentials")
	}

	return "", overwrite, nil
}
