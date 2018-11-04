package core

import ()

type Contacts map[id]*Contact

func (cs Contacts) fetch(contact id) *Contact {
	return cs[contact]
}

func (cs Contacts) New(email string) *Contact {
	id := newID()
	cs[id] = &Contact{id, email}
	return cs[id]
}
