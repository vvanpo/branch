package titian

import (
	"encoding"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

// FieldValue represents a value for a particular field of a contact.
type FieldValue interface {
	Field() *Field
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	fmt.Stringer
}

// LabelFieldValue
type LabelFieldValue struct {
	id
	field *Field
	value string
}

func (l *LabelFieldValue) Field() *Field {
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

func (l *LabelFieldValue) String() string {
	return value
}

// TextFieldValue
type TextFieldValue struct {
	id
	field *Field
	value string
}

func (t *TextFieldValue) Field() *Field {
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

func (t *TextFieldValue) String() string {
	return value
}

// DateFieldValue
type DateFieldValue struct {
	id
	field *Field
	date  time.Time
}

func (d *DateFieldValue) Field() *Field {
	return d.field
}

func (d *DateFieldValue) MarshalBinary() ([]byte, error) {
	return d.date.MarshalBinary()
}

func (d *DateFieldValue) UnmarshalBinary(data []byte) error {
	return d.date.UnmarshalBinary(data)
}

func (d *DateFieldValue) SetDate(date time.Time) {
	d.date = date
}

func (d *DateFieldValue) String() string {
	return d.date.String()
}
