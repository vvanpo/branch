package titian

import (
	"testing"
)

// Tests that a newly-created contact is in the contact list and has a non-zero
// identifier.
func TestNewContact(t *testing.T) {
	app := New(Config{})
	email, _ := NewEmailAddress("victor@example.com")
	contact, err := app.contacts.New(email)
	var zeroId id

	if err != nil {
		t.Fail()
	}

	if contact.id == zeroId {
		t.Fail()
	}

	if contact.EmailAddress().String() != "victor@example.com" {
		t.Fail()
	}

	if app.contacts.fetch(contact.id) != contact {
		t.Fail()
	}
}
