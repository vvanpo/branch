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
	fields := make([]*Field, 0)
	walkCategories(func(fc *FieldCategory) {
		for _, field := range fc.fields {
			fields = append(fields, field)
		}
	})
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

	*field = Field{}
}

func (fs *Fields) Move(field *Field, category *FieldCategory) {

}

func (fs *Fields) Find(name string) *Field {
	for _, category := range fs.categories {
		if field := category.Find(name); field != nil {
			return field
		}
	}
}

// NewCategory
func (fs *Fields) NewCategory(name string, parent *FieldCategory) (*FieldCategory, error) {
	if fs.FindCategory(name) != nil {
		return nil, errors.New("Duplicate category name")
	}

	category := &FieldCategory{
		app:           fs.app,
		id:            newId(),
		name:          name,
		fields:        make(map[id]*Field),
		subcategories: make(map[id]*Field),
	}

	if parent == nil {
		parent = fs.categories
	}

	parent.subcategories[category.id] = category
	return category, nil
}

//
func (fs *Fields) DeleteCategory(category *FieldCategory) {

}

func (fs *Fields) MoveCategory(category, parent *FieldCategory) {

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

// walkCategories applies fn to every category.
func (fs *Fields) walkCategories(fn func(*FieldCategory)) {
	walk := func(fcs []*FieldCategory) {
		for _, category := range fcs {
			fn(category)
			walk(fcs.subcategories)
		}
	}

	walk(fs.categories)
}
