package branch

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type id uuid.UUID
type db struct {
	*sql.DB
}

// newID
func newID() id {
	return id(uuid.New())
}

// openDB
func openDB() db {
	d, err := sql.Open("postgres", "")
	if err != nil {
		panic("Failed to connect to database")
	}

	return db{d}
}
