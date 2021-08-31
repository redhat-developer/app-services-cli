package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Editor struct {
	content  []byte
	filename string
}

func New(content []byte, filename string) *Editor {
	return &Editor{
		content:  content,
		filename: filename,
	}
}

func (e *Editor) Run() ([]byte, error) {
	shell := defaultShell
	if os.Getenv("SHELL") != "" {
		shell = os.Getenv("SHELL")
	}
	editor := defaultEditor
	if os.Getenv("EDITOR") != "" {
		editor = os.Getenv("EDITOR")
	}

	path := filepath.Join(os.TempDir(), e.filename)
	err := ioutil.WriteFile(path, e.content, 0600)
	if err != nil {
		return nil, err
	}
	defer os.Remove(path)

	args := []string{shell, shellCommandFlag, fmt.Sprintf("%s %s", editor, path)}
	// #nosec 204
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
