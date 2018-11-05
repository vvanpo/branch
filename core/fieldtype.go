package titian

import (
	"fmt"
	"strconv"
)

type FieldType interface {
	Name() string
	NewValue() FieldValue
}

// LabelField
type LabelField struct{}

func (_ LabelField) Name() string {
	return "Label"
}

func (_ LabelField) NewValue() FieldValue {
	return &LabelFieldValue{}
}

// TextField
type TextField struct{}

func (_ TextField) Name() string {
	return "Text"
}

func (_ TextField) NewValue() FieldValue {
	return &TextFieldValue{}
}

// NumberField
type NumberField struct {
	id
	unsigned bool
	integer  bool
}

func (n NumberField) Name() string {
	return "Number"
}

func (n NumberField) NewValue(value string) (string, error) {
	number, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(number), nil
}

// FlagField
type FlagField struct{}

func (_ FlagField) Name() string {
	return "Flag"
}

func (_ FlagField) NewValue(value string) (string, error) {
	b, err := strconv.ParseBool(value)

	if err == nil {
		if b {
			return "True", nil
		}

		return "False", nil
	}

	return "", err
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

// ChecklistField
type ChecklistField struct {
	id
	items []string
}

func (c ChecklistField) Name() string {
	return "Checklist"
}

func (c ChecklistField) NewValue(value string) (string, error) {
	//stub
	return value, nil
}

// SelectionField
type SelectionField struct {
	id
	items []string
}

func (s SelectionField) Name() string {
	return "Selection"
}

func (s SelectionField) NewValue(value string) (string, error) {
	//stub
	return value, nil
}
