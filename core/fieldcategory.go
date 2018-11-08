package titian

import ()

//
type FieldCategory struct {
	app Container
	id
	name          string
	description   string
	fields        map[id]*Field
	subcategories map[id]*FieldCategory
}

//
func (fc *FieldCategory) NewField(name string, datatype FieldType) (*Field, error) {
	field := &Field{
		id:         newID(),
		name:       name,
		datatype:   datatype,
		groups:     make(map[*Group]struct{ bool }),
		userAccess: make(map[*Group]struct{ bool }),
		visibility: make(map[*Group]struct{ bool }),
	}

	return field
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
}
