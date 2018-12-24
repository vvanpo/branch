package titian

import (
	"errors"
	"unicode/utf8"
)

func NewText(name, description string) (*Field, error) {
	return newField(name, description, &textType{})
}

// textType represents a multi-line text field.
type textType struct{}

func (_ textType) NewValue(value string) (Value, error) {
	text := new(textValue)

	if err := text.Set(value); err != nil {
		return nil, err
	}

	return text, nil
}

// textValue
type textValue string

func (t *textValue) String() string {
	return string(*t)
}

func (t *textValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	*t = textValue(value)
	return nil
}
