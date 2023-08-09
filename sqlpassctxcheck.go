package sqlpassctxcheck

import (
	"fmt"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "sqlpassctxcheck",
	Doc:  "check for sql module method call without ctx",
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	s, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		return nil, fmt.Errorf("failed to get SSA")
	}

	restrictedMethods := getRestrictedMethodFuncList(pass)
	pass.Report = analysisutil.ReportWithoutIgnore(pass)
	for _, sf := range s.SrcFuncs {
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				for _, m := range restrictedMethods {
					if analysisutil.Called(instr, nil, m.method) {
						pass.Reportf(
							instr.Pos(),
							"use (*%s.%s).%s instead of (*%s.%s).%s",

							m.definition.PackagePath,
							m.definition.ReceiverType,
							m.definition.AlternateMethodName,
							m.definition.PackagePath,
							m.definition.ReceiverType,
							m.definition.MethodName,
						)
						break
					}
				}
			}
		}
	}
	return nil, nil
}
