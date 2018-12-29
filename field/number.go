package field

import (
	"math/big"
)

func NewNumber(name, description string) Type {
	field := newField
}

type numberType struct {
	minimum *big.Rat
	maximum *big.Rat
	step    *big.Rat
	// Whether to display the number as a fraction or decimal.
	isFraction bool
}

func (n *numberType) NewValue(value *big.Rat) (Value, error) {

}

type number big.Rat
