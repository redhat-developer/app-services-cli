// Package dump contains functions used to print documents to JSON, YAML and Table formats
package dump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"

	"github.com/landoop/tableprinter"
	"gitlab.com/c0b/go-ordered-json"
	"gopkg.in/yaml.v2"
)

const (
	JSONFormat  = "json"
	YAMLFormat  = "yaml"
	YMLFormat   = "yml"
	EmptyFormat = ""
)

// JSON dumps the given data to the given stream so that it looks pretty. If the data is a valid
// JSON document then it will be indented before printing it. If the `jq` tool is available in the
// path then it will be used for syntax highlighting.
func JSON(stream io.Writer, body []byte) error {
	if len(body) == 0 {
		return nil
	}
	data := ordered.NewOrderedMap()
	err := json.Unmarshal(body, data)
	if err != nil {
		return dumpBytes(stream, body)
	}
	if haveJQ() {
		return dumpJQ(stream, body)
	}
	return dumpJSON(stream, data)
}

// YAML dumps the given data to the given stream so that it looks pretty. If the data is a valid
// YAML document then it will be indented before printing it. If the `yq` tool is available in the
// path then it will be used for syntax highlighting.
func YAML(stream io.Writer, body []byte) error {
	if len(body) == 0 {
		return nil
	}
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal(body, data)
	if err != nil {
		return dumpBytes(stream, body)
	}
	if haveYQ(4) {
		return dumpYQ(stream, body)
	}

	return dumpYAML(stream, data)
}

// Table prints the given data into a formatted table. Only properties that have a `header`
// tag will be printed. See https://github.com/lensesio/tableprinter
func Table(stream io.Writer, in interface{}) {
	printer := tableprinter.New(stream)
	printer.Print(in)
}

func dumpBytes(stream io.Writer, data []byte) error {
	_, err := stream.Write(data)
	if err != nil {
		return err
	}
	_, err = stream.Write([]byte("\n"))
	return err
}

func dumpJQ(stream io.Writer, data []byte) error {
	// #nosec 204
	jq := exec.Command("jq", ".")
	jq.Stdin = bytes.NewReader(data)
	jq.Stdout = stream
	jq.Stderr = os.Stderr
	return jq.Run()
}

func dumpYQ(stream io.Writer, data []byte) error {
	// #nosec 204
	yq := exec.Command("yq", "eval")
	yq.Stdin = bytes.NewReader(data)
	yq.Stdout = stream
	yq.Stderr = os.Stderr
	return yq.Run()
}

func dumpJSON(stream io.Writer, data interface{}) error {
	encoder := json.NewEncoder(stream)
	encoder.SetIndent("", cmdutil.DefaultJSONIndent)
	return encoder.Encode(data)
}

func dumpYAML(stream io.Writer, data interface{}) error {
	encoder := yaml.NewEncoder(stream)
	return encoder.Encode(data)
}

func haveJQ() bool {
	_, err := exec.LookPath("jq")
	return err == nil
}

// detect if YQ is on the user's path.
// minVersion requires a minimum version of YQ to be installed.
func haveYQ(minVersion int) bool {
	_, err := exec.LookPath("yq")
	if err != nil {
		return false
	}

	versionOutput, err := exec.Command("yq", "--version").Output()
	if err != nil {
		return false
	}

	prefixes := []string{"yq", "(https://github.com/mikefarah/yq/)", "version"}

	versionStr := string(versionOutput)
	for _, p := range prefixes {
		versionStr = strings.TrimSpace(strings.TrimLeft(string(versionStr), p))
	}

	if versionStr == "" {
		return false
	}

	// if the version number starts with the minimum version number
	// YQ can be used
	if strings.HasPrefix(versionStr, fmt.Sprintf("%v.", minVersion)) {
		return true
	}

	return false
}

// Formatted prints the given data to the given format
func Formatted(writer io.Writer, format string, data interface{}) error {
	switch format {
	case YAMLFormat, YMLFormat:
		data, err := yaml.Marshal(data)
		if err != nil {
			return err
		}
		return YAML(writer, data)
	default:
		data, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return JSON(writer, data)
	}
}
