package core

import (
	"github.com/vvanpo/branch/core/repo"
)

// Contact
type Contact struct {
	repo  *repo.Contacts
	id    repo.ID
	email string
}

func (c Contact) Email() string {
	return c.email
}

func (c *Contact) SetEmail(email string) error {
	c.email = email
	c.repo.Save(c.getEntity())
	return nil
}

func (c *Contact) getEntity() repo.Contact {
	return repo.Contact{c.id, c.email}
}
