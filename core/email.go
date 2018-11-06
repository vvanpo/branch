package titian

import (
	"errors"
	"strings"
)

type Email struct {
	local  string
	domain string
}

func NewEmail(email string) (Email, error) {
	s := strings.Split(email, "@")

	if len(s) < 2 {
		return errors.New("Invalid e-mail address")
	}

	local := strings.Join(s[:len(s)-1], "@")
	domain := s[len(s)-1]
	e := Email{local, domain}
	return e, nil
}

func (e Email) String() string {
	return e.local + "+" + e.domain
}
