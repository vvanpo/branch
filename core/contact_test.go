package titian

import (
	"testing"
)

// Tests setting an e-mail address.
func TestContactEmail(t *testing.T) {
	app := New(Config{})
	email, _ := NewEmailAddress("victor@example.com")
	contact, err := app.contacts.New(email)

	if err != nil {
		t.Fail()
	}

	email, _ = NewEmailAddress("another@example.com")
	contact.SetPrimaryEmailAddress(email)

	if contact.EmailAddress().String() != "another@example.com" {
		t.Fail()
	}
}
