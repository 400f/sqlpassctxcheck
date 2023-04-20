package sqlpassctxcheck_test

import (
	"testing"

	"github.com/400f/sqlpassctxcheck"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, sqlpassctxcheck.Analyzer, "a")
}
