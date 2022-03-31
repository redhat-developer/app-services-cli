package generate

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// configuration types for generate-config command
const (
	envFormat        = "env"
	jsonFormat       = "json"
	propertiesFormat = "properties"
)

var configurationTypes = []string{envFormat, jsonFormat, propertiesFormat}

var (
	envConfig        = template.Must(template.New(envFormat).Parse(templateEnv))
	jsonConfig       = template.Must(template.New(jsonFormat).Parse(templateJSON))
	propertiesConfig = template.Must(template.New(propertiesFormat).Parse(templateProperties))
)

// WriteConfig saves the configurations to a file
// in the specified output format
func WriteConfig(configType string, config *configValues) error {

	var fileBody bytes.Buffer
	fileTemplate := getFileFormat(configType)
	err := fileTemplate.Execute(&fileBody, config)
	if err != nil {
		return err
	}

	fileData := []byte(fileBody.String())
	filePath := getDefaultPath(configType)

	return ioutil.WriteFile(filePath, fileData, 0o600)
}

// getDefaultPath returns the default absolute path for the configuration file
func getDefaultPath(configType string) (filePath string) {
	switch configType {
	case envFormat:
		filePath = "rhoas.env"
	case propertiesFormat:
		filePath = "rhoas.properties"
	case jsonFormat:
		filePath = "rhoas.json"
	}

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "./"
	}

	filePath = filepath.Join(pwd, filePath)

	return filePath
}

func getFileFormat(configType string) (template *template.Template) {

	switch configType {
	case envFormat:
		template = envConfig
	case propertiesFormat:
		template = propertiesConfig
	case jsonFormat:
		template = jsonConfig
	}

	return template
}
