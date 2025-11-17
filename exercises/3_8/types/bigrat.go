package types

import (
	"math/big"
)

type BigRatComplex struct {
	Real *big.Rat
	Imag *big.Rat
}

func NewBigRatComplex(real float64, imag float64) BigRatComplex {
	result := BigRatComplex{
		Real: new(big.Rat).SetFloat64(real),
		Imag: new(big.Rat).SetFloat64(imag),
	}
	return result
}

/* Add ... (a + bi) + (c + di) = (a+c) + (b+d)i*/
func (z BigRatComplex) Add(other BigRatComplex) BigRatComplex {
	result := BigRatComplex{
		Real: new(big.Rat).Add(z.Real, other.Real),
		Imag: new(big.Rat).Add(z.Imag, other.Imag),
	}
	return result
}

/* Sub ... (a + bi) - (c + di) = (a-c) + (b-d)i*/
func (z BigRatComplex) Sub(other BigRatComplex) BigRatComplex {
	result := BigRatComplex{
		Real: new(big.Rat).Sub(z.Real, other.Real),
		Imag: new(big.Rat).Sub(z.Imag, other.Imag),
	}
	return result
}

/* MulScalar ... (a + bi)(s) = as + bsi*/
func (z BigRatComplex) MulScalar(s *big.Rat) BigRatComplex {
	result := BigRatComplex{
		Real: new(big.Rat).Mul(z.Real, s),
		Imag: new(big.Rat).Mul(z.Imag, s),
	}
	return result
}

// Mul ... (a+bi)(c+di) = (ac-bd) + (ad+bc)i
func (z BigRatComplex) Mul(other BigRatComplex) BigRatComplex {
	ac := new(big.Rat).Mul(z.Real, other.Real)
	bd := new(big.Rat).Mul(z.Imag, other.Imag)
	ad := new(big.Rat).Mul(z.Real, other.Imag)
	bc := new(big.Rat).Mul(z.Imag, other.Real)

	result := BigRatComplex{
		Real: new(big.Rat).Sub(ac, bd),
		Imag: new(big.Rat).Add(ad, bc),
	}
	return result
}

// Div ... (a+bi) / (c+di) = [(ac+bd) + (bc-ad)i] / (c² + d²)
func (z BigRatComplex) Div(other BigRatComplex) BigRatComplex {
	ac := new(big.Rat).Mul(z.Real, other.Real)
	bd := new(big.Rat).Mul(z.Imag, other.Imag)
	real := new(big.Rat).Add(ac, bd)

	ad := new(big.Rat).Mul(z.Real, other.Imag)
	bc := new(big.Rat).Mul(z.Imag, other.Real)
	imag := new(big.Rat).Sub(bc, ad)

	c2 := new(big.Rat).Mul(other.Real, other.Real)
	d2 := new(big.Rat).Mul(other.Imag, other.Imag)
	denom := new(big.Rat).Add(c2, d2)

	result := BigRatComplex{
		Real: new(big.Rat).Quo(real, denom),
		Imag: new(big.Rat).Quo(imag, denom),
	}
	return result
}

// AbsSquared ... : |a+bi|² = a² + b²
func (z BigRatComplex) AbsSquared() *big.Rat {
	a2 := new(big.Rat).Mul(z.Real, z.Real)
	b2 := new(big.Rat).Mul(z.Imag, z.Imag)
	sum := new(big.Rat).Add(a2, b2)
	return sum
}
