package contact

import (
	"errors"

	"github.com/vvanpo/titian/email"
	"github.com/vvanpo/titian/field"
)

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	emails []email.Address
	fields map[*field.Field]field.Value
}

// EmailAddresses returns a slice of all e-mail addresses associated with the
// contact.
func (c Contact) EmailAddresses() []email.Address {
	return c.email[:]
}

// AddEmailAddress adds an e-mail address to the contact's e-mail list. Will
// panic if passed the zero value.
func (c *Contact) AddEmailAddress(address email.Address) error {
	if address == (email.Address{}) {
		panic("Attempting to add an invalid e-mail address")
	}

	c.email = append(c.email, address)
	return nil
}

// RemoveEmailAddress removes an e-mail address from the contact. Panics if
// removing the address would leave the contact without any e-mail addresses.
// Passing an e-mail address not belonging to the contact is a no-op.
func (c *Contact) RemoveEmailAddress(address email.Address) error {
	if len(c.emails) <= 1 {
		panic("Cannot remove all e-mail addresses from contact")
	}

	for i, a := range c.emails {
		if a == address {
			c.emails = append(c.emails[:i], c.emails[i+1:]...)
			return nil
		}
	}

	return nil
}

// Fields
func (c Contact) Fields() []*field.Field {
	fields := make([]*field.Field, len(c.fields))

	for field := range c.fields {
		fields = append(fields, field)
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

// DeleteField
func (c *Contact) DeleteField(field *field.Field) {
	delete(c.fields, field)
}

// Merge combines two contacts. The receiver retains its fields and primary
// e-mail address, and additional fields from the associated contact are added.
func (c *Contact) Merge(associated *Contact) {
	emails := associated.EmailAddresses()
	fields := associated.fields
	c.emails = append(c.emails, emails...)

	for field, value := range fields {
		if _, ok := c.fields[field]; !ok {
			c.fields[field] = value
		}
	}
}
