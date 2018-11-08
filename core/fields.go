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
	fs.walkFields(func(field *Field) {
		fields = append(fields, field)
	})
	return fields
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
func (fs *Fields) Move(field *Field, category *FieldCategory) {
	delete(field.Category().fields, field.id)
	category.fields[field.id] = field
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
		id:            newID(),
		name:          name,
		fields:        make(map[id]*Field),
		subcategories: make(map[id]*FieldCategory),
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
		make(map[id]*FieldCategory),
	}
}

// walkCategories applies fn to every category.
func (fs *Fields) walkCategories(fn func(*FieldCategory)) {
	var walk func(map[id]*FieldCategory)
	walk = func(fcs map[id]*FieldCategory) {
		for _, category := range fcs {
			fn(category)
			walk(category.subcategories)
		}
	}

	walk(fs.categories)
}

// walkFields applies fn to every field.
func (fs *Fields) walkFields(fn func(*Field)) {
	fs.walkCategories(func(fc *FieldCategory) {
		for _, field := range fc.fields {
			fn(field)
		}
	})
}
