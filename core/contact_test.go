package titian

import (
	"testing"
)

// Tests the SetEmail method.
func TestContactSetEmail(t *testing.T) {
	cs := Contacts{}
	contact := cs.New("victor@example.com")
	contact.SetEmail("another@example.com")

	if contact.Email() != "another@example.com" {
		t.Fail()
	}
}
