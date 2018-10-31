package repo

import ()

type Contacts struct {
	Repo
}

// new
func (cs *contacts) new(email string) (*Contact, error) {
	contact := &Contact{cs, newID(), email}

	if err := cs.store(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

// store
func (cs *contacts) store(c *Contact) error {
	return nil
}

// fetch
func (cs *contacts) fetch(email string) (*Contact, error) {

}
