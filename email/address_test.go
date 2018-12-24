package email

import (
	"testing"
)

var validAddresses = [...]string{
	"victor@gmail.com",
	"victor_@gmail.com",
	"!def!xyz%abc@example.com",
	"a@b",
}

var invalidAddresses = [...]string{
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

func TestAddressString(t *testing.T) {
	address, _ := NewAddress(validAddresses[0])
	address2, _ := NewAddress(invalidAddresses[0])

	if address.String() != validAddresses[0] {
		t.Errorf("E-mail input does not match string output: %v => %v", validAddresses[0], address.String())
	}

	if address2.String() != "" {
		t.Errorf("Invalid e-mail address does not output empty string: %v => %v", invalidAddresses[0], address2.String())
	}
}

func TestInvalidAddress(t *testing.T) {
	for _, invalid := range invalidAddresses {
		address, err := NewAddress(invalid)

		if address != (Address{}) || err == nil {
			t.Errorf("Invalid e-mail address '%v' passed validation", invalid)
		}
	}
}

func TestValidAddress(t *testing.T) {
	for _, valid := range validAddresses {
		address, err := NewAddress(valid)

		if address == (Address{}) || err != nil {
			t.Errorf("Valid e-mail address '%v' failed validation", valid)
		}
	}
}
