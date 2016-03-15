// Package gorpish provides mockable interfaces to gorp
// databases and transactions.
package gorpish

import (
	"database/sql"
	"database/sql/driver"

	"gopkg.in/gorp.v1"
)

type IDB interface {
	Begin() (ITX, error)
	gorp.SqlExecutor
}

// ITX is an interface to the transaction structure
type ITX interface {
	gorp.SqlExecutor
	driver.Tx
}

// IStmt is an interface to be used with testing statements
type IStmt interface {
	Close() error
	Exec(args ...interface{}) (sql.Result, error)
	Query(args ...interface{}) (*sql.Rows, error)
	QueryRow(args ...interface{}) *sql.Row
}

// DB is our database type
type DB struct {
	*gorp.DbMap
}

// TX is our transaction type
type TX struct {
	*gorp.Transaction
}

// Stmt is our statement type
type Stmt struct {
	*sql.Stmt
}

// Begin will return an instance of ITX
func (db *DB) Begin() (ITX, error) {
	var tx *TX
	gorpTx, err := db.DbMap.Begin()
	if err != nil {
		return tx, err
	}
	tx.Transaction = gorpTx
	return tx, nil
}

// Prepare will return an instance of Stmt for the database.
func (db *DB) Prepare(query string) (IStmt, error) {
	var stmt *Stmt
	sqlStmt, err := db.DbMap.Prepare(query)
	if err != nil {
		return stmt, err
	}
	stmt.Stmt = sqlStmt
	return stmt, nil
}

// Prepare will return an instance of Stmt for the transaction.
func (tx *TX) Prepare(query string) (IStmt, error) {
	var stmt *Stmt
	sqlStmt, err := tx.Transaction.Prepare(query)
	if err != nil {
		return stmt, err
	}
	stmt.Stmt = sqlStmt
	return stmt, nil
}
