package core

import ()

//
type User struct {
	app Container
	id
	contact *Contact
}

//
func (u User) Contact() *Contact {
	return u.contact
}

//
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
