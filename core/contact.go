package titian

import (
	"errors"
	"fmt"
)

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	app *Container
	id
	// email is a required field, and denotes the contact's primary e-mail
	// address.
	email      EmailAddress
	alternates []EmailAddress
	fields     map[id]FieldValue
}

// EmailAddress returns the contact's primary e-mail address.
func (c Contact) EmailAddress() EmailAddress {
	return c.email
}

// EmailAddresses returns a slice of all e-mail addresses associated with the
// contact.
func (c Contact) EmailAddresses() []EmailAddress {
	return append([]EmailAddress{c.email}, c.alternates...)
}

// AddEmailAddress adds an e-mail address to the contact's e-mail list. Returns
// an error for non-unique e-mails, and will panic if passed the zero value.
// The passed e-mail address is assumed to be verified. A verified e-mail
// address is one that the application has received an e-mail from in the past,
// was provided by a third-party authentication provider, or has been manually
// set as verified. A contact can have multiple associated e-mail addresses if
// they have been linked together (e.g. by the owning user, or by an
// administrator).
func (c *Contact) AddEmailAddress(email EmailAddress) error {
	if email == (EmailAddress{}) {
		panic("Attempting to add an invalid e-mail address")
	}

	if c.app.contacts.Find(email) != nil {
		return fmt.Errorf("E-mail address '%v' is already in use", email)
	}

	c.alternates = append(c.alternates, email)
	return nil
}

// SetPrimaryEmailAddress sets the primary e-mail address for the contact,
// bumping the existing primary e-mail address onto the list of alternate
// addresses. Must be passed a verified e-mail address, which either belongs to
// the contact or is otherwise unique. Will panic if passed a zero-value.
func (c *Contact) SetPrimaryEmailAddress(email EmailAddress) error {
	emails := append([]EmailAddress{c.email}, c.alternates...)

	for i, e := range emails {
		if e == email {
			c.email = email
			c.alternates = append(emails[:i], emails[i+1:]...)
			return nil
		}
	}

	if err := c.AddEmailAddress(email); err != nil {
		return err
	}

	c.email = email
	c.alternates = emails[:len(emails)-1]
	return nil
}

// RemoveEmailAddress removes a secondary e-mail address. Returns an error if
// the passed e-mail is the contact's primary address. Passing an e-mail address
// not belonging to the contact is a no-op.
func (c *Contact) RemoveEmailAddress(email EmailAddress) error {
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

// Merge combines two contacts. The receiver retains its fields and primary
// e-mail address.
func (c *Contact) Merge(associated *Contact) {
	newEmails := associated.EmailAddresses()
	c.app.contacts.Delete(associated)
	c.alternates = append(c.alternates, newEmails...)
}

// Field
func (c Contact) Field(field *Field) FieldValue {
	if value, ok := c.fields[field.id]; ok {
		return value
	}

	return nil
}

func (c *Contact) SetField(value FieldValue) {
	c.fields[value.Field().id] = value
}

//func (c *Contact) DeleteField(
