package main

import (
	"github.com/redhat-developer/app-services-cli/hack/linters/go-i18n-linter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
