package titian

import ()

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	name        string
	description string
	fieldtype   Type
}

// Name
func (f Field) Name() string {
	return f.name
}

// Description
func (f Field) Description() string {
	return f.description
}

// SetDescription
func (f *Field) SetDescription(description string) {
	f.description = description
}

// Type retrieves the field type.
func (f Field) Type() Type {
	return f.fieldtype
}

func (f *Field) setName(name string) error {
	f.name = name
	return nil
}