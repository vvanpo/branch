package titian

import (
	"testing"
)

// Tests that a newly-created contact is in the contact list and has a non-zero
// identifier.
func TestNewContact(t *testing.T) {
	cs := Contacts{}
	contact := cs.New("victor@example")
	var zeroId id

	if contact.id == zeroId {
		t.Fail()
	}

	if contact.Email() != "victor@example" {
		t.Fail()
	}

	if cs.fetch(contact.id) != contact {
		t.Fail()
	}
}
