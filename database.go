package sqlmock

// Database Database
type Database struct {
	tables map[string]*Table
}

// NewDatabase NewDatabase
func NewDatabase() *Database {
	return &Database{}
}

// AddTable AddTable
func (db *Database) Table(name string) *Table {
	table, exist := db.tables[name]
	if !exist {
		table = NewTable()
		db.tables[name] = table
	}
	return table
}
