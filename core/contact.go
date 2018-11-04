package core

import ()

// Contact
type Contact struct {
	id
	// The first e-mail in the slice is the primary e-mail.
	emails []string
	fields map[*Field]string
}

func (c Contact) Email() string {
	return c.emails[0]
}

func (c Contact) Field(field *Field) *string {
	var value string

	if value, ok := c.fields[field]; ok {
		return &value
	}

	return nil
}

func (c *Contact) SetEmail(email string) error {
	c.emails[0] = email
	return nil
}

func (c *Contact) SetField(field *Field, value string) error {
	c.fields[field] = value
	return nil
}
