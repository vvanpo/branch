package branch

import ()

type contacts struct {
	db
}

func (cs *contacts) store(c *Contact) error {
	return nil
}
