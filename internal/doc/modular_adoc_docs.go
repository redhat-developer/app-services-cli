// Copyright 2015 Red Hat Inc. All rights reserved.
// Copyright 2021 Red Hat Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package doc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var linkTemplate = template.Must(template.New("linkTemplate").Parse(`
ifdef::env-github,env-browser[]
* link:{{.Link}}.adoc#{{.GitHubId}}[{{.Name}}]	 - {{.Short}}
endif::[]
ifdef::pantheonenv[]
* link:{path}#{{.PantheonId}}[{{.Name}}]	 - {{.Short}}
endif::[]
`))

func nameToPantheonId(name string) string {
	return fmt.Sprintf("ref-%s_{context}", strings.ReplaceAll(name, " ", "-"))
}

func nameToGitHubId(name string) string {
	return strings.ReplaceAll(name, " ", "-")
}

func writeXref(buf *bytes.Buffer, name string, short string) error {
	pantheonId := nameToPantheonId(name)
	gitHubId := nameToGitHubId(name)
	link := strings.ReplaceAll(name, " ", "_")
	err := linkTemplate.Execute(buf, struct {
		PantheonId string
		GitHubId   string
		Name       string
		Link       string
		Short      string
	}{
		pantheonId,
		gitHubId,
		name,
		link,
		short,
	})
	if err != nil {
		return err
	}
	return nil
}

// FlagUsages returns a string containing the usage information
// for all flags in the FlagSet.
func FlagUsages(f *pflag.FlagSet) string {
	buf := new(bytes.Buffer)

	lines := make([]string, 0)

	maxlen := 0
	f.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		line := ""
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			line = fmt.Sprintf("  `-%s`, `--%s`", flag.Shorthand, flag.Name)
		} else {
			line = fmt.Sprintf("      `--%s`", flag.Name)
		}

		varname, usage := pflag.UnquoteUsage(flag)
		if varname != "" {
			line += fmt.Sprintf(" _%s_", varname)
		}
		line += "::"
		if flag.NoOptDefVal != "" {
			switch flag.Value.Type() {
			case "string":
				line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
			case "bool":
				if flag.NoOptDefVal != "true" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			case "count":
				if flag.NoOptDefVal != "+1" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			default:
				line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
			}
		}

		// This special character will be replaced with spacing once the
		// correct alignment is calculated
		line += "\x00"
		if len(line) > maxlen {
			maxlen = len(line)
		}

		line += usage
		if !defaultIsZeroValue(flag) {
			if flag.Value.Type() == "string" {
				line += fmt.Sprintf(" (default %q)", flag.DefValue)
			} else {
				line += fmt.Sprintf(" (default %s)", flag.DefValue)
			}
		}
		if len(flag.Deprecated) != 0 {
			line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
		}

		lines = append(lines, line)
	})

	for _, line := range lines {
		sidx := strings.Index(line, "\x00")
		spacing := strings.Repeat(" ", maxlen-sidx)
		// maxlen + 2 comes from + 1 for the \x00 and + 1 for the (deliberate) off-by-one in maxlen-sidx
		fmt.Fprintln(buf, line[:sidx], spacing, line[sidx+1:])
	}

	return buf.String()
}

// defaultIsZeroValue returns true if the default value for this flag represents
// a zero value.
func defaultIsZeroValue(f *pflag.Flag) bool {
	switch f.Value.Type() {
	case "bool":
		return f.DefValue == "false"
	case "duration":
		// Beginning in Go 1.7, duration zero values are "0s"
		return f.DefValue == "0" || f.DefValue == "0s"
	case "int", "int8", "int32", "int64", "uint", "uint8", "uint16", "unit32", "uint64", "count", "float32", "float64":
		return f.DefValue == "0"
	case "string":
		return f.DefValue == ""
	case "ip", "ipMask", "ipNet":
		return f.DefValue == "<nil>"
	case "intSlice", "stringSlice", "stringArray":
		return f.DefValue == "[]"
	default:
		switch f.Value.String() {
		case "false":
			return true
		case "<nil>":
			return true
		case "":
			return true
		case "0":
			return true
		}
		return false
	}
}

func printOptions(buf *bytes.Buffer, cmd *cobra.Command) error {
	flags := cmd.NonInheritedFlags()
	flags.SetOutput(buf)
	if flags.HasAvailableFlags() {
		buf.WriteString("[discrete]\n")
		buf.WriteString("== Options\n\n")
		buf.WriteString(FlagUsages(flags))
		buf.WriteString("\n")
	}

	parentFlags := cmd.InheritedFlags()
	parentFlags.SetOutput(buf)
	if parentFlags.HasAvailableFlags() {
		buf.WriteString("[discrete]\n")
		buf.WriteString("== Options inherited from parent commands\n\n")
		buf.WriteString(FlagUsages(parentFlags))
		buf.WriteString("\n")
	}
	return nil
}

// GenAsciidoc creates asciidocs documentation
func GenAsciidoc(cmd *cobra.Command, w io.Writer) error {
	return GenAsciidocCustom(cmd, w, func(s string) string { return s })
}

// GenAsciidocCustom creates custom asciidoc documentation
func GenAsciidocCustom(cmd *cobra.Command, w io.Writer, linkHandler func(string) string) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	buf := new(bytes.Buffer)
	path := GetNormalizedCommandPath(cmd)
	headerName := GetShortCommandPath(cmd)
	buf.WriteString("ifdef::env-github,env-browser[:context: cmd]\n")
	buf.WriteString(fmt.Sprintf("[id='%s']\n", nameToPantheonId(path)))
	buf.WriteString(fmt.Sprintf("= %s\n\n", headerName))
	buf.WriteString("[role=\"_abstract\"]\n")
	buf.WriteString(cmd.Short + "\n\n")
	if len(cmd.Long) > 0 {
		buf.WriteString("[discrete]\n")
		buf.WriteString("== Synopsis\n\n")
		buf.WriteString(cmd.Long + "\n\n")
	}

	if cmd.Runnable() {
		buf.WriteString(fmt.Sprintf("....\n%s\n....\n\n", cmd.UseLine()))
	}

	if len(cmd.Example) > 0 {
		buf.WriteString("[discrete]\n")
		buf.WriteString("== Examples\n\n")
		buf.WriteString(fmt.Sprintf("....\n%s\n....\n\n", cmd.Example))
	}

	if err := printOptions(buf, cmd); err != nil {
		return err
	}
	if hasSeeAlso(cmd) {
		buf.WriteString("[discrete]\n")
		buf.WriteString("== See also\n\n")
		if cmd.HasParent() {
			parent := cmd.Parent()
			pname := parent.CommandPath()
			if err := writeXref(buf, pname, parent.Short); err != nil {
				return err
			}
			cmd.VisitParents(func(c *cobra.Command) {
				if c.DisableAutoGenTag {
					cmd.DisableAutoGenTag = c.DisableAutoGenTag
				}
			})
		}

		children := cmd.Commands()
		sort.Sort(byName(children))

		for _, child := range children {
			if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
				continue
			}
			cname := path + " " + child.Name()
			if err := writeXref(buf, cname, child.Short); err != nil {
				return err
			}
		}
		buf.WriteString("\n")
	}
	if !cmd.DisableAutoGenTag {
		buf.WriteString("====== Auto generated by spf13/cobra on " + time.Now().Format("2-Jan-2006") + "\n")
	}
	_, err := buf.WriteTo(w)
	return err
}

// GetNormalizedCommandPath returns name of the command without cmd name and with underscore instead of spaces
func GetNormalizedCommandPath(c *cobra.Command) string {
	commands := strings.Split(c.CommandPath(), " ")

	if len(commands) == 1 {
		return commands[0]
	}

	return strings.Join(commands[1:], "-")
}

func GetShortCommandPath(c *cobra.Command) string {
	commands := strings.Split(c.CommandPath(), " ")

	if len(commands) == 1 {
		return commands[0]
	}

	return strings.Join(commands[1:], " ")
}

// GenAsciidocTree will generate a markdown page for this command and all
// descendants in the directory given. The header may be nil.
// This function may not work correctly if your command names have `-` in them.
// If you have `cmd` with two subcmds, `sub` and `sub-third`,
// and `sub` has a subcommand called `third`, it is undefined which
// help output will be in the file `cmd-sub-third.1`.
func GenAsciidocTree(cmd *cobra.Command, options *GeneratorOptions) error {
	if options.Dir == "" {
		return errors.New("dir must be specified")
	}

	if options.LinkHandler == nil {
		options.LinkHandler = func(name string) string {
			return name
		}
	}

	if options.FileNameGenerator == nil {
		options.FileNameGenerator = func(c *cobra.Command) string {
			basename := GetNormalizedCommandPath(c) + ".adoc"
			return filepath.Join(options.Dir, basename)
		}
	}

	err := GenAsciidocTreeCustom(cmd, options)
	if err != nil {
		return err
	}

	if options.GenerateIndex && options.IndexFile != "" {
		return CreateIndexFile(cmd, options)
	}
	return nil
}

// GenAsciidocTreeCustom
func GenAsciidocTreeCustom(cmd *cobra.Command, options *GeneratorOptions) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := GenAsciidocTreeCustom(c, options); err != nil {
			return err
		}
	}

	filename := options.FileNameGenerator(cmd)
	fmt.Println("Generating Command", cmd.CommandPath(), "to", filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if options.FilePrepender != nil {
		outputFile := options.FilePrepender(filename)
		if _, err := io.WriteString(f, outputFile); err != nil {
			return err
		}
	}

	return GenAsciidocCustom(cmd, f, options.LinkHandler)
}
