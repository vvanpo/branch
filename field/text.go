package field

import (
	"errors"
	"unicode/utf8"
)

func NewTextField(name, description string) (*Field, error) {
	return newField(name, description, &textType{})
}

// textType represents a multi-line text field.
type textType struct{}

func (_ textType) NewValue(value string) (Value, error) {
	if !utf8.ValidString(value) {
		return nil, errors.New("Invalid string encoding")
	}

	return textValue(value), nil
}

// text
type text string

func (t text) String() string {
	return string(t)
}
