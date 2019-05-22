package field

import (
//	"errors"
//	"strings"
//	"unicode/utf8"
)

type LabelField struct {
}

/**
type label string

func (l LabelType) NewValue() Value {
	return new(label)
}

func (l LabelType) Validate() error {
	return nil
}

func (l *label) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return errors.New("Labels cannot span multiple lines")
	}

	*l = label(value)
	return nil
}
*/
