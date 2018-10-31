package branch

import ()

// users is the repository of User objects.
type users struct {
	*db
}

func (us *users) fetch(email string) (*Contact, error) {
	return nil, nil
}

func (us *users) store(contact *Contact) error {
	return nil
}
