package titian

import (
	"errors"
	"strings"
)

type EmailAddress struct {
	local  string
	domain string
}

func NewEmailAddress(address string) (EmailAddress, error) {
	s := strings.Split(address, "@")

	if len(s) < 2 {
		return EmailAddress{}, errors.New("Invalid e-mail address")
	}

	local := strings.Join(s[:len(s)-1], "@")
	domain := s[len(s)-1]
	e := EmailAddress{local, domain}
	return e, nil
}

func (e EmailAddress) String() string {
	return e.local + "+" + e.domain
}
