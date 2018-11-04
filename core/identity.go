package core

import (
	"github.com/google/uuid"
)

// Universal indentifier for objects.
type id uuid.UUID

// newID
func newID() id {
	return id(uuid.New())
}
