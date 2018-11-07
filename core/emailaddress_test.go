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
	"ç$€§/az@gmail.com",
	"@gmail.com",
	"victor@_gmail.com",
	"victor@",
}

var validEmails = [...]string{
	"a@b",
	"victor@gmail.com",
}

func TestValidateEmailAddress(t *testing.T) {
	for _, invalid := range invalidEmails {
		email, err := NewEmailAddress(invalid)

		if email != (EmailAddress{}) || err == nil {
			t.Errorf("Invalid e-mail address '%v' passed validation", invalid)
		}
	}

	for _, valid := range validEmails {
		email, err := NewEmailAddress(valid)

		if email == (EmailAddress{}) || err != nil {
			t.Errorf("Valid e-mail address '%v' failed validation", valid)
		}
	}
}
