package field

import (
	"fmt"
)

// Value represents data stored in a field.
type Value interface {
	fmt.Stringer
	Set(string) error
}

/*
// DateFieldValue
type DateFieldValue struct {
	id        fvid
	fieldtype DateField
	date      time.Time
}

func (d *DateFieldValue) MarshalBinary() ([]byte, error) {
	return d.date.MarshalBinary()
}

func (d *DateFieldValue) UnmarshalBinary(data []byte) error {
	return d.date.UnmarshalBinary(data)
}

func (d *DateFieldValue) String() string {
	return d.date.String()
}*/
