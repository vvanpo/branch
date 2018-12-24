package contact

import (
	"errors"

	"github.com/vvanpo/titian/email"
	"github.com/vvanpo/titian/field"
)

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	// email is a required field, and denotes the contact's primary e-mail
	// address.
	email      email.Address
	alternates []email.Address
	fields     map[*field.Field]field.Value
}

// EmailAddress returns the contact's primary e-mail address.
func (c Contact) EmailAddress() email.Address {
	return c.email
}

// EmailAddresses returns a slice of all e-mail addresses associated with the
// contact.
func (c Contact) EmailAddresses() []email.Address {
	return append([]email.Address{c.email}, c.alternates...)
}

// AddEmailAddress adds an e-mail address to the contact's e-mail list. Returns
// an error for non-unique e-mails, and will panic if passed the zero value.
// The passed e-mail address is assumed to be verified. A verified e-mail
// address is one that the application has received an e-mail from in the past,
// was provided by a third-party authentication provider, or has been manually
// set as verified. A contact can have multiple associated e-mail addresses if
// they have been linked together (e.g. by the owning user, or by an
// administrator).
func (c *Contact) AddEmailAddress(address email.Address) error {
	if address == (email.Address{}) {
		panic("Attempting to add an invalid e-mail address")
	}

	c.alternates = append(c.alternates, address)
	return nil
}

// SetPrimaryEmailAddress sets the primary e-mail address for the contact,
// bumping the existing primary e-mail address onto the list of alternate
// addresses. If the passed e-mail address does not already belong to the
// contact, it is assumed to be verified. Will panic if passed a zero-value.
func (c *Contact) SetPrimaryEmailAddress(address email.Address) error {
	emails := append([]email.Address{c.email}, c.alternates...)

	for i, e := range emails {
		if e == address {
			c.email = address
			c.alternates = append(emails[:i], emails[i+1:]...)
			return nil
		}
	}

	if err := c.AddEmailAddress(address); err != nil {
		return err
	}

	c.email = address
	c.alternates = emails[:len(emails)-1]
	return nil
}

// RemoveEmailAddress removes a secondary e-mail address. Returns an error if
// the passed e-mail is the contact's primary address. Passing an e-mail address
// not belonging to the contact is a no-op.
func (c *Contact) RemoveEmailAddress(email email.Address) error {
	if email == c.email {
		return errors.New("Cannot remove your primary e-mail address")
	}

	for i, e := range c.alternates {
		if email == e {
			c.alternates = append(c.alternates[:i], c.alternates[i+1:]...)
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

func (c *Contact) DeleteField(field *field.Field) {
	delete(c.fields, field)
}

// Merge combines two contacts. The receiver retains its fields and primary
// e-mail address, and additional fields from the associated contact are added.
func (c *Contact) Merge(associated *Contact) {
	emails := associated.EmailAddresses()
	fields := associated.fields
	c.alternates = append(c.alternates, emails...)

	for field, value := range fields {
		if _, ok := c.fields[field]; !ok {
			c.fields[field] = value
		}
	}
}
