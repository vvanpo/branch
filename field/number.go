package field

import (
	"math/big"
)

func NewNumberField(name, description string, num NumberType) (*Field, error) {
	return newField(name, description, num)
}

type NumberType struct {
	minimum *big.Rat
	maximum *big.Rat
	step    *big.Rat
	// Whether to display the number as a fraction or decimal.
	isFraction bool
}

func (n NumberType) NewValue(x, y int64) (Number, error) {
	number := Number{*math.NewRat(x, y)}
}

type Number struct {
	value big.Rat
}
