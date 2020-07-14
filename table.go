package sqlmock

import "reflect"

// Table Table
type Table struct {
	schema reflect.Type
	data   []interface{}
}

// NewTable NewTable
func NewTable() *Table {
	return &Table{}
}

// Schema Schema
func (t *Table) Schema(schema interface{}) *Table {
	t.schema = reflect.TypeOf(schema)
	return t
}

// Add Add
func (t *Table) Add(data interface{}) *Table {
	schema := reflect.TypeOf(data)
	if t.schema != schema {
		panic("schema mismatch")
	}
	t.data = append(t.data, data)
	return t
}

// BatchAdd BatchAdd
func (t *Table) BatchAdd(data ...interface{}) *Table {
	for _, each := range data {
		t.Add(each)
	}
	return t
}
