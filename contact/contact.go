package contact

import (
	"sort"

	"github.com/vvanpo/titian/email"
	"github.com/vvanpo/titian/field"
)

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	emails map[email.Address]struct{}
	fields map[*field.Field]field.Value
}

// EmailAddresses returns a slice of all e-mail addresses associated with the
// contact.
func (c Contact) EmailAddresses() []email.Address {
	addresses := make([]string, len(c.emails))

	for address := range c.emails {
		addresses = append(addresses, address)
	}

	sort.Strings(addresses)
	return addresses
}

// AddEmailAddress adds an e-mail address or addresses to the contact's e-mail
// list. Will panic if passed the zero value.
func (c *Contact) AddEmailAddress(addresses ...email.Address) {
	for _, address := range addresses {
		if address == (email.Address{}) {
			panic("Attempting to add an invalid e-mail address")
		}

		c.emails[address] = struct{}{}
	}
}

// RemoveEmailAddress removes an e-mail address from the contact. Panics if
// removing the address would leave the contact without any e-mail addresses.
// Passing an e-mail address not belonging to the contact is a no-op.
func (c *Contact) RemoveEmailAddress(address email.Address) {
	if len(c.emails) <= 1 {
		panic("Cannot remove all e-mail addresses from contact")
	}

	delete(c.emails, address)
}

// Fields returns a map of all field values.
func (c Contact) Fields() map[*field.Field]string {
	fields := make(map[*field.Field]string, len(c.fields))

	for field, value := range c.fields {
		fields[field] = value.String()
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
