package core

import ()

type User struct {
	id      ID
	contact ID
	roles   []ID
	app Container
}

func (u User) ID() ID {
	return u.id
}

func (u User) Contact() ID {
	return u.contact
}

func (u User) Access(permission string) bool {
	roles := u.roles

	// Concatenate all role lists from groups
	for _, group := range u.app.Groups {
		roles = append(roles, group.Roles()...)
	}

	// Use a map to ensure we only check each role once
	hashmap := make(map[ID]struct{})
	for _, r := range roles {
		if _, ok := hashmap[r]; !ok {
			for _, p := range u.app.Roles.Fetch(r).Permissions() {
				if p == permission {
					return true
				}
			}
		}

		hashmap[r] = struct{}{}
	}

	return false
}
