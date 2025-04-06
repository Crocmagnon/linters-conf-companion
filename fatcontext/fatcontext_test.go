package fatcontext_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"linters-conf-companion/fatcontext"
	"os"
	"path/filepath"
	"testing"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(wd), "sample-project")

	analysistest.Run(t, testdata, fatcontext.Analyzer, "./...")
}
