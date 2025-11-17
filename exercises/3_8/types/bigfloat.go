/*Package types Test comment in complex/big*/
package types

import "math/big"

type BigFloatComplex struct {
	Real *big.Float
	Imag *big.Float
	Prec uint
}

func NewBigFloatComplex(real float64, imag float64, precision uint) BigFloatComplex {
	return BigFloatComplex{
		Real: big.NewFloat(real).SetPrec(precision),
		Imag: big.NewFloat(imag).SetPrec(precision),
		Prec: precision,
	}
}

func (z BigFloatComplex) Add(other BigFloatComplex) BigFloatComplex {
	result := BigFloatComplex{
		Real: new(big.Float).SetPrec(z.Prec),
		Imag: new(big.Float).SetPrec(z.Prec),
		Prec: z.Prec,
	}
	result.Real.Add(z.Real, other.Real)
	result.Imag.Add(z.Imag, other.Imag)
	return result
}

func (z BigFloatComplex) Sub(other BigFloatComplex) BigFloatComplex {
	result := BigFloatComplex{
		Real: new(big.Float).SetPrec(z.Prec),
		Imag: new(big.Float).SetPrec(z.Prec),
		Prec: z.Prec,
	}
	result.Real.Sub(z.Real, other.Real)
	result.Imag.Sub(z.Imag, other.Imag)
	return result
}

// Mul ... (a+bi)(c+di) = (ac-bd) + (ad+bc)i
func (z BigFloatComplex) Mul(other BigFloatComplex) BigFloatComplex {
	ac := new(big.Float).SetPrec(z.Prec).Mul(z.Real, other.Real)
	bd := new(big.Float).SetPrec(z.Prec).Mul(z.Imag, other.Imag)
	ad := new(big.Float).SetPrec(z.Prec).Mul(z.Real, other.Imag)
	bc := new(big.Float).SetPrec(z.Prec).Mul(z.Imag, other.Real)

	result := BigFloatComplex{
		Real: new(big.Float).SetPrec(z.Prec),
		Imag: new(big.Float).SetPrec(z.Prec),
		Prec: z.Prec,
	}
	result.Real.Sub(ac, bd)
	result.Imag.Add(ad, bc)
	return result
}

// Div ... (a+bi) / (c+di) = [(ac+bd) + (bc-ad)i] / (c² + d²)
func (z BigFloatComplex) Div(other BigFloatComplex) BigFloatComplex {
	ac := new(big.Float).SetPrec(z.Prec).Mul(z.Real, other.Real)
	bd := new(big.Float).SetPrec(z.Prec).Mul(z.Imag, other.Imag)
	numeratorReal := new(big.Float).SetPrec(z.Prec).Add(ac, bd)

	ad := new(big.Float).SetPrec(z.Prec).Mul(z.Real, other.Imag)
	bc := new(big.Float).SetPrec(z.Prec).Mul(z.Imag, other.Real)
	numeratorImag := new(big.Float).SetPrec(z.Prec).Sub(bc, ad)

	c2 := new(big.Float).SetPrec(z.Prec).Mul(other.Real, other.Real)
	d2 := new(big.Float).SetPrec(z.Prec).Mul(other.Imag, other.Imag)
	denom := new(big.Float).SetPrec(z.Prec).Add(c2, d2)

	return BigFloatComplex{
		Real: new(big.Float).SetPrec(z.Prec).Quo(numeratorReal, denom),
		Imag: new(big.Float).SetPrec(z.Prec).Quo(numeratorImag, denom),
		Prec: z.Prec,
	}
}

// Abs ... : |a+bi| = sqrt(a² + b²)
func (z BigFloatComplex) Abs() *big.Float {
	a2 := new(big.Float).SetPrec(z.Prec).Mul(z.Real, z.Real)
	b2 := new(big.Float).SetPrec(z.Prec).Mul(z.Imag, z.Imag)
	sum := new(big.Float).SetPrec(z.Prec).Add(a2, b2)
	return sum.Sqrt(sum)
}

func (z BigFloatComplex) MulScalar(scalar float64) BigFloatComplex {
	s := big.NewFloat(scalar).SetPrec(z.Prec)
	result := BigFloatComplex{
		Real: new(big.Float).SetPrec(z.Prec).Mul(z.Real, s),
		Imag: new(big.Float).SetPrec(z.Prec).Mul(z.Imag, s),
		Prec: z.Prec,
	}
	return result
}
