package stonefruit

import (
	"github.com/google/uuid"
)

// Universal indentifier for objects.
type id uuid.UUID

// newID generates a unique identifier.
func newID() id {
	return id(uuid.New())
}
