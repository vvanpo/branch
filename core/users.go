package titian

import ()

type Users map[id]*User

func (us Users) fetch(user id) *User {
	return us[user]
}

func (us Users) New(contact *Contact) *User {
	for _, u := range us {
		if u.id == contact.id {
			panic("Contact already belongs to an existing user")
		}
	}

	id := newID()
	us[id] = &User{id: id, contact: contact}
	return us[id]
}
