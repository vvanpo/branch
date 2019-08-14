package contact

import (
	"github.com/vvanpo/stonefruit/email"
	"github.com/vvanpo/stonefruit/field"
)

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	fields map[*field.Field]field.Value
}

// Fields returns a map of all field values.
func (c Contact) Fields() map[*field.Field]field.Value {
	fields := make(map[*field.Field]field.Value, len(c.fields))

	for field, value := range c.fields {
		fields[field] = value
	}

	return fields
}

// SetField
func (c *Contact) SetField(field *field.Field, value string) error {
	if _, ok := c.fields[field]; ok {
		return c.fields[field].Set(value)
	}

	newValue, err := field.Type().NewValue(value)

	if err != nil {
		return err
	}

	c.fields[field] = newValue
	return nil
}

// DeleteField deletes the value associated with the passed field.
func (c *Contact) DeleteField(field *field.Field) {
	delete(c.fields, field)
}

// Merge combines two contacts. The receiver retains its fields while additional
// fields from the associated contact are added.
func (c *Contact) Merge(associated *Contact) {
	for address := range associated.EmailAddresses() {
		c.AddEmailAddress
	}

	for field, value := range associated.fields {
		if _, ok := c.fields[field]; !ok {
			c.fields[field] = value
		}
	}
}
