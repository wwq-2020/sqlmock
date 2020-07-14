package sqlmock

import (
	"database/sql/driver"
	"errors"
	"reflect"

	"github.com/pingcap/parser/ast"
)

type stmt struct {
	*Sqlmock
	sn           ast.StmtNode
	query        string
	numInput     int
	placeholders []string
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return s.numInput
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	stmt := s.sn.(*ast.SelectStmt)
	columns := getQueryColumns(stmt)
	tableName := stmt.From.TableRefs.Left.(*ast.TableSource).Source.(*ast.TableName).Name.L
	database, exist := s.databases[s.config.DBName]
	if !exist {
		return nil, errors.New("database not exist")
	}
	table, exist := database.tables[tableName]
	if !exist {
		return nil, errors.New("table not exist")
	}
	rows := &rows{columns: columns}
	for _, each := range table.data {
		rt := reflect.ValueOf(each)
		value := make([]driver.Value, 0, len(columns))
		for _, column := range columns {
			fv := rt.FieldByName(column)
			value = append(value, fv.Interface())
		}
		rows.values = append(rows.values, value)
	}
	return rows, nil
}
