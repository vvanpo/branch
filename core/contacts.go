package core

import ()

type Contacts map[ID]*Contact

func (cs Contacts) New(email string) ID {
	id := newID()
	cs[id] = &Contact{id, email}
	return id
}

func (cs Contacts) Fetch(id ID) *Contact {
	return cs[id]
}
