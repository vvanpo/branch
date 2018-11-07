package titian

import ()

type Fields struct {
	app        *Container
	categories map[id]*FieldCategory
	list       map[id]*Field
}

func (fs Fields) All() []*Field {
	fields := make([]*Field, 0, len(fs.list))

	for _, field := range fs.list {
		fields = append(fields, field)
	}

	return fields
}

func (fs *Fields) New(name string, datatype FieldType) *Field {
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

// Delete removes a field, deleting all values for that field from contacts.
func (fs *Fields) Delete(field *Field) {
	for _, group := range fs.app.groups.All() {
		group.RemoveRequiredField(field)
	}

	for _, contact := range fs.app.contacts.All() {
		contact.DeleteField(field)
	}

	delete(fs.list, field.id)
	*field = Field{}
}

func newFields(app *Container) *Fields {
	return &Fields{
		app,
		make(map[id]*FieldCategory),
		make(map[id]*Field),
	}
}

func (fs Fields) fetch(field id) *Field {
	return fs.list[field]
}
