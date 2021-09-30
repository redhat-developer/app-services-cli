package analyzer

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestAll(t *testing.T) {
	wd, _ := os.Getwd()
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	if err := Analyzer.Flags.Set("path", filepath.Join(filepath.Dir(wd), "localize", "locales")); err != nil {
		log.Fatal("Unable to set Analyzer flags")
	}
	analysistest.Run(t, testdata, Analyzer, "p")
}
