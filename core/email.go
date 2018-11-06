package titian

import ()

type Email struct {
	local string
	domain string
}

func NewEmail(email string) (Email, error) {
	e := Email{}
	return e
}

func (e Email) String() string {
	return e.local + "+" + e.domain
}

