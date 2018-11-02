package core

import ()

type User struct {
	id ID
	contact ID
	roles []ID
}

func (u User) ID() ID {
	return u.id
}

func (u User) Contact() ID {
	return u.contact
}

