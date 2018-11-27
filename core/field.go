package titian

import ()

type fid id

// A Field is used to label and format information associated with contacts.
type Field struct {
	app         Container
	id          fid
	name        string
	description string
	fieldtype   FieldType
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
