package contact

import (
	"github.com/vvanpo/titian/email"
	"github.com/vvanpo/titian/field"
)

type Contacts interface {
	New(email.Address) (*Contact, error)
	Delete(*Contact)
}
