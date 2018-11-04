package titian

import ()

type FieldCategory struct {
	app Container
	id
	name          string
	description   string
	subcategories []*FieldCategory
}
