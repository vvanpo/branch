package titian

import (
	"testing"
)

func TestEmailAddressString(t *testing.T) {
	email, _ := NewEmailAddress("victor@example.com")

	if email != "victor@example.com" {
		t.Fail()
	}
}

var invalidEmails = [...]string{
	"",
	"name",
}

func TestValidateEmailAddress(t *testing.T) {
	for _, email := range invalidEmails {
		email, err := NewEmailAddress("victor")

		if email != (EmailAddress{}) || err == nil {
			t.Fail()
		}
	}
}
