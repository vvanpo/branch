package titian

import (
	"testing"
)

// Set and retrieve a primary e-mail address.
func TestContactEmailAddress(t *testing.T) {
	contact, _ := stubContact()
	email := contact.EmailAddress()
	email2, _ := NewEmailAddress(validEmails[1])

	if contact.EmailAddress().String() != validEmails[0] {
		t.Error("Failure to set initial e-mail address")
	}

	contact.SetPrimaryEmailAddress(email2)

	if contact.EmailAddress().String() != validEmails[1] {
		t.Fatal("Failure to set primary e-mail address")
	}

	if err := contact.RemoveEmailAddress(email); err != nil {
		t.Error("Failure to remove alternate e-mail address")
	}
}

// Add multiple alternate e-mail addresses.
func TestContactEmailAddresses(t *testing.T) {
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

	if list := contact.EmailAddresses(); len(list) != len(validEmails) {
		t.Errorf("E-mail address count doesn't match inputs: %d => %d", len(validEmails), len(list))
	}
}

func stubContact() (*Contact, error) {
	app := New(Config{})
	email, _ := NewEmailAddress(validEmails[0])
	return app.contacts.New(email)
}
