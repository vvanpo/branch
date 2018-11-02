package core

import (
	"github.com/google/uuid"
)

// Universal indentifier for objects.
type ID uuid.UUID

// newID
func newID() ID {
	return ID(uuid.New())
}
