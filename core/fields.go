package titian

import ()

type Fields struct {
	categories map[id]*FieldCategory
	fields     map[id]*Field
}

func (fs Fields) fetch(field id) *Field {
	return fs.fields[field]
}

func (fs Fields) New(name string, datatype string) *Field {
	id := newID()
	fs.fields[id] = &Field{id: id, name: name, datatype: datatype}
	return fs.fields[id]
}
