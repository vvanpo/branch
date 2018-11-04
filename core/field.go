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
	access map[id]struct {
		group    *Group
		readonly bool
	}
	datatype string
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Description() string {
	return f.description
}

func (f Field) Category() *FieldCategory {
	return f.category
}

func (f *Field) SetName(name string) {
	f.name = name
}

func (f *Field) SetDescription(description string) {
	f.description = description
}

func (f *Field) SetCategory(category *FieldCategory) {
	f.category = category
}
