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

type text string

func (t TextType) NewValue() Value {
	return new(text)
}

func (t TextType) Validate() error {
	return nil
}

func (t *text) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	*t = text(value)
	return nil
}
