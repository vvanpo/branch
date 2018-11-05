package titian

import ()

type Fields struct {
	app        *Container
	categories map[id]*FieldCategory
	list       map[id]*Field
}

func (fs Fields) fetch(field id) *Field {
	return fs.list[field]
}

func (fs Fields) New(name string, datatype FieldType) *Field {
	if fs.list == nil {
		fs.list = make(map[id]*Field)
	}

	id := newID()
	fs.list[id] = &Field{
		id:       id,
		name:     name,
		datatype: datatype,
	}
	return fs.list[id]
}
