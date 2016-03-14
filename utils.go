package gorper

import (
	"github.com/stretchr/testify/mock"
)

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
