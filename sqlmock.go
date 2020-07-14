package sqlmock

import (
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

// Sqlmock Sqlmock
type Sqlmock struct {
	databases map[string]*Database
	config    *mysql.Config
}

// New New
func New() *Sqlmock {
	return &Sqlmock{
		databases: make(map[string]*Database),
	}
}

// Database Database
func (sm *Sqlmock) Database(name string) *Database {
	db, exist := sm.databases[name]
	if !exist {
		db = NewDatabase()
		sm.databases[name] = db
	}
	return db
}

// Open Open
func (sm *Sqlmock) Open(name string) (driver.Conn, error) {
	config, err := mysql.ParseDSN(name)
	if err != nil {
		return nil, err
	}
	sm.config = config
	return &conn{sm}, nil
}
