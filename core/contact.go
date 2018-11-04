package core

import ()

// Contact
type Contact struct {
	id
	email string
}

func (c Contact) Email() string {
	return c.email
}

func (c *Contact) SetEmail(email string) error {
	c.email = email
	return nil
}
