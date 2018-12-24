package titian

import (
	"errors"

	"github.com/vvanpo/titian/field"
)

type Fields struct {
	app *Container
	// A list of top-level categories
	categories []*field.Category
}

// Delete removes a field, deleting all values for that field from all contacts.
func (fs *Fields) Delete(field *Field) {
	for _, group := range fs.app.groups.All() {
		group.RemoveRequiredField(field)
	}

	for _, contact := range fs.app.contacts.All() {
		contact.DeleteField(field)
	}

	*field = Field{}
}

// Move changes a field's category.
func (fs *Fields) Move(field *Field, category *FieldCategory) error {
	if category.Find(field.name) != nil {
		return errors.New("Duplicate category name")
	}

	delete(field.Category().fields, field.id)
	category.fields[field.id] = field
	return nil
}

// Find searches the fields collection for a field.
func (fs *Fields) Find(name string) *Field {
	for _, category := range fs.categories {
		if field := category.Find(name); field != nil {
			return field
		}
	}

	return nil
}

// NewCategory
func (fs *Fields) NewCategory(name string, parent *FieldCategory) (*FieldCategory, error) {
	if fs.FindCategory(name) != nil {
		return nil, errors.New("Duplicate category name")
	}

	category := &FieldCategory{
		app:           fs.app,
		id:            fcid(newID()),
		name:          name,
		fields:        make(map[fid]*Field),
		subcategories: make(map[fcid]*FieldCategory),
	}
	collection := fs.categories

	if parent != nil {
		collection = parent.subcategories
	}

	collection[category.id] = category
	return category, nil
}

//
func (fs *Fields) DeleteCategory(category *FieldCategory) {

}

func (fs *Fields) MoveCategory(category, parent *FieldCategory) {

}

func (fs Fields) FindCategory(name string) *FieldCategory {
	return nil
}

func newFields(app *Container) *Fields {
	return &Fields{
		app,
		make(map[fcid]*FieldCategory),
	}
}
