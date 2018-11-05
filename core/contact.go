package titian

import ()

// Contact
type Contact struct {
	id
	// The first e-mail in the slice is the primary e-mail.
	emails []string
	fields map[id]FieldValue
}

// Email
func (c Contact) Email() string {
	return c.emails[0]
}

func (c *Contact) SetEmail(email string) error {
	c.emails[0] = email
	return nil
}

// Field
func (c Contact) Field(field Field) FieldValue {
	if value, ok := c.fields[field.id]; ok {
		return value
	}

	return nil
}

func (c *Contact) SetField(value FieldValue) {
	c.fields[value.Field().id] = value
}
