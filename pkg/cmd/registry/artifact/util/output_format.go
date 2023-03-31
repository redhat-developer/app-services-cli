package util

import (
	"errors"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"io"
	"strings"
)

type OutputFormat int

const (
	UnknownOutputFormat OutputFormat = iota
	TableOutputFormat
	YamlOutputFormat
	JsonOutputFormat
)

// Check for the UnknownOutputFormat format case and raise an error
func OutputFormatFromString(raw string) OutputFormat {
	raw = strings.TrimSpace(raw)
	raw = strings.ToLower(raw)
	if len(raw) == 0 {
		return TableOutputFormat
	}
	switch raw {
	case "table":
		return TableOutputFormat
	case "yaml", "yml":
		return YamlOutputFormat
	case "json":
		return JsonOutputFormat
	}
	return UnknownOutputFormat
}

func Dump(out io.Writer, format OutputFormat, tableData interface{}, fullData interface{}) error {
	if fullData == nil {
		fullData = tableData
	}
	switch format {
	case TableOutputFormat:
		_, _ = out.Write([]byte("\n"))
		dump.Table(out, tableData)
		_, _ = out.Write([]byte("\n"))
	case JsonOutputFormat:
		err := dump.Formatted(out, dump.JSONFormat, fullData)
		if err != nil {
			return err
		}
	case YamlOutputFormat:
		err := dump.Formatted(out, dump.YAMLFormat, fullData)
		if err != nil {
			return err
		}
	case UnknownOutputFormat:
		return errors.New("unknown output format")
	}
	return nil
}
