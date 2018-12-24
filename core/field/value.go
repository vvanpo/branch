package titian

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Value represents a value for a particular field type.
type Value interface {
	fmt.Stringer
	Set(string) error
}

// TextValue
type TextValue struct {
	id    uuid.UUID
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
