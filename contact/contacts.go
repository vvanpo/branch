package contact

import (
	"github.com/vvanpo/stonefruit/email"
	"github.com/vvanpo/stonefruit/field"
)

type Contacts interface {
	New(email.Address) (*Contact, error)
	Delete(*Contact)
}
