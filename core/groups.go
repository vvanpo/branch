package titian

import ()

type Groups struct {
	app  *Container
	list map[id]*Group
}

func (gs Groups) All() []*Group {
	groups := make([]*Group, 0, len(gs.list))

	for _, group := range gs.list {
		groups = append(groups, group)
	}

	return groups
}

func (gs Groups) New(name string) *Group {
	group := &Group{
		app:            gs.app,
		id:             newID(),
		name:           name,
		contacts:       make(map[id]*Contact),
		permissions:    make([]string, 0),
		requiredFields: make(map[id]*Field),
	}
	gs.list[group.id] = group
	return group
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

func (gs Groups) fetch(group id) *Group {
	return gs.list[group]
}
