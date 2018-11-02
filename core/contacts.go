package core

import (
	"errors"
)

type Contacts map[ID]Contact

func (cs Contacts) New(email string) (ID, error) {
	id := newID()
	cs[id] = Contact{id, email}
	return id, nil
}

func (cs Contacts) Fetch(id ID) (*Contact, error) {
	if c, ok := cs[id]; ok {
		return &c, nil
	}

	return nil, errors.New("Invalid contact ID")
}
