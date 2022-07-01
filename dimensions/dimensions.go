package dimensions

import (
	"errors"
	"log"
	"math"
)

type Imp struct {
	Inches      int
	Numerator   int
	Denominator int
}

type Fraction struct {
	Numerator   int
	Denominator int
}

type Rectangle struct {
	X Imp
	Y Imp
}

func Inches(length int) Imp {
	return Imp{
		Inches:      length,
		Numerator:   0,
		Denominator: 32,
	}
}

func BuildImp(inches int, numerator int, denominator int) Imp {
	if denominator == 0 {
		log.Panicln("Denominator cannot be zero")
	}
	return Imp{
		Inches:      inches,
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func (i *Imp) ToFloat() float64 {
	fractional := float64(i.Numerator) / float64(i.Denominator)
	return float64(i.Inches) + fractional
}

func ToFractional(total float64) (Imp, error) {
	inches, fractional := math.Modf(total)

	var denominators = []int{
		2, 4, 8, 32,
	}

	for _, denominator := range denominators {
		for numerator := 1; numerator < denominator; numerator++ {
			if float64(numerator)/float64(denominator) == float64(fractional) {
				return Imp{
					Inches:      int(inches),
					Numerator:   numerator,
					Denominator: denominator,
				}, nil
			}
		}
	}
	return Imp{}, errors.New("Cannot convert to fractional")
}
