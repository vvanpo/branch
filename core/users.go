package titian

import ()

type Users struct {
	app  *Container
	list map[uid]*User
}

func (us *Users) New(contact *Contact) *User {
	if us.list == nil {
		us.list = make(map[uid]*User)
	}

	for _, u := range us.list {
		if u.contact.id == contact.id {
			panic("Contact already belongs to an existing user")
		}
	}

	id := uid(newID())
	us.list[id] = &User{
		id:      id,
		contact: contact,
	}
	return us.list[id]
}

func (us Users) user(id uid) *User {
	if user, ok := us.list[id]; ok {
		return user
	}

	return nil
}
