package titian

import (
	"errors"

	"github.com/google/uuid"
)

//
type FieldCategory struct {
	id            uuid.UUID
	name          string
	description   string
	fields        map[uuid.UUID]*Field
	subcategories map[uuid.UUID]*FieldCategory
}

// NewField
func (fc *FieldCategory) NewField(name string, datatype FieldType) (*Field, error) {
	if fc.app.fields.Find(name) != nil {
		return nil, errors.New("Duplicate field name")
	}

	field := &Field{
		id:        uuid.New(),
		name:      name,
		fieldtype: datatype,
	}

	return field, nil
}

// Find
func (fc *FieldCategory) Find(name string) *Field {
	for _, field := range fc.fields {
		if name == field.name {
			return field
		}
	}

	for _, category := range fc.subcategories {
		if field := category.Find(name); field != nil {
			return field
		}
	}

	return nil
}
