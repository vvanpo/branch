package titian

import (
	"testing"
)

func TestEmailAddressString(t *testing.T) {
	email, _ := NewEmailAddress("victor@example.com")

	if email.String() != "victor@example.com" {
		t.Fail()
	}
}

var invalidEmails = [...]string{
	"",
	"name",
}

func TestValidateEmailAddress(t *testing.T) {
	for _, invalid := range invalidEmails {
		email, err := NewEmailAddress(invalid)

		if email != (EmailAddress{}) || err == nil {
			t.Fail()
		}
	}
}
