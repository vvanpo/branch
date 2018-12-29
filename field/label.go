package field

import (
	"errors"
	"strings"
	"unicode/utf8"
)

func NewLabel(name, description string) (*Field, error) {
	return newField(name, description, &labelType{})
}

// labelType
type labelType struct{}

func (_ labelType) NewValue(value string) (Value, error) {
	if !utf8.ValidString(value) {
		return nil, errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return nil, errors.New("Labels cannot span multiple lines")
	}

	return labelValue(value), nil
}

// labelValue
type labelValue string

func (l labelValue) String() string {
	return string(l)
}
