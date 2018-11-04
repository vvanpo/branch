package titian

import ()

type Groups struct {
	app  Container
	list map[id]*Group
}

func (gs Groups) fetch(group id) *Group {
	return gs.list[group]
}

func (gs Groups) New(name string) *Group {
	id := newID()
	gs.list[id] = &Group{id: id, name: name}
	return gs.list[id]
}

func (gs *Groups) Remove(name string) {
	for _, group := range gs.list {
		if group.name == name {
			delete(gs.list, group.id)
			return
		}
	}
}
