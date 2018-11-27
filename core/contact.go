package titian

import (
	"errors"
	"fmt"
)

type cid id

// A Contact represents any individual or organization with a known e-mail
// address or addresses.
type Contact struct {
	app *Container
	id  cid
	// email is a required field, and denotes the contact's primary e-mail
	// address.
	email      EmailAddress
	alternates []EmailAddress
	fields     map[fid]FieldValue
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
// addresses. If the passed e-mail address does not already belong to the
// contact, it is assumed to be verified. Will panic if passed a zero-value.
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

// Field returns the value of the specified field, and nil if it is not present.
func (c Contact) Field(field *Field) FieldValue {
	if value, ok := c.fields[field.id]; ok {
		return value
	}

	return nil
}

func (c Contact) Fields() []*Field {
	fields := make([]*Field, len(c.fields))

	for id := range c.fields {
		fields = append(fields, c.app.fields.field(id))
	}

	return fields
}

func (c *Contact) SetField(field *Field, value string) error {
	if _, ok := c.fields[field.id]; ok {
		return c.fields[field.id].Set(value)
	}

	newValue, err := field.Type().New(value)

	if err != nil {
		return err
	}

	c.fields[field.id] = newValue
	return nil
}

func (c *Contact) DeleteField(field *Field) {
	delete(c.fields, field.id)
}

// Groups
func (c Contact) Groups() []*Group {
	groups := make([]*Group, 0)

	for _, group := range c.app.groups.list {
		for _, member := range group.Members() {
			if c.id == member.id {
				groups = append(groups, group)
				break
			}
		}
	}

	return groups
}

// Merge combines two contacts. The receiver retains its fields and primary
// e-mail address, and additional fields and groups from the associated contact
// are added.
func (c *Contact) Merge(associated *Contact) {
	emails := associated.EmailAddresses()
	fields := associated.fields
	groups := associated.Groups()
	c.app.contacts.Delete(associated)
	c.alternates = append(c.alternates, emails...)

	for id, value := range fields {
		if _, ok := c.fields[id]; !ok {
			c.fields[id] = value
		}
	}

	for _, group := range groups {
		group.AddContact(c)
	}
}
