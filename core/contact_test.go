package titian

import (
	"testing"
)

func stubContact() (*Contact, error) {
	app := New(Config{})
	email, _ := NewEmailAddress(validEmails[0])
	return app.contacts.New(email)
}

// Tests setting an e-mail address.
func TestContactEmailAddress(t *testing.T) {
	contact, _ := stubContact()
	email2, _ := NewEmailAddress(validEmails[1])

	if contact.EmailAddress().String() != validEmails[0] {
		t.Error("Failure to set initial e-mail address")
	}

	contact.SetPrimaryEmailAddress(email2)

	if contact.EmailAddress().String() != validEmails[1] {
		t.Error("Failure to set primary e-mail address")
	}
}

//
func TestContactAddEmailAddress(t *testing.T) {
	contact, _ := stubContact()

	for _, address := range validEmails[1:] {
		email, _ := NewEmailAddress(address)

		if err := contact.AddEmailAddress(email); err != nil {
			t.Errorf("Failed to add valid e-mail address '%v'", email)
		}
	}

	if contact.AddEmailAddress(contact.email) == nil {
		t.Fatalf("Failed to identify duplicate e-mail address")
	}

	list := contact.EmailAddresses()
	if len(list) != len(validEmails) {
		t.Errorf("E-mail address count doesn't match inputs: %d => %d", len(validEmails), len(list))
	}
}
