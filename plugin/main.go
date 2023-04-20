package main

import (
	"github.com/400f/sqlpassctxcheck"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (a analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		sqlpassctxcheck.Analyzer,
	}
}

// nolint
var AnalyzerPlugin analyzerPlugin
