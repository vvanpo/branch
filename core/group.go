package core

import ()

type Group struct {
	id       ID
	name string
	contacts []ID
	users    []ID
	roles    []ID
}

func (g Group) ID() ID {
	return g.id
}

func (g Group) Name() string {
	return g.name
}

func (g Group) Contacts() []ID {
	return g.contacts[:]
}

func (g Group) Users() []ID {
	return g.users[:]
}

func (g Group) Roles() []ID {
	return g.roles[:]
}
