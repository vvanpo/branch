package titian

import (
	"encoding"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

type fvid id

// FieldValue represents a value for a particular field type.
type FieldValue interface {
	fmt.Stringer
	Set(string) error
}

// LabelFieldValue
type LabelFieldValue struct {
	id    fvid
	value string
}

func (l *LabelFieldValue) String() string {
	return l.value
}

func (l *LabelFieldValue) Set(value string) error {
	if !utf8.Valid(value) {
		return errors.New("Invalid string encoding")
	}

	l.value = value
	return nil
}

// TextFieldValue
type TextFieldValue struct {
	id    fvid
	value string
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
	return t.value
}

// DateFieldValue
type DateFieldValue struct {
	id        fvid
	fieldtype FieldType
	date      time.Time
}

func (d *DateFieldValue) MarshalBinary() ([]byte, error) {
	return d.date.MarshalBinary()
}

func (d *DateFieldValue) UnmarshalBinary(data []byte) error {
	return d.date.UnmarshalBinary(data)
}

func (d *DateFieldValue) String() string {
	return d.date.String()
}
