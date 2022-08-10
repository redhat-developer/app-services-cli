# Generating AsciiDoc Docs For Your Own cobra.Command

Generating AsciiDoc pages from a cobra command is incredibly easy. An example is as follows:

```go
package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
	options := doc.GeneratorOptions{
		Dir: "/tmp/",
	}
	err := doc.GenAsciidocTree(cmd, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

That will get you a AsciiDoc document `/tmp/test.adoc`

## Generation options


```golang
	// Directory to write the documentation files
	Dir string

	// FileNameGenerator - provides custom file name for each documentation file
	FileNameGenerator func(cmd *cobra.Command) string

	// FilePrepender - Prepend content to the generated file (add header)
	FilePrepender func(string) string

	// LinkHandler - function to handle links that lets you to transform them to different format
	LinkHandler func(string) string

	// GenerateIndex - generate index file
	GenerateIndex bool

	// IndexLocation - name and folder of the assembly file (typically ./README.adoc)
	IndexLocation string
```	

## Generate AsciiDoc docs for the entire command tree

This program can actually generate docs for the kubectl command in the kubernetes project

```go
package main

import (
	"io"
	"log"
	"os"

	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/spf13/cobra/doc"
)

func main() {
	kubectl := cmd.NewKubectlCommand(cmdutil.NewFactory(nil), os.Stdin, io.Discard, io.Discard)
	err := doc.GenAsciidocTree(kubectl, doc.GeneratorOptions{
		Dir: "/.",
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

This will generate a whole series of files, one for each command in the tree, in the directory specified (in this case "./")

## Generate AsciiDoc docs for a single command

You may wish to have more control over the output, or only generate for a single command, instead of the entire command
tree. If this is the case you may prefer to `GenAsciidoc` instead of `GenAsciidocTree`

```go
out := new(bytes.Buffer)
err := doc.GenAsciidoc(cmd, out)
if err != nil {
	log.Fatal(err)
}
```

This will write the AsciiDoc doc for ONLY "cmd" into the out, buffer.

## Customize the output

Both `GenAsciidocTree` GeneratorOptions object provides number of customizations we can execute.
By default only Dir folder is required to be present.

```go
options := doc.GeneratorOptions{
		Dir: "/.",
}
```
 
The `FilePrepender` field will prepend the return value given the full filepath to the rendered Asciidoc file. A common use case is to add front matter to use the generated documentation with [Hugo](http://gohugo.io/):

```go
const fmTemplate = `---
date: %s
title: "%s"
slug: %s
url: %s
---
`

options.FilePrepender := func(filename string) string {
	now := time.Now().Format(time.RFC3339)
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	url := "/commands/" + strings.ToLower(base) + "/"
	return fmt.Sprintf(fmTemplate, now, strings.Replace(base, "_", " ", -1), base, url)
}
```

The `LinkHandler` can be used to customize the rendered internal links to the commands, given a filename:

```go
options.LinkHandler := func(name string) string {
	base := strings.TrimSuffix(name, path.Ext(name))
	return "/commands/" + strings.ToLower(base) + "/"
}
```

The `FileNameGenerator` can be used to customize file names: 

```go
options.FileNameGenerator = func(c *cobra.Command) string {
		basename := c.CommandPath() + ".adoc"
		return filepath.Join(options.Dir, basename)
}
```