package titian

import ()

type gid id

type Group struct {
	app         *Container
	id          gid
	name        string
	description string
	contacts    []*Contact
}

// Members returns an unordered list of all contacts belonging to the group.
func (g Group) Members() []*Contact {
	contacts := make([]*Contact, 0, len(g.contacts))

	for _, contact := range g.contacts {
		contacts = append(contacts, contact)
	}

	return contacts
}

func (g *Group) AddContact(contact *Contact) {
	g.contacts[contact.id] = contact
}

func (g *Group) RemoveContact(contact *Contact) {
	delete(g.contacts, contact.id)
}

// Name returns the name of the group.
func (g Group) Name() string {
	return g.name
}

func (g *Group) SetName(name string) {
	g.name = name
}

// Description returns an optional description of the group.
func (g Group) Description() string {
	return g.description
}

func (g *Group) SetDescription(desc string) {
	g.description = desc
}

// RequiredFields returns a list of fields required for members of the group.
func (g Group) RequiredFields() []*Field {
	fields := make([]*Field, 0, len(g.requiredFields))

	for _, field := range g.requiredFields {
		fields = append(fields, field)
	}

	return fields
}

// AddRequiredField
func (g *Group) AddRequiredField(field *Field) {
	g.requiredFields[field.id] = field
}

// RemoveRequiredField
func (g *Group) RemoveRequiredField(field *Field) {
	delete(g.requiredFields, field.id)
}
