package titian

import ()

type Groups struct {
	app  *Container
	list map[gid]*Group
}

// All returns a list of all groups.
func (gs Groups) All() []*Group {
	groups := make([]*Group, 0, len(gs.list))

	for _, group := range gs.list {
		groups = append(groups, group)
	}

	return groups
}

// New creates a new group and adds it to the collection.
func (gs *Groups) New(name string) *Group {
	group := &Group{
		app:            gs.app,
		gid:             gid(newID()),
		name:           name,
		contacts:       make(map[id]*Contact),
		permissions:    make([]string, 0),
		requiredFields: make(map[id]*Field),
	}
	gs.list[group.gid] = group
	return group
}

// Delete removes a group from the collection and deletes associated data.
func (gs *Groups) Delete(group *Group) {
	delete(gs.list, group.gid)
	*group = Group{}
}

func newGroups(app *Container) *Groups {
	return &Groups{
		app,
		make(map[gid]*Group),
	}
}

func (gs Groups) fetch(group gid) *Group {
	return gs.list[group]
}
