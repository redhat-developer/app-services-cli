package util

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CreateFileFromStdin() (*os.File, error) {
	var specifiedFile *os.File
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	specifiedFile, err = os.CreateTemp("", "rhoas-std-input")
	if err != nil {
		return nil, err
	}
	_, err = (*specifiedFile).Write(data)
	if err != nil {
		return nil, err
	}
	_, err = (*specifiedFile).Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return specifiedFile, nil
}

// IsURL accepts a string and determines if it is a URL
func IsURL(s string) bool {
	return strings.HasPrefix(s, "http:/") || strings.HasPrefix(s, "https:/")
}

// GetContentFromFileURL loads file content from the provided URL
func GetContentFromFileURL(ctx context.Context, url string) (*os.File, error) {

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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	tmpfile, err := os.CreateTemp("", "rhoas-std-input")
	if err != nil {
		return nil, fmt.Errorf("error initializing temporary file: %w", err)
	}

	_, err = (*tmpfile).Write(data)
	if err != nil {
		return nil, err
	}
	_, err = (*tmpfile).Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return tmpfile, nil
}
