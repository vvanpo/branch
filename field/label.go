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
	label := new(labelValue)

	if err := label.Set(value); err != nil {
		return nil, err
	}

	return label, nil
}

// labelValue
type labelValue string

func (l *labelValue) String() string {
	return string(*l)
}

func (l *labelValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return errors.New("Labels cannot span multiple lines")
	}

	*l = labelValue(value)
	return nil
}
