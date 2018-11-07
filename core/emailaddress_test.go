package titian

import (
	"testing"
)

var validEmails = [...]string{
	"victor@gmail.com",
	"victor_@gmail.com",
	"!def!xyz%abc@example.com",
	"a@b",
}

var invalidEmails = [...]string{
	"",
	"@b",
	"a@",
	"name",
	"ç$€§/az@gmail.com",
	"@gmail.com",
	"victor@_gmail.com",
	"victor@",
	"victor@yahooo.",
}

func TestEmailAddressString(t *testing.T) {
	email, _ := NewEmailAddress(validEmails[0])
	email2, _ := NewEmailAddress(invalidEmails[0])

	if email.String() != validEmails[0] {
		t.Errorf("E-mail input does not match string output: %v => %v", validEmails[0], email.String())
	}

	if email2.String() != "" {
		t.Errorf("Invalid e-mail address does not output empty string: %v => %v", invalidEmails[0], email2.String())
	}
}

func TestInvalidEmailAddress(t *testing.T) {
	for _, invalid := range invalidEmails {
		email, err := NewEmailAddress(invalid)

		if email != (EmailAddress{}) || err == nil {
			t.Errorf("Invalid e-mail address '%v' passed validation", invalid)
		}
	}
}

func TestValidEmailAddress(t *testing.T) {
	for _, valid := range validEmails {
		email, err := NewEmailAddress(valid)

		if email == (EmailAddress{}) || err != nil {
			t.Errorf("Valid e-mail address '%v' failed validation", valid)
		}
	}
}
