package mocks

import (
	"database/sql"

	"github.com/simplehq/gorpish"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gorp.v1"
)

// TestDB is our test database.
// It is mockable, and implements DB
type TestDB struct {
	mock.Mock
	*gorpish.DB
}

// Begin implements the IDB Begin method.
func (db *TestDB) Begin() (gorpish.ITX, error) {
	args := db.Called()
	return args.Get(0).(gorpish.ITX), args.Error(1)
}

// Prepare implements the DB Prepare method
func (db *TestDB) Prepare(query string) (gorpish.IStmt, error) {
	args := db.Called(query)
	return args.Get(0).(gorpish.IStmt), args.Error(1)
}

// TestTX is our test transaction.
// It is mockable and implements TX
type TestTX struct {
	mock.Mock
	*gorpish.TX
}

// Insert implements ITX Insert method.
func (tx *TestTX) Insert(list ...interface{}) error {
	args := tx.Called(list...)
	return args.Error(0)
}

// Rollback implements the ITX Rollback method.
func (tx *TestTX) Rollback() error {
	args := tx.Called()
	return args.Error(0)
}

// Commit implements the ITX Commit method.
func (tx *TestTX) Commit() error {
	args := tx.Called()
	return args.Error(0)
}

// TestStmt is our test statement.
// It is mockable and implements Stmt
type TestStmt struct {
	mock.Mock
	*gorpish.Stmt
}

// Exec runs the statement
func (stmt *TestStmt) Exec(execArgs ...interface{}) (sql.Result, error) {
	args := stmt.Called(execArgs...)
	return args.Get(0).(sql.Result), args.Error(1)
}

// NewTestDB will create a new test database.
func NewTestDB(dialect gorp.Dialect) *TestDB {
	sqlDb, _ := sql.Open("testdb", "")

	gorpMap := &gorp.DbMap{Db: sqlDb, Dialect: dialect}

	newDB := &gorpish.DB{DbMap: gorpMap}

	return &TestDB{DB: newDB}
}

// TestResult gives us a simple struct to use when expecting results.
type TestResult struct {
	sql.Result
}

// NewTestTX will create a empty transaction
// ready for testing.
func NewTestTX() *TestTX {
	tx := &gorpish.TX{Transaction: &gorp.Transaction{}}

	return &TestTX{TX: tx}
}

// NewTestStmt will create a empty Stmt
// ready for testing.
func NewTestStmt() *TestStmt {
	stmt := &gorpish.Stmt{Stmt: &sql.Stmt{}}

	return &TestStmt{Stmt: stmt}
}
