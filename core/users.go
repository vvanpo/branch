package branch

import (
)

// users is the repository of User objects.
type users struct {
	db Database
}

func (us *users) fetch(email string) (Contact, error) {

}

func (us *users) create(contact Contact) error {

}

