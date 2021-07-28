package util

import (
	"io"
	"io/ioutil"
	"os"
)

func CreateFileFromStdin() (*os.File, error) {
	var specifiedFile *os.File
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	specifiedFile, err = ioutil.TempFile("", "rhoas-std-input")
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
