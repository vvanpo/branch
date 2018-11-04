package titian

import ()

type Field struct {
	app Container
	id
	name        string
	description string
	category    *FieldCategory
	groups      map[id]struct {
		group    *Group
		required bool
	}
	datatype string
	// If access is a nil map, the owning user has full access to their field.
	// Otherwise, user access to their field is determined by group membership.
	access map[*Group]struct {
		readonly bool
	}
	// If the visibility slice is empty, the field defaults to public.
	visibility map[*Group]struct {
		administer bool
	}
}

// Name
func (f Field) Name() string {
	return f.name
}

func (f *Field) SetName(name string) {
	f.name = name
}

// Description
func (f Field) Description() string {
	return f.description
}

func (f *Field) SetDescription(description string) {
	f.description = description
}

// Category
func (f Field) Category() *FieldCategory {
	return f.category
}

func (f *Field) SetCategory(category *FieldCategory) {
	f.category = category
}

// Groups
func (f Field) Groups() []*Group {
	groups := make([]*Group, 0, len(f.groups))

	for _, g := range f.groups {
		groups = append(groups, g.group)
	}

	return groups
}

// Visible
func (f Field) Visible(group *Group) bool {
	if len(f.visibility) == 0 {
		return true
	}

	_, ok := f.visibility[group]
	return ok
}

// Administer
func (f Field) Administer(group *Group) bool {
	if g, ok := f.visibility[group]; ok {
		return g.administer
	}

	return false
}
