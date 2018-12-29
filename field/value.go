package field

import (
	"fmt"
)

// Value represents data stored in a particular field.
type Value interface {
	fmt.Stringer
}
