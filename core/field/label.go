package titian

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Label
type Label struct{}

func (_ Label) Name() string {
	return "Label"
}

func (_ Label) NewValue(value string) (Value, error) {
	label := &LabelValue{}

	if err := label.Set(value); err != nil {
		return nil, err
	}

	return label, nil
}

// LabelValue
type LabelValue struct {
	value string
}

func (l *LabelValue) String() string {
	return l.value
}

func (l *LabelValue) Set(value string) error {
	if !utf8.ValidString(value) {
		return errors.New("Invalid string encoding")
	}

	if strings.Contains(value, "\n") {
		return errors.New("Labels cannot span multiple lines")
	}

	l.value = value
	return nil
}
