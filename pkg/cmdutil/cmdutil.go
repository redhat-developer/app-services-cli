package cmdutil

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ConvertPageValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 1
	}

	return int32(val)
}

func ConvertSizeValueToInt32(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return 10
	}

	return int32(val)
}

// StringSliceToListStringWithQuotes converts a string slice to a
// comma-separated list with each value in quotes.
// Example: "a", "b", "c"
func StringSliceToListStringWithQuotes(validOptions []string) string {
	var listF string
	for i, val := range validOptions {
		listF += fmt.Sprintf("\"%v\"", val)
		if i < len(validOptions)-1 {
			listF += ", "
		}
	}
	return listF
}

// IsURL accepts a string and determines if it is a URL
func IsURL(s string) bool {
	return strings.HasPrefix(s, "http:/") || strings.HasPrefix(s, "https:/")
}

// GetContentFromFileURL loads file content from the provided URL
func GetContentFromFileURL(url string, ctx context.Context) (*os.File, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error loading file: %s", http.StatusText(resp.StatusCode))
	}

	respbody := resp.Body

	defer resp.Body.Close()

	tmpfile, err := ioutil.TempFile("", "rhoas_file-*")
	if err != nil {
		return nil, fmt.Errorf("error initializing temporary file: %w", err)
	}

	defer func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}()

	_, err = io.Copy(tmpfile, respbody)
	if err != nil {
		return nil, err
	}

	specifiedFile, err := os.Open(tmpfile.Name())
	if err != nil {
		return nil, err
	}

	return specifiedFile, nil
}
