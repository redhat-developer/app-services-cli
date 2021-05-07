package credentials

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/color"

	"github.com/AlecAivazis/survey/v2"

	"github.com/MakeNowJust/heredoc"
)

// Templates
var (
	templateProperties = heredoc.Doc(`
	## Generated by rhoas cli
	clientID=%v
	clientSecret=%v
	`)

	templateEnv = heredoc.Doc(`
	## Generated by rhoas cli
	CLIENT_ID=%v
	CLIENT_SECRET=%v
	`)

	templateJSON = heredoc.Doc(`
	{ 
		"clientID":"%v", 
		"clientSecret":"%v" 
	}`)
)

// Credentials is a type which represents the credentials
// for a service account
type Credentials struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// GetDefaultPath returns the default absolute path for the credentials file
func GetDefaultPath(outputFormat string) (filePath string) {
	switch outputFormat {
	case "env":
		filePath = ".env"
	case "properties":
		filePath = "credentials.properties"
	case "json":
		filePath = "credentials.json"
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
func Write(output string, filepath string, credentials *Credentials) error {
	fileTemplate := getFileFormat(output)
	fileBody := fmt.Sprintf(fileTemplate, credentials.ClientID, credentials.ClientSecret)

	fileData := []byte(fileBody)

	// replace any env vars in the file path
	trueFilePath := os.ExpandEnv(filepath)

	return ioutil.WriteFile(trueFilePath, fileData, 0600)
}

func getFileFormat(output string) (format string) {
	switch output {
	case "env":
		format = templateEnv
	case "properties":
		format = templateProperties
	case "json":
		format = templateJSON
	}

	return format
}

// ChooseFileLocation starts an interactive prompt to get the path to the credentials file
// a while loop will be entered as it can take multiple attempts to find a suitable location
// if the file already exists
func ChooseFileLocation(outputFormat string, filePath string, overwrite bool) (string, bool, error) {
	chooseFileLocation := true

	defaultPath := GetDefaultPath(outputFormat)

	for chooseFileLocation {
		// choose location
		fileNamePrompt := &survey.Input{
			Message: "Credentials file location",
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
			Message: fmt.Sprintf("File %v already exists. Do you want to overwrite it?", color.CodeSnippet(filePath)),
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
		return "", overwrite, errors.New("you must specify a file to save the service account credentials")
	}

	return filePath, overwrite, nil
}
