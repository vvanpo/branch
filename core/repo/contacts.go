package repo

import ()

// Contact
type Contact struct {
	ID
	Email string
}

// Contacts
type Contacts struct {
	db
}

// New
func (cs *Contacts) New(email string) (Contact, error) {
	contact := Contact{NewID(), email}

	return contact, nil
}

// Save
func (cs *Contacts) Save(c Contact) error {
	return nil
}

// Fetch
func (cs *Contacts) Fetch(email string) Contact {
	return Contact{}
}
