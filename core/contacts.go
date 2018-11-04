package core

import ()

type Contacts struct {
	list   map[id]*Contact
	fields map[id]*Field
}

func (cs Contacts) fetch(contact id) *Contact {
	return cs.list[contact]
}

func (cs Contacts) New(email string) *Contact {
	id := newID()
	cs.list[id] = &Contact{
		id:     id,
		emails: []string{email},
	}
	return cs.list[id]
}
