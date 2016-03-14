package gorpish_test

import (
	"errors"
	"testing"

	_ "github.com/erikstmartin/go-testdb"
	"github.com/simplehq/gorpish"
)

var db *gorpish.TestDB

func init() {
	db = gorpish.NewTestDB()
}

func TestBeginCalled(t *testing.T) {
	tx := gorpish.NewTestTransaction()

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

func TestInsertCalledReturnsErrorAndRollback(t *testing.T) {
	tx := gorpish.NewTestTransaction()

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

func TestOKInsertDoesCommit(t *testing.T) {
	tx := gorpish.NewTestTransaction()

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
