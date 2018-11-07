package titian

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
)

type EmailAddress struct {
	local  string
	domain string
}

func NewEmailAddress(address string) (EmailAddress, error) {
	s := strings.Split(address, "@")

	if checkmail.ValidateFormat(address) != nil {
		return EmailAddress{}, errors.New("Invalid e-mail address format")
	}

	local := strings.Join(s[:len(s)-1], "@")
	domain := s[len(s)-1]
	return EmailAddress{local, domain}, nil
}

// String implements the fmt.Stringer interface.
func (e EmailAddress) String() string {
	return e.local + "@" + e.domain
}
