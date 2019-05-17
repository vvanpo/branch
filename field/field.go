// Package field
//
package field

import ()

type Interface interface {
	TypeName() string
	Name() string
	Description() string
}

type field struct {
	name        string
	description string
}

func (f field) Name() string {
	return f.name
}

func (f field) Description() string {
	return f.description
}

/**
// NewField
func NewField(name, description string, definition Type, repository Fields) (*Field, error) {
	field := &Field{
		repo: repository,
	}

	if err := field.name.Set(name); err != nil {
		return nil, err
	}

	if err := field.description.Set(description); err != nil {
		return nil, err
	}

	if err := definition.Validate(); err != nil {
		return nil, err
	}

	field.definition = definition
	return field, nil
}

// Name
func (f Field) Name() string {
	return string(f.name)
}

// Description
func (f Field) Description() string {
	return string(f.description)
}

// SetDescription
func (f *Field) SetDescription(description string) error {
	if err := f.description.Set(description); err != nil {
		return err
	}

	f.repo.Persist(f)
	return nil
}
*/
