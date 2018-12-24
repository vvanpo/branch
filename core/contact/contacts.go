package titian

import ()

type Contacts struct {
	app    *Container
	list   map[cid]*Contact
	fields map[fid]*Field
}

func (cs *Contacts) All() []*Contact {
	contacts := make([]*Contact, 0, len(cs.list))

	for _, contact := range cs.list {
		contacts = append(contacts, contact)
	}

	return contacts
}

// New creates a Contact. It must be passed a unique e-mail address. Passing a
// zero-value causes the method to panic.
func (cs *Contacts) New(email EmailAddress) (*Contact, error) {
	contact := &Contact{
		app:        cs.app,
		id:         cid(newID()),
		email:      email,
		alternates: nil,
		fields:     make(map[fid]FieldValue),
	}

	if err := contact.AddEmailAddress(email); err != nil {
		return nil, err
	}

	contact.alternates = nil
	cs.list[contact.id] = contact
	return contact, nil
}

// Delete removes a contact from the list of contacts, and deletes any
// information unique to the contact. The passed Contact object is set to the
// zero value.
func (cs *Contacts) Delete(contact *Contact) {
	for _, group := range contact.Groups() {
		group.RemoveContact(contact)
	}

	for _, field := range contact.Fields() {
		contact.DeleteField(field)
	}

	delete(cs.list, contact.id)
	*contact = Contact{}
}

// Find searches the contact list for a contact with a matching associated
// e-mail address.
func (cs *Contacts) Find(email EmailAddress) *Contact {
	for _, contact := range cs.list {
		for _, e := range contact.EmailAddresses() {
			if email == e {
				return contact
			}
		}
	}

	return nil
}

func newContacts(app *Container) *Contacts {
	return &Contacts{
		app,
		make(map[cid]*Contact),
		make(map[fid]*Field),
	}
}

func (cs Contacts) contact(id cid) *Contact {
	if contact, ok := cs.list[id]; ok {
		return contact
	}

	return nil
}
