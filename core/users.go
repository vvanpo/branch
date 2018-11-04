package titian

import ()

type Users struct {
	app  *Container
	list map[id]*User
}

func (us Users) fetch(user id) *User {
	return us.list[user]
}

func (us *Users) New(contact *Contact) *User {
	for _, u := range us.list {
		if u.id == contact.id {
			panic("Contact already belongs to an existing user")
		}
	}

	id := newID()
	us.list[id] = &User{
		id:      id,
		contact: contact,
	}
	return us.list[id]
}
