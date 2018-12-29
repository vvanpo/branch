package field

import (
	"errors"
	"math/big"
)

type NumberType struct {
	Minimum *big.Rat
	Maximum *big.Rat
	Step    *big.Rat
}

func (n NumberType) NewField(name, description string) (*Field, error) {
	return newField(name, description, n)
}

func (n NumberType) NewValue(x, y int64) error {
	if y == 0 {
		return errors.New("Denominator cannot be zero")
	}

	return nil
}

type Number struct {
	value big.Rat
}
