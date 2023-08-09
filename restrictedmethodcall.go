package sqlpassctxcheck

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

type restrictedMethodCall struct {
	PackagePath         string
	ReceiverType        string
	MethodName          string
	AlternateMethodName string
}

var restrictedSqlModuleMethods = []restrictedMethodCall{
	// sql
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "Begin",
		AlternateMethodName: "BeginTx",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "Exec",
		AlternateMethodName: "ExecContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "Ping",
		AlternateMethodName: "PingContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "Prepare",
		AlternateMethodName: "PrepareContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "Query",
		AlternateMethodName: "QueryContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "DB",
		MethodName:          "QueryRow",
		AlternateMethodName: "QueryRowContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Stmt",
		MethodName:          "Exec",
		AlternateMethodName: "ExecContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Stmt",
		MethodName:          "Query",
		AlternateMethodName: "QueryContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Stmt",
		MethodName:          "QueryRow",
		AlternateMethodName: "QueryRowContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Tx",
		MethodName:          "Exec",
		AlternateMethodName: "ExecContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Tx",
		MethodName:          "Prepare",
		AlternateMethodName: "PrepareContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Tx",
		MethodName:          "Query",
		AlternateMethodName: "QueryContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Tx",
		MethodName:          "QueryRow",
		AlternateMethodName: "QueryRowContext",
	},
	{
		PackagePath:         string(sql),
		ReceiverType:        "Tx",
		MethodName:          "Stmt",
		AlternateMethodName: "StmtContext",
	},
	// sqlx
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "Beginx",
		AlternateMethodName: "BeginTxx",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "Get",
		AlternateMethodName: "GetContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "MustBegin",
		AlternateMethodName: "MustBeginTx",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "MustExec",
		AlternateMethodName: "MustExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "NamedExec",
		AlternateMethodName: "NamedExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "NamedQuery",
		AlternateMethodName: "NamedQueryContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "PrepareNamed",
		AlternateMethodName: "PrepareNamedContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "Preparex",
		AlternateMethodName: "PreparexContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "QueryRowx",
		AlternateMethodName: "QueryRowxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "Queryx",
		AlternateMethodName: "QueryxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "DB",
		MethodName:          "Select",
		AlternateMethodName: "SelectContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "Exec",
		AlternateMethodName: "ExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "Get",
		AlternateMethodName: "GetContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "MustExec",
		AlternateMethodName: "MustExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "Query",
		AlternateMethodName: "QueryContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "QueryRow",
		AlternateMethodName: "QueryRowContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "QueryRowx",
		AlternateMethodName: "QueryRowxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "Queryx",
		AlternateMethodName: "QueryxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "NamedStmt",
		MethodName:          "Select",
		AlternateMethodName: "SelectContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Stmt",
		MethodName:          "Get",
		AlternateMethodName: "GetContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Stmt",
		MethodName:          "MustExec",
		AlternateMethodName: "MustExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Stmt",
		MethodName:          "QueryRowx",
		AlternateMethodName: "QueryRowxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Stmt",
		MethodName:          "Queryx",
		AlternateMethodName: "QueryxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Stmt",
		MethodName:          "Select",
		AlternateMethodName: "SelectContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "Get",
		AlternateMethodName: "GetContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "MustExec",
		AlternateMethodName: "MustExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "NamedExec",
		AlternateMethodName: "NamedExecContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "NamedStmt",
		AlternateMethodName: "NamedStmtContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "PrepareNamed",
		AlternateMethodName: "PrepareNamedContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "Preparex",
		AlternateMethodName: "PreparexContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "QueryRowx",
		AlternateMethodName: "QueryRowxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "Queryx",
		AlternateMethodName: "QueryxContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "Select",
		AlternateMethodName: "SelectContext",
	},
	{
		PackagePath:         string(sqlx),
		ReceiverType:        "Tx",
		MethodName:          "Stmtx",
		AlternateMethodName: "StmtxContext",
	},
}

type restrictedMethod struct {
	method     *types.Func
	definition restrictedMethodCall
}

func getRestrictedMethodFuncList(pass *analysis.Pass) []restrictedMethod {
	fs := make([]restrictedMethod, 0, len(restrictedSqlModuleMethods))

	for _, m := range restrictedSqlModuleMethods {
		t := analysisutil.TypeOf(pass, m.PackagePath, m.ReceiverType)
		if t == nil {
			continue
		}

		method := analysisutil.MethodOf(t, m.MethodName)
		if method != nil {
			fs = append(fs, restrictedMethod{
				method:     method,
				definition: m,
			})
		}
	}
	return fs
}
