package core

import (
	"errors"
)

type Users map[ID]User

func (us Users) New(contact ID) (ID, error) {
	var id ID

	for _, u := range us {
		if u.id == contact {
			return id, errors.New("Contact already belongs to an existing user")
		}
	}

	id = newID()
	us[id] = User{id: id, contact: contact}
	return id, nil
}
