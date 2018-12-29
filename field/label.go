package field

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// LabelType
type LabelType struct{}

func NewField(name, description string) (*Field, error) {
	return newField(name, description, &LabelType{})
}

func (l LabelType) Validate(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return errors.New("Labels cannot span multiple lines")
	}

	return nil
}
