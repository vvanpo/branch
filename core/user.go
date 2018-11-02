package core

import ()

type User struct {
	id ID
	contact *Contact
	roles map[ID]*Role
}

func (u User) ID() ID {
	return u.id
}

func (u User) Contact() *Contact {
	return u.contact
}

func (u User) Groups() map[ID]*Group {
	groups := make(map[ID]*Group)
	return groups
}
