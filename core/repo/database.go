package repo

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type ID uuid.UUID

// NewID
func NewID() ID {
	return ID(uuid.New())
}

type db struct {
	*sql.DB
}

// open
func open() db {
	conn, err := sql.Open("postgres", "")
	if err != nil {
		panic("Failed to connect to database")
	}

	return db{conn}
}
