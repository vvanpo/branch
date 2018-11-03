package core

import (
	"testing"
)

func TestNewContact(t *testing.T) {
	cs := Contacts{}
	contact := cs.New("victor@example")
	var id ID
	if contact.ID() == id {
		t.Fail()
	}
}
