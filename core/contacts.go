package titian

import ()

type Contacts struct {
	app    *Container
	list   map[id]*Contact
	fields map[id]*Field
}

func (cs Contacts) fetch(contact id) *Contact {
	return cs.list[contact]
}

func (cs *Contacts) New(email string) *Contact {
	if cs.list == nil {
		cs.list = make(map[id]*Contact)
	}

	id := newID()
	cs.list[id] = &Contact{
		id:     id,
		emails: []string{email},
	}
	return cs.list[id]
}

func (cs *Contacts) Find(email string) *Contact {
	for _, contact := range cs.list {
		for _, e := range contact.VerifiedEmails() {
			if email == e {
				return contact
			}
		}
	}

	return nil
}
