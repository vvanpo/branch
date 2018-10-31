package branch

import (
	"database/sql"
	"github.com/satori/go.uuid"
)

type id [16]byte

func newID() id {
	// UUID version 4 is the only properly pseudorandom version.
	u := uuid.Must(uuid.NewV4())
	return id(uuid.Bytes())
}
