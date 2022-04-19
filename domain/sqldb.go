package domain

import "database/sql"

type SqlDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Select(dest interface{}, query string, args ...interface{}) error
	Begin() (*sql.Tx, error)
}
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type SqlRow interface {
	Scan(dest ...interface{}) error
}
