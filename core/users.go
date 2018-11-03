package core

import ()

type Users map[ID]*User

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

func (us Users) Fetch(user ID) *User {
	return us[user]
}
