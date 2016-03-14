## gorpish

gorpish provides a re-usable interface to a gorp databse. It also provides some basic testing utilities.

### Usage

Creating a DB:

```Go
func NewDB() *gorpish.DB {
    sqlDb, _ := sql.Open("postgres", "")
    gorpMap := &gorp.DbMap{Db: sqlDb, Dialect: gorp.PostgresDialect{}}
    return &gorpish.DB{DbMap: gorpMap}
}

db := NewDB()
// Use the db as a normal gorp.DbMap
```

### Testing

Some basic mocked objects are available in the `mocks` package. You can embed them and mock further methods using the api found at https://github.com/stretchr/testify/mock.

```
import (
    "testing"
    "github.com/simplehq/gorpish/mocks"
)

func TestInsert(t *testing.T) {
    db := mocks.NewTestDB()

    db.On("Insert", 1).Return(nil)

    err := db.Insert(1)

    db.AssertExpectations(t)
}
```
