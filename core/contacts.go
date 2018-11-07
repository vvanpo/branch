package titian

import ()

type Contacts struct {
	app    *Container
	list   map[id]*Contact
	fields map[id]*Field
}

// New creates a Contact. It must be passed a unique e-mail address. Passing a
// zero-value causes the method to panic.
func (cs *Contacts) New(email EmailAddress) (*Contact, error) {
	contact := &Contact{
		app:        cs.app,
		id:         newID(),
		email:      email,
		alternates: nil,
		fields:     make(map[id]FieldValue),
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

	for _, value := range contact.Fields() {
		contact.DeleteField(value.Field())
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
		make(map[id]*Contact),
		make(map[id]*Field),
	}
}

func (cs Contacts) fetch(contact id) *Contact {
	return cs.list[contact]
}
