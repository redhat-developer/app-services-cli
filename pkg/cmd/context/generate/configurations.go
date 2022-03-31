package generate

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	envConfig        = template.Must(template.New("env").Parse(templateEnv))
	jsonConfig       = template.Must(template.New("json").Parse(templateJSON))
	propertiesConfig = template.Must(template.New("properties").Parse(templateProperties))
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
	case "env":
		filePath = "rhoas.env"
	case "properties":
		filePath = "rhoas.properties"
	case "json":
		filePath = "rhoas.json"
	}

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "./"
	}

	filePath = filepath.Join(pwd, filePath)

	return filePath
}

func getFileFormat(configType string) *template.Template {

	switch configType {
	case "env":
		return envConfig
	case "properties":
		return propertiesConfig
	default:
		return jsonConfig
	}
}
