package titian

import ()

type Container struct {
	contacts *Contacts
	users    *Users
	groups   *Groups
	fields   *Fields
	emails   *Emails
}

func New(config Config) *Container {
	return &Container{
		&Contacts{},
		&Users{},
		&Groups{},
		&Fields{},
		&Emails{},
	}
}
