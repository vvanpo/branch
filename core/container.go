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
	app := &Container{}
	app.contacts = newContacts(app)
	app.groups = newGroups(app)
	app.fields = newFields(app)
	return app
}
