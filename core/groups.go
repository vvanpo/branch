package titian

import ()

type Groups map[id]*Group

func (gs Groups) fetch(group id) *Group {
	return gs[group]
}

func (gs Groups) New(name string) *Group {
	id := newID()
	gs[id] = &Group{id: id, name: name}
	return gs[id]
}

func (gs *Groups) Remove(name string) {
	for _, group := range gs {
		if group.name == name {
			delete(gs, group.id)
			return
		}
	}
}
