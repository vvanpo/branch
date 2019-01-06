package field

import (
	"errors"
	"fmt"
)

// Categories is a collection that controls persistence of Category objects.
type Categories interface {
	// Query returns a Category and a channel for persisting it. The channel can
	// only be sent to once, after which it is closed. The returned Category is
	// guaranteed to be consistent until it is sent down the channel, or the
	// channel is closed. If the caller does not need to persist the Category it
	// should close the channel once consistency is no longer needed. This frees
	// up the Category for use by another caller.
	Query() (Category, chan<- Category)
}

//
type Category struct {
	repository  Categories
	name        label
	description text
	fields      []*Field
}

// NewCategory
func NewCategory(name, description string, repository Categories) (*Category, error) {
	c := &Category{
		repository: repository,
		fields:     make([]*Field, 0),
	}

	if err := c.name.Set(name); err != nil {
		return nil, err
	}

	if err := c.description.Set(description); err != nil {
		return nil, err
	}

	return c, nil
}

// Name
func (c Category) Name() string {
	return string(c.name)
}

// Description
func (c Category) Description() string {
	return string(c.description)
}

// SetDescription
func (c *Category) SetDescription(description string) error {
	if err := c.description.Set(description); err != nil {
		return err
	}

	c.repository.Persist(c)
	return nil
}

// Fields
func (c *Category) Fields() []*Field {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.repository.Fields().Fetch(c)
}

// Subcategories
func (c *Category) Subcategories() []*Category {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.repository.Fetch(c)
}

// AppendField
func (c *Category) AppendField(field *Field) error {
	if c.GetField(field.Name()) != nil {
		return errors.New("Duplicate field name")
	}

	if c.GetSubcategory(field.Name()) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is already the name of a subcategory", field.Name())
	}

	c.fields = append(c.fields, field)
	return nil
}

// AppendSubcategory
func (c *Category) AppendSubcategory(category *Category) error {
	if c.GetSubcategory(category.Name()) != nil {
		return errors.New("Duplicate subcategory name")
	}

	if c.GetField(category.Name()) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is already the name of a field", category.Name())
	}

	c.subcategories = append(c.subcategories, category)
	return nil
}

// RenameField
func (c *Category) RenameField(from, to string) error {
	if c.GetField(to) != nil {
		return errors.New("Duplicate field name")
	}

	if c.GetSubcategory(to) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is the name of a subcategory", to)
	}

	field := c.GetField(from)

	if field == nil {
		panic("Cannot rename nil field")
	}

	return field.name.Set(to)
}

// RenameSubcategory
func (c *Category) RenameSubcategory(from, to string) error {
	if c.GetSubcategory(to) != nil {
		return errors.New("Duplicate subcategory name")
	}

	if c.GetField(to) != nil {
		return fmt.Errorf("Duplicate name, \"%s\" is already the name of a field", to)
	}

	category := c.GetSubcategory(from)

	if category == nil {
		panic("Cannot rename nil subcategory")
	}

	return category.name.Set(to)
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
	var walk func([]*Category)
	walk = func(categories []*Category) {
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
