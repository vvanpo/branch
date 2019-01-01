package field

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	name        label
	description text
	definition  Type
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

	return nil
}

func NewField(name, description string, definition Type) (*Field, error) {
	field := &Field{}

	if err := field.name.Set(name); err != nil {
		return nil, err
	}

	if err := field.SetDescription(description); err != nil {
		return nil, err
	}

	if err := definition.Validate(); err != nil {
		return nil, err
	}

	field.definition = definition
	return field, nil
}
