package fields

import (
	"github.com/vvanpo/titian/category"
	"github.com/vvanpo/titian/field"
)

type Fields struct {
	root category.Category
}

func (f Fields) Add(f field.Field) {
}

func (f Fields) NewCategory(name, description string) {

}
