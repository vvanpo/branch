package core

import ()

type Container struct {
	Contacts
	Users
	Groups
	Fields
}

func New(config Config) *Container {
	app := &Container{}
	return app
}
