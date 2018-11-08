package titian

import (
	"errors"
)

//
type FieldCategory struct {
	app *Container
	id
	name          string
	description   string
	fields        map[id]*Field
	subcategories map[id]*FieldCategory
}

//
func (fc *FieldCategory) NewField(name string, datatype FieldType) (*Field, error) {
	if fc.app.fields.Find(name) != nil {
		return nil, errors.New("Duplicate field name")
	}

	field := &Field{
		id:         newID(),
		name:       name,
		fieldtype:  datatype,
		groups:     make(map[id]*Group),
		visible:    make(map[id]*Group),
		administer: make(map[id]*Group),
	}

	return field, nil
}

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
