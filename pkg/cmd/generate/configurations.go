package generate

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

// configuration types for generate-config command
const (
	envFormat        = "env"
	jsonFormat       = "json"
	propertiesFormat = "properties"
	configmapFormat  = "configmap"
)

var configurationTypes = []string{envFormat, jsonFormat, propertiesFormat, configmapFormat}

var (
	envConfig               = template.Must(template.New(envFormat).Parse(templateEnv))
	jsonConfig              = template.Must(template.New(jsonFormat).Parse(templateJSON))
	propertiesConfig        = template.Must(template.New(propertiesFormat).Parse(templateProperties))
	configMapTemplateConfig = template.Must(template.New(configmapFormat).Parse(templateConfigMap))
)

// WriteConfig saves the configurations to a file
// in the specified output format
func WriteConfig(opts *options, config *configValues) (string, error) {

	var fileBody bytes.Buffer
	fileTemplate := getFileFormat(opts.configType)
	err := fileTemplate.Execute(&fileBody, config)
	if err != nil {
		return "", err
	}

	fileData := []byte(fileBody.String())
	if opts.fileName == "" {
		opts.fileName = getDefaultPath(opts.configType)
	}

	// If the file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	_, err = os.Stat(opts.fileName)
	if err == nil && !opts.overwrite {
		return "", opts.localizer.MustLocalizeError("generate.error.configFileAlreadyExists", localize.NewEntry("FilePath", opts.fileName))
	}

	return opts.fileName, ioutil.WriteFile(opts.fileName, fileData, 0o600)
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
	case configmapFormat:
		filePath = "rhoas-services.yaml"
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
	case configmapFormat:
		template = configMapTemplateConfig
	}

	return template
}
