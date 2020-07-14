package sqlmock

import (
	"database/sql/driver"
	"errors"

	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

type conn struct {
	*Sqlmock
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	p := parser.New()
	sns, _, err := p.Parse(string(query), "utf8mb4", "utf8mb4")
	if err != nil {
		return nil, err
	}
	if len(sns) != 1 {
		return nil, errors.New("invalid sql")
	}
	sn := sns[0]
	return &stmt{
		query:    query,
		numInput: getNumInput(sn),
		sn:       sn,
		Sqlmock:  c.Sqlmock,
	}, nil
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return &tx{}, nil
}
