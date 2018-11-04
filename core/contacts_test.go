package core

import (
	"testing"
)

func TestNewContact(t *testing.T) {
	cs := Contacts{}
	contact := cs.New("victor@example")
	var id id

	if contact.id == id {
		t.Fail()
	}
}
