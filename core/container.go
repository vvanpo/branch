package titian

import ()

type Container struct {
	contacts *Contacts
	users    *Users
	groups   *Groups
	fields   *Fields
}

func New(config Config) *Container {
	app := &Container{}
	return app
}
