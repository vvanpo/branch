package field

import (
	"github.com/vvanpo/titian/email"
)

type EmailAddressesType struct {
	Maximum uint
}

func (e EmailAddressesType) NewValue() Value {
	return make([]email.Address, 0)
}

func (e EmailAddressesType) Validate() error {
	return nil
}

/*

func (e *EmailAddressesType) NewValue(addresses ...email.Address) EmailAddresses {
	if len(addresses) == 0 {
		panic("Cannot create e-mail address list without addresses")
	}

	list := EmailAddresses{make(map[email.Address]struct{})}
	list.Add(addresses...)
	return list
}

type EmailAddresses struct {
	list map[email.Address]struct{}
}

// List returns a sorted slice of the e-mail addresses.
func (e EmailAddresses) List() []email.Address {
	addresses := make([]email.Address, len(e.list))

	for address := range e.list {
		addresses = append(addresses, address)
	}

	return addresses
}

// Add appends an e-mail address or addresses to the list. Will panic if any
// passed address is the zero value.
func (e *EmailAddresses) Add(addresses ...email.Address) {
	for _, address := range addresses {
		if address == (email.Address{}) {
			panic("Attempting to add an invalid e-mail address")
		}

		e.list[address] = struct{}{}
	}
}

// Remove deletes an e-mail address from the list. Attempting to remove the last
// address will cause the method to panic. Passing an address not present in the
// list is a no-op.
func (e *EmailAddresses) Remove(address email.Address) {
	if len(e.list) <= 1 {
		panic("Cannot remove all e-mail addresses from list")
	}

	delete(e.list, address)
} */
