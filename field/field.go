package field

import ()

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	name        string
	description string
	fieldtype   interface{}
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
	if err := (TextType{}).Validate(description); err != nil {
		return err
	}

	f.description = description
	return nil
}

func newField(name, description string, fieldtype interface{}) (*Field, error) {
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
	if err := (LabelType{}).Validate(name); err != nil {
		return err
	}

	f.name = name
	return nil
}
