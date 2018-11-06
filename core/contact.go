package titian

import (
	"errors"
)

// Contact
type Contact struct {
	app *Container
	id
	email Email
	verifiedEmails []string
	fields map[id]FieldValue
}

// Email
func (c Contact) Email() Email {
	return c.email
}

// VerifiedEmails
func (c Contact) VerifiedEmails() [:]string {
	return append([]string{c.email}, c.verifiedEmails...)
}

// SetPrimaryEmail swaps the current primary e-mail for a verified e-mail.
// Returns an error if the passed e-mail is not in the user's verified e-mails
// list.
func (c *Contact) SetPrimaryEmail(email Email) error {
	emails := verifiedEmails()

	for i, e := range emails {
		if e == email {
			c.email = email
			c.verifiedEmails = append(emails[:i], emails[i+1:]...)
			return nil
		}
	}

	return errors.New("'" + email + "' is not a verified e-mail address")
}

// AddVerifiedEmail adds an e-mail address to the verified e-mails list. Returns
// an error if passed the zero value or adding an address already in use.
func (c *Contact) AddVerifiedEmail(email Email) error {
	if email == Email{} {
		return errors.New("E-mail address is invalid")
	}

	if app.contacts.Find(email) != nil {
		return errors.New("E-mail address '" + email + "' is already in use")
	}

	c.verifiedEmails = append(c.verifiedEmails, email)
	return nil
}

// RemoveEmail removes a verified secondary e-mail address. Returns an error if
// the passed e-mail is the contact's primary address.
func (c *Contact) RemoveEmail(email Email) error {
	if email == c.email {
		return errors.New("Cannot remove your primary e-mail address")
	}

	for i, e := range c.verifiedEmails {
		if email == e {
			c.verifiedEmails = append(c.verifiedEmails[:i], c.verifiedEmails[i+1:]...)
			return nil
		}
	}

	return nil
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

