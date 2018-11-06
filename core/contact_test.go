package titian

import (
	"testing"
)

// Tests the SetEmail method.
func TestContactSetEmail(t *testing.T) {
	cs := Contacts{}
	email, _ := NewEmail("victor@example.com")
	contact, err := cs.New(email)

	if err != nil {
		t.Fail()
	}

	contact.SetEmail("another@example.com")

	if contact.Email() != "another@example.com" {
		t.Fail()
	}
}
