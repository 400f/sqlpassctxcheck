package main

import (
	"github.com/400f/sqlpassctxcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(sqlpassctxcheck.Analyzer)
}
