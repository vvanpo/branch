package titian

import ()

type uid id

//
type User struct {
	app     Container
	id      uid
	contact *Contact
}

//
func (u User) Contact() *Contact {
	return u.contact
}

