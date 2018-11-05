package titian

import (
	"encoding"
	"errors"
	"unicode/utf8"
)

// FieldValue represents a value for a particular field of a contact.
type FieldValue interface {
	Field() Field
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// LabelFieldValue
type LabelFieldValue struct {
	id
	field Field
	value string
}

func (l *LabelFieldValue) Field() Field {
	return l.field
}

func (l *LabelFieldValue) MarshalBinary() ([]byte, error) {
	return []byte(l.value), nil
}

func (l *LabelFieldValue) UnmarshalBinary(data []byte) error {
	if !utf8.Valid(data) {
		return errors.New("Invalid UTF-8")
	}

	l.value = string(data)
	return nil
}

// TextFieldValue
type TextFieldValue struct {
	id
	field Field
	value string
}

func (t *TextFieldValue) Field() Field {
	return t.field
}

func (t *TextFieldValue) MarshalBinary() ([]byte, error) {
	return []byte(t.value), nil
}

func (t *TextFieldValue) UnmarshalBinary(data []byte) error {
	if !utf8.Valid(data) {
		return errors.New("Invalid UTF-8")
	}

	t.value = string(data)
	return nil
}
