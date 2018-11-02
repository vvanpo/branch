package core

import ()

type Role struct {
	id          ID
	name        string
	permissions []string
}

func (r Role) ID() ID {
	return r.id
}

type Roles map[ID]Role

func (rs Roles) New(name string) ID {
	id := newID()
	rs[id] = Role{id: id, name: name}
	return id
}
