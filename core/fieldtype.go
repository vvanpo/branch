package titian

import ()

type FieldType interface {
	Name() string
	NewValue(string) (string, error)
}

// LabelField
type LabelField struct{}

func (_ LabelField) Name() string {
	return "Label"
}

// NewValue validates UTF-8 input.
func (_ LabelField) NewValue(value string) (string, error) {
	// stub
	return value, nil
}

// TextField
type TextField struct {}

func (_ LabelField) Name() string {
	return "Text"
}

// NewValue validates markdown input.
func (_ LabelField) NewValue(value string) (string, error) {
	return value, nil
}

// DateField
type DateField struct {
	id
	datetime bool
	timezone bool
}

func (d DateField) Name() string {
	return "Date"
}

func (d DateField) NewValue(value string) (string, error) {
	// stub
	return value, nil
}
