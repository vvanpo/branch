package core

import ()

// Contact
type Contact struct {
	id
	email  string
	fields map[*Field]string
}

func (c Contact) Email() string {
	return c.email
}

func (c *Contact) SetEmail(email string) error {
	c.email = email
	return nil
}

func (c *Contact) SetField(field *Field, value string) error {
	c.fields[field] = value
}
