package titian

import (
	"errors"
	"fmt"
)

//
type Category struct {
	name          string
	description   string
	fields        []*Field
	subcategories []*Category
}

// NewCategory
func NewCategory(name, description string) (*Category, error) {
	c := &Category{}

	if err := c.SetName(name); err != nil {
		return nil, err
	}

	if err := c.SetDescription(description); err != nil {
		return nil, err
	}

	c.fields = make([]*Field, 0)
	c.subcategories = make([]*Category, 0)
	return c, nil
}

// Name
func (c Category) Name() string {
	return c.name
}

// SetName
func (c *Category) SetName(name string) error {
	c.name = name
	return nil
}

// Description
func (c Category) Description() string {
	return c.description
}

// SetDescription
func (c *Category) SetDescription(description string) error {
	c.description = description
	return nil
}

// Fields
func (c Category) Fields() []*Field {
	return c.fields[:]
}

// Subcategories
func (c Category) Subcategories() []*Category {
	return c.subcategories[:]
}

// GetField retrieves the field with name if it is a member.
func (c Category) GetField(name string) *Field {
	for _, field := range c.fields {
		if name == field.Name() {
			return field
		}
	}

	return nil
}

// GetSubcategory retrieves the category with name if it is a member.
func (c Category) GetSubcategory(name string) *Category {
	for _, category := range c.subcategories {
		if name == category.Name() {
			return category
		}
	}

	return nil
}

// AppendField
func (c *Category) AppendField(field *Field) error {
	if c.GetField(name) != nil {
		return errors.New("Duplicate field name")
	}

	if c.GetSubcategory(name) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is already the name of a subcategory", name)
	}

	c.fields = append(c.fields, field)
	return nil
}

// AppendSubcategory
func (c *Category) AppendSubcategory(category *Category) error {
	if c.GetSubcategory(name) != nil {
		return errors.New("Duplicate category name")
	}

	if c.GetField(name) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is already the name of a field", name)
	}

	c.subcategories = append(c.subcategories, category)
	return nil
}

// MoveField
func (c *Category) MoveField(from, dest uint) {
	field := c.fields[from]
	c.RemoveField(field)
	c.fields = append(c.fields[:dest], field, c.fields[dest:]...)
}

// MoveSubcategory
func (c *Category) MoveSubcategory(from, dest uint) {
	category := c.subcategories[from]
	c.RemoveSubcategory(category)
	c.subcategories = append(c.subcategories[:dest], category, c.subcategories[dest:]...)
}

// RemoveField
func (c *Category) RemoveField(field *Field) {
	for i, f := range c.fields {
		if f == field {
			c.fields = append(c.fields[:i], c.fields[i+1:]...)
		}
	}
}

// RemoveSubcategory
func (c *Category) RemoveSubcategory(category *Category) {
	for i, sc := range c.subcategories {
		if sc == category {
			c.subcategories = append(c.subcategories[:i], c.subcategories[i+1:]...)
		}
	}
}

// WalkCategories applies fn recursively to each category in the tree, starting
// with the receiver and proceeding in a breadth-first manner.
func (c *Category) WalkCategories(fn func(*Category)) {
	//	var walk func([]*Category)
	var walk = func(categories []*Category) {
		for _, category := range categories {
			fn(category)
			walk(category.subcategories)
		}
	}

	fn(c)
	walk(c.subcategories)
}

// WalkFields applies fn to every field in the receiver, and recursively to
// every field in the receiver's subcategories.
func (c *Category) WalkFields(fn func(*Field)) {
	c.WalkCategories(func(category *Category) {
		for _, field := range category.fields {
			fn(field)
		}
	})
}
