package titian

import (
	"github.com/google/uuid"
)

// A Field is used to label and format information associated with contacts.
type Field struct {
	id          uuid.UUID
	name        string
	description string
	fieldtype   FieldType
}

// Name
func (f Field) Name() string {
	return f.name
}

// SetName
func (f *Field) SetName(name string) {
	f.name = name
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
func (f Field) Type() FieldType {
	return f.fieldtype
}
