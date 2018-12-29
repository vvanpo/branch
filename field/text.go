package field

import (
	"errors"
	"unicode/utf8"
)

// TextType represents a multi-line text field.
type TextType struct {
	// The max length of the text field, in runes. Zero implies no limit.
	Maximum uint
}

func (t TextType) NewField(name, description string) (*Field, error) {
	return newField(name, description, t)
}

func (t TextType) Validate(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	return nil
}
