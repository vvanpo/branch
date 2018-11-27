package titian

import (
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
	if !utf8.ValidString(value) {
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

func (t *TextFieldValue) String() string {
	return t.value
}

func (t *TextFieldValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	t.value = value
	return nil
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
