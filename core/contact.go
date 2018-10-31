package branch

import ()

type Contact struct {
	contacts
	id
	email string
}

// NewContact
func (b *Branch) NewContact(email string) (Contact, error) {
	contact := Contact{b.contacts, newID(), email}

	if err := b.contacts.store(contact); err != nil {
		return nil, err
	}

	return contact, nil
}
