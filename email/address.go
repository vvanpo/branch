package email

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Address struct {
	local  string
	domain string
}

func NewAddress(address string) (Address, error) {
	a := Address{}
	s := strings.Split(address, "@")

	if checkmail.ValidateFormat(address) != nil {
		return a, errors.New("Invalid e-mail address format")
	}

	a.local = strings.Join(s[:len(s)-1], "@")
	a.domain = s[len(s)-1]
	return a, nil
}

// String implements the fmt.Stringer interface.
func (a Address) String() string {
	if a == (Address{}) {
		return ""
	}

	return a.local + "@" + a.domain
}
