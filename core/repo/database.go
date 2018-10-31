package repo

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type id uuid.UUID
type Repo struct {
	db *sql.DB
}

// newID
func newID() id {
	return id(uuid.New())
}

// openDB
func Open() Repo {
	conn, err := sql.Open("postgres", "")
	if err != nil {
		panic("Failed to connect to database")
	}

	return Repo{conn}
}
