package core

import ()

type Groups map[ID]*Group

func (gs Groups) New(name string) *Group {
	id := newID()
	gs[id] = &Group{id: id, name: name}
	return gs[id]
}

func (gs Groups) Fetch(group ID) *Group {
	return gs[group]
}
