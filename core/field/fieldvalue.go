package titian

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

// FieldValue represents a value for a particular field type.
type FieldValue interface {
	fmt.Stringer
	Set(string) error
}

// LabelFieldValue
type LabelFieldValue struct {
	id    uuid.UUID
	value string
}

func (l *LabelFieldValue) String() string {
	return l.value
}

func (l *LabelFieldValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return errors.New("Labels cannot span multiple lines")
	}

	l.value = value
	return nil
}

// TextFieldValue
type TextFieldValue struct {
	id    uuid.UUID
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

/*
// DateFieldValue
type DateFieldValue struct {
	id        fvid
	fieldtype DateField
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
}*/
