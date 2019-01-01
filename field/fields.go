package field

type Categories interface {
	Fetch(parent *Category) []*Category
	Fields() Fields
	Persist(*Category)
}

type Fields interface {
	Fetch(parent *Category) []*Field
	Persist(*Field)
}
