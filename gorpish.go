package gorpish

import (
	"database/sql"
	"database/sql/driver"

	"github.com/stretchr/testify/mock"

	"gopkg.in/gorp.v1"
)

// IDB is an interface to the database structure
type IDB interface {
	Begin() (ITX, error)
}

// ITX is an interface to the transaction structure
type ITX interface {
	Insert(...interface{}) error
	driver.Tx
}

// DB is our database type
type DB struct {
	IDB
	*gorp.DbMap
}

// TX is our transaction type
type TX struct {
	*gorp.Transaction
}

// Begin will return an instance of ITX
func (db *DB) Begin() (ITX, error) {
	var tx TX
	gorpTx, err := db.DbMap.Begin()
	if err != nil {
		return nil, err
	}
	tx.Transaction = gorpTx
	return &tx, nil
}

// TestDB is our test database.
// It is mockable, and implements DB
type TestDB struct {
	mock.Mock
	*DB
}

// Begin implements the IDB Begin method.
func (db *TestDB) Begin() (ITX, error) {
	args := db.Called()
	return args.Get(0).(ITX), args.Error(1)
}

// TestTransaction is our test transaction.
// It is mockable and implements TX
type TestTransaction struct {
	mock.Mock
	*TX
}

// Insert implements ITX Insert method.
func (tx *TestTransaction) Insert(list ...interface{}) error {
	args := tx.Called(list...)
	return args.Error(0)
}

// Rollback implements the ITX Rollback method.
func (tx *TestTransaction) Rollback() error {
	args := tx.Called()
	return args.Error(0)
}

// Commit implements the ITX Commit method.
func (tx *TestTransaction) Commit() error {
	args := tx.Called()
	return args.Error(0)
}

// NewTestDB will create a new test database.
func NewTestDB() *TestDB {
	sqlDb, _ := sql.Open("testdb", "")

	gorpMap := &gorp.DbMap{Db: sqlDb, Dialect: gorp.PostgresDialect{}}

	newDB := new(DB)
	newDB.IDB = newDB
	newDB.DbMap = gorpMap

	db := new(TestDB)
	db.DB = newDB

	return db
}

// NewTestTransaction will create a empty transaction
// ready for testing.
func NewTestTransaction() *TestTransaction {
	tx := new(TX)
	tx.Transaction = &gorp.Transaction{}

	testTx := new(TestTransaction)
	testTx.TX = tx

	return testTx
}
