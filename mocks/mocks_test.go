package mocks_test

import (
	"errors"
	"testing"

	_ "github.com/erikstmartin/go-testdb"
	"github.com/simplehq/gorpish/mocks"
)

var db *mocks.TestDB

func init() {
	db = mocks.NewTestDB()
}

// TestBeginCalled will assert Begin was called
// on the TestDB.
func TestBeginCalled(t *testing.T) {
	tx := mocks.NewTestTX()

	db.On("Begin").Return(tx, nil).Once()

	newTx, err := db.Begin()
	if err != nil {
		t.Error("Error was not nil.")
	}

	if newTx == nil {
		t.Error("Transaction should not be nil.")
	}

	db.AssertExpectations(t)
}

// TestInsertCalledReturnsErrorAndRollback will assert
// that Insert returns and error and calls Rollback
// on the TestTx.
func TestInsertCalledReturnsErrorAndRollback(t *testing.T) {
	tx := mocks.NewTestTX()

	tx.On("Insert", "random").Return(errors.New("Did not work"))
	tx.On("Rollback").Return(nil).Once()

	err := tx.Insert("random")
	if err == nil {
		t.Error("Insert error should not be nil.")
	}

	err = tx.Rollback()
	if err != nil {
		t.Error("Rollback error should be nil.")
	}

	tx.AssertExpectations(t)
}

// TestOKInsertDoesCommit will assert that a
// successful Insert will Commit the TestTx.
func TestOKInsertDoesCommit(t *testing.T) {
	tx := mocks.NewTestTX()

	tx.On("Insert", "random").Return(nil)
	tx.On("Commit").Return(nil).Once()

	err := tx.Insert("random")
	if err != nil {
		t.Error("Insert error should be nil.")
	}

	err = tx.Commit()
	if err != nil {
		t.Error("Commit error should be nil.")
	}

	tx.AssertExpectations(t)
}

// TestDbPrepare will assert that Prepare
// returns a TestStmt from the TestDB.
func TestDbPrepare(t *testing.T) {
	stmt := mocks.NewTestStmt()
	query := "SELECT * FROM this"

	db.On("Prepare", query).Return(stmt, nil)

	newStmt, err := db.Prepare(query)
	if err != nil {
		t.Error("Prepare error should be nil.")
	}
	if newStmt == nil {
		t.Error("Prepare statement should not be nil.")
	}

	db.AssertExpectations(t)
}

// TestStmtExec will assert that Exec
// is called on the TestStmt.
func TestStmtExec(t *testing.T) {
	result := mocks.TestResult{}
	stmt := mocks.NewTestStmt()

	stmt.On("Exec", 1).Return(result, nil)

	newResult, err := stmt.Exec(1)
	if err != nil {
		t.Error("Exec error should be nil.")
	}
	if _, ok := newResult.(mocks.TestResult); !ok {
		t.Error("Exec result is not of type mocks.TestResult")
	}

	stmt.AssertExpectations(t)
}
