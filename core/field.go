package titian

import ()

// A Field is used to label and format information associated with contacts.
type Field struct {
	app Container
	id
	name        string
	description string
	fieldtype   FieldType
	// The list of groups whose members can have a value stored in this field.
	// If this list is empty, the field is common to all contacts.
	groups map[id]*Group
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

func (f Field) Type() FieldType {
	return f.fieldtype
}

// Category
func (f Field) Category() *FieldCategory {
	return nil
}

// Groups
func (f Field) Groups() []*Group {
	groups := make([]*Group, 0, len(f.groups))

	for _, group := range f.groups {
		groups = append(groups, group)
	}

	return groups
}

func (f *Field) AddGroup(group *Group) {
	f.groups[group.id] = group
}

func (f *Field) RemoveGroup(group *Group) {
	delete(f.groups, group.id)
}
