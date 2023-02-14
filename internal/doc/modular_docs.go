package doc

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// CreateModularDocs Create Modular Documentation from the CLI generated docs
func CreateModularDocs() error {
	srcDir := os.Getenv("SRC_DIR")
	if srcDir == "" {
		return errors.New("SRC_DIR must be set")
	}
	files, err := filepath.Glob(fmt.Sprintf("%s/*.adoc", srcDir))
	if err != nil {
		return errors.WithStack(err)
	}
	outDir := os.Getenv("DEST_DIR")
	if outDir == "" {
		outDir = "dist"
	}
	modulesDir := path.Join(outDir, "modules")
	err = os.RemoveAll(modulesDir)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(modulesDir, 0o755)
	if err != nil {
		return errors.WithStack(err)
	}
	moduleFiles, err := CreateModules(modulesDir, files)
	if err != nil {
		return errors.WithStack(err)
	}

	assembliesDir := path.Join(outDir, "assemblies")
	err = os.RemoveAll(assembliesDir)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(assembliesDir, 0o755)
	if err != nil {
		return errors.WithStack(err)
	}
	err = CreateAssembly(assembliesDir, moduleFiles)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func CreateModules(modulesDir string, commandAdocFiles []string) ([]string, error) {
	answer := make([]string, 0)
	for _, f := range commandAdocFiles {
		destName := fmt.Sprintf("ref-cli%s", strings.Replace(strings.ReplaceAll(filepath.Base(f), "_", "-"), "rhoas", "", 1))
		destPath := path.Join(modulesDir, destName)
		_, err := copyFile(f, destPath)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		answer = append(answer, destPath)
	}
	return answer, nil
}

func CreateAssembly(assembliesDir string, files []string) error {
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	commandFileNames := make([]string, 0)

	for _, f := range files {
		relPath, err := filepath.Rel(assembliesDir, f)
		if err != nil {
			return errors.WithStack(err)
		}
		commandFileNames = append(commandFileNames, relPath)
	}

	contentTemplate := `:context: rhoas-cli-command-reference
[id="cli-command-reference_{context}"]
= OpenShift Application Services CLI command reference

[role="_abstract"]
You use the ` + "`rhoas`" + ` CLI to manage your application services from the command line.

{{ range .Commands}}
include::{{.}}[leveloffset=+1]
{{ end }}
`
	type Vars struct {
		Commands []string
	}

	vars := Vars{
		Commands: commandFileNames,
	}

	filename := "assembly-cli-command-reference.adoc"
	output, err := os.Create(path.Join(assembliesDir, filename))
	if err != nil {
		return errors.WithStack(err)
	}

	err = template.Must(template.New("content").Parse(contentTemplate)).Execute(output, vars)
	if err != nil {
		return errors.WithStack(err)
	}

	err = output.Sync()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
