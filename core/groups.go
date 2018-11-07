package titian

import ()

type Groups struct {
	app  *Container
	list map[id]*Group
}

func (gs Groups) fetch(group id) *Group {
	return gs.list[group]
}

func (gs Groups) New(name string) *Group {
	if gs.list == nil {
		gs.list = make(map[id]*Group)
	}

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

func newGroups(app *Container) *Groups {
	return &Groups{
		app,
		make(map[id]*Group),
	}
}
