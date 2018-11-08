package titian

import (
	"errors"
)

type Fields struct {
	app *Container
	// A list of top-level categories
	categories map[id]*FieldCategory
}

func (fs Fields) All() []*Field {
	fields := make([]*Field, 0, len(fs.list))

	for _, field := range fs.list {
		fields = append(fields, field)
	}

	return fields
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

// NewCategory
func (fs *Fields) NewCategory(name string, parent *FieldCategory) (*FieldCategory, error) {
	if fs.FindCategory(name) != nil {
		return nil, errors.New()
	}

	category := &FieldCategory{
		app:           fs.app,
		id:            newId(),
		name:          name,
		fields:        make(map[id]*Field),
		subcategories: make(map[id]*Field),
	}

	parent.subcategories[category.id] = category
	return category, nil
}

func (fs Fields) FindCategory(name string) *FieldCategory {
}

func newFields(app *Container) *Fields {
	return &Fields{
		app,
		make(map[id]*FieldCategory),
		make(map[id]*Field),
	}
}

func (fc *FieldCategory) walkCategories(fn func(*FieldCategory)) {
	for _, category := range fc.subcategories {
		fn(category)
		category.walkCategories(fn)
	}
}
