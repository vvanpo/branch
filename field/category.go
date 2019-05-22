package field

type Category struct {
	Name          string
	Description   string
	Fields        []Field
	Subcategories []Category
}
