package field

import (
	"strings"
)

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	name        string
	description string
	fieldtype   Type
}

// Name
func (f Field) Name() string {
	return f.name
}

// Description
func (f Field) Description() string {
	return f.description
}

// SetDescription
func (f *Field) SetDescription(description string) error {
	var t textType
	v, err := t.NewValue(description)

	if err != nil {
		return err
	}

	f.description = v.String()
	return nil
}

// Type retrieves the field type.
func (f Field) Type() Type {
	return f.fieldtype
}

func newField(name, description string, fieldtype Type) (*Field, error) {
	field := &Field{}

	if err := field.setName(name); err != nil {
		return nil, err
	}

	if err := field.SetDescription(description); err != nil {
		return nil, err
	}

	field.fieldtype = fieldtype
	return field, nil
}

func (f *Field) setName(name string) error {
	var t labelType
	v, err := t.NewValue(name)

	if err != nil {
		return err
	}

	f.name = v.String()
	return nil
}
