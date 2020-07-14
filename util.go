package sqlmock

import (
	"github.com/pingcap/parser/ast"
	driver "github.com/pingcap/tidb/types/parser_driver"
)

func getQueryColumns(sn *ast.SelectStmt) []string {
	columns := make([]string, 0, len(sn.Fields.Fields))
	for _, field := range sn.Fields.Fields {
		column := field.Expr.(*ast.ColumnNameExpr).Name.Name.L
		columns = append(columns, column)
	}
	return columns
}

func getNumInput(sn ast.StmtNode) int {
	switch snReal := sn.(type) {
	case *ast.InsertStmt:
		return getNumInputForInsert(snReal)
	case *ast.SelectStmt:
		return getWhereNumInput(0, snReal.Where.(*ast.BinaryOperationExpr))
	case *ast.DeleteStmt:
		return getWhereNumInput(0, snReal.Where.(*ast.BinaryOperationExpr))
	case *ast.UpdateStmt:
		return getNumInputForUpdate(snReal.List, snReal.Where.(*ast.BinaryOperationExpr))
	default:
		panic("unexpected stmt")
	}
}

func getNumInputForInsert(sn *ast.InsertStmt) int {
	numInput := 0
	for _, list := range sn.Lists {
		for _, each := range list {
			_, ok := each.(*driver.ParamMarkerExpr)
			if ok {
				numInput++
			}
		}
	}
	return numInput
}

func getWhereNumInput(curNumInput int, be *ast.BinaryOperationExpr) int {
	switch beReal := be.L.(type) {
	case *ast.ColumnNameExpr:
		_, ok := be.R.(*driver.ParamMarkerExpr)
		if ok {
			return 1
		}
		return curNumInput + 1
	case *ast.BinaryOperationExpr:
		return getWhereNumInput(curNumInput, beReal)
	default:
		return curNumInput
	}
}

func getNumInputForUpdate(as []*ast.Assignment, be *ast.BinaryOperationExpr) int {
	curNumInput := getWhereNumInput(0, be)
	for _, each := range as {
		_, ok := each.Expr.(*driver.ParamMarkerExpr)
		if ok {
			curNumInput++
		}
	}
	return curNumInput
}
