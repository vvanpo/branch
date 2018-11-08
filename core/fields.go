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

	if parent = nil {
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

func (fc *FieldCategory) walkCategories(fn func(*FieldCategory)) {
	for _, category := range fc.subcategories {
		fn(category)
		category.walkCategories(fn)
	}
}
