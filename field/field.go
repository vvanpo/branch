package field

// A Field is used to categorize and format information associated with
// contacts.
type Field struct {
	name        string
	description string
	definition  Type
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
	t := (TextType{}).NewValue()

	if err := t.(*text).Set(description); err != nil {
		return err
	}

	f.description = description
	return nil
}

func NewField(name, description string, definition Type) (*Field, error) {
	field := &Field{}

	if err := field.setName(name); err != nil {
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

func (f *Field) setName(name string) error {
	l := (LabelType{}).NewValue()

	if err := l.(*label).Set(name); err != nil {
		return err
	}

	f.name = name
	return nil
}
