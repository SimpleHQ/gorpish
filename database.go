package gorper

import (
	"database/sql/driver"

	"gopkg.in/gorp.v1"
)

// IDB is an interface to the databse structure
type IDB interface {
	Begin() (ITX, error)
}

// ITX is an interface to the transaction structure
type ITX interface {
	Insert(...interface{}) error
	driver.Tx
}

// DB is our databse type
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
