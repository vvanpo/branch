package titian

import (
	"errors"
)

type fcid id

//
type FieldCategory struct {
	app           *Container
	id            fcid
	name          string
	description   string
	fields        map[fid]*Field
	subcategories map[fcid]*FieldCategory
}

func (fc *FieldCategory) NewField(name string, datatype FieldType) (*Field, error) {
	if fc.app.fields.Find(name) != nil {
		return nil, errors.New("Duplicate field name")
	}

	field := &Field{
		id:        fid(newID()),
		name:      name,
		fieldtype: datatype,
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
