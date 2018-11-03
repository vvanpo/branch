package core

import ()

type User struct {
	app     Container
	id      ID
	contact *Contact
}

func (u User) ID() ID {
	return u.id
}

func (u User) Contact() *Contact {
	return u.contact
}

func (u User) Access(permission string) bool {
	for _, group := range u.app.Groups {
		if _, ok := group.contacts[u.contact.id]; ok {
			for _, p := range group.permissions {
				if p == permission {
					return true
				}
			}
		}
	}

	return false
}
