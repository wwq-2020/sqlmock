package sqlmock

import (
	"database/sql/driver"
	"io"
)

type rows struct {
	columns  []string
	values   [][]driver.Value
	curIndex int
}

func (r *rows) Columns() []string {
	return r.columns
}

func (r *rows) Close() error {
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	if r.curIndex == len(r.values) {
		return io.EOF
	}
	value := r.values[r.curIndex]
	for index := range dest {
		dest[index] = value[index]
	}
	r.curIndex++
	return nil
}
