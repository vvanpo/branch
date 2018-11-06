package titian

import (
	"testing"
)

// Tests that a newly-created contact is in the contact list and has a non-zero
// identifier.
func TestNewContact(t *testing.T) {
	const address = "victor@example.com"
	app := New(Config{})
	email, _ := NewEmailAddress(address)
	contact, err := app.contacts.New(email)
	var zeroId id

	if err != nil {
		t.Errorf("New method returned error: %v", err)
	}

	if contact.id == zeroId {
		t.Fatal("Contact does not have initialized ID")
	}

	if contact.EmailAddress().String() != "victor@example.com" {
		t.Errorf("Contact e-mail address does not match expected value: %v => %v", address, contact.EmailAddress())
	}

	if app.contacts.fetch(contact.id) != contact {
		t.Errorf("Cannot fetch contact ID: %v", contact.id)
	}
}
