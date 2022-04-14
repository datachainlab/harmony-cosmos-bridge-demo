package types

import (
	tmmath "github.com/tendermint/tendermint/libs/math"
)

func NewFractionFromTm(f tmmath.Fraction) *Fraction {
	return &Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}
