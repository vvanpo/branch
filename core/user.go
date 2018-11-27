package titian

import ()

type uid id

//
type User struct {
	app     Container
	id      uid
	contact *Contact
}

//
func (u User) Contact() *Contact {
	return u.contact
}

//
func (u User) Access(permission string) bool {
	for _, group := range u.app.groups.list {
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

func (u User) AccessField(field *Field) (read, write bool) {
	return
}

func (u User) AccessGroupField(field *Field, group *Group) (read, write bool) {
	return
}
