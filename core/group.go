package core

import ()

type Group struct {
	id ID
	contacts map[ID]*Contact
	users map[ID]*User
	roles map[ID]*Role
}
