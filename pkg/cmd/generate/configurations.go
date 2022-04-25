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
	secretFormat     = "secret"
)

var configurationTypes = []string{envFormat, jsonFormat, propertiesFormat, secretFormat}

var (
	envConfig            = template.Must(template.New(envFormat).Parse(templateEnv))
	jsonConfig           = template.Must(template.New(jsonFormat).Parse(templateJSON))
	propertiesConfig     = template.Must(template.New(propertiesFormat).Parse(templateProperties))
	secretTemplateConfig = template.Must(template.New(secretFormat).Parse(templateSecret))
)

// WriteConfig saves the configurations to a file
// in the specified output format
func WriteConfig(configType string, filePath string, config *configValues) error {

	var fileBody bytes.Buffer
	fileTemplate := getFileFormat(configType)
	err := fileTemplate.Execute(&fileBody, config)
	if err != nil {
		return err
	}

	fileData := []byte(fileBody.String())
	if filePath == "" {
		filePath = getDefaultPath(configType)
	}

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
	case secretFormat:
		filePath = "rhoas-services-secret.yaml"
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
	case secretFormat:
		template = secretTemplateConfig
	}

	return template
}
