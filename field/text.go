package field

import (
	"errors"
	"unicode/utf8"
)

func NewText(name, description string) (*Field, error) {
	return newField(name, description, &textType{})
}

// text represents a multi-line text field.
type text struct{}

func (_ text) NewValue(value string) (Value, error) {
	if !utf8.ValidString(value) {
		return nil, errors.New("Invalid string encoding")
	}

	return textValue(value), nil
}

// textValue
type textValue string

func (t textValue) String() string {
	return string(t)
}
