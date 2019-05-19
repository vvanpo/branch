package fields

import (
	"github.com/vvanpo/titian/field"
)

type Fields struct {
	category
}

// Walk applies fn recursively to each field in the tree in a breadth-first
// manner.
func (f Fields) Walk(fn func(field field.Interface)) {
	var walk = func(fields []field.Interface) {
		for _, field := range fields {
			fn(.fields)
			walk()
		}
	}

	fn(c)
	walk(c.subcategories)
}
