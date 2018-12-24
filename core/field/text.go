package titian

import (
	"errors"
	"unicode/utf8"
)

// Text represents a multi-line text field.
type Text struct{}

func (_ Text) Name() string {
	return "Text"
}

func (_ Text) NewValue(value string) (Value, error) {
	text := &TextValue{}

	if err := text.Set(value); err != nil {
		return nil, err
	}

	return text, nil
}

// TextValue
type TextValue struct {
	value string
}

func (t *TextValue) String() string {
	return t.value
}

func (t *TextValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	t.value = value
	return nil
}
