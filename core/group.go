package titian

import ()

type Group struct {
	app Container
	id
	name           string
	description    string
	contacts       map[id]*Contact
	permissions    []string
	requiredFields []*Field
}

func (g Group) Name() string {
	return g.name
}

func (g Group) Description() string {
	return g.description
}

func (g Group) Contacts() []*Contact {
	contacts := make([]*Contact, 0, len(g.contacts))

	for _, contact := range g.contacts {
		contacts = append(contacts, contact)
	}

	return contacts
}

func (g Group) Permissions() []string {
	return g.permissions[:]
}

func (g Group) RequiredFields() []*Field {
	return g.requiredFields[:]
}

func (g *Group) SetName(name string) {
	g.name = name
}

func (g *Group) SetDescription(desc string) {
	g.description = desc
}

func (g *Group) AddContact(contact *Contact) {
	g.contacts[contact.id] = contact
}

func (g *Group) RemoveContact(contact *Contact) {
	delete(g.contacts, contact.id)
}

func (g *Group) AddPermission(perm string) {
	for _, p := range g.permissions {
		if p == perm {
			return
		}
	}

	g.permissions = append(g.permissions, perm)
}

func (g *Group) RemovePermission(perm string) {
	for i, p := range g.permissions {
		if p == perm {
			g.permissions = append(g.permissions[:i], g.permissions[i+1:]...)
			return
		}
	}
}
