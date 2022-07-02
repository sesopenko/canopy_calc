// Package dimensions provides dimensional representation and mathematics for Imperial distance measurements
package dimensions

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Imp struct {
	Inches      int
	Numerator   int
	Denominator int
}

func (i Imp) Format() string {
	return fmt.Sprintf("%d %d/%d\"", i.Inches, i.Numerator, i.Denominator)
}

type Fraction struct {
	Numerator   int
	Denominator int
}

type Rectangle struct {
	Width  Imp
	Height Imp
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

func (i Imp) ToFloat() float64 {
	fractional := float64(i.Numerator) / float64(i.Denominator)
	return float64(i.Inches) + fractional
}

func ToFractional(total float64) (Imp, error) {
	inches, fractional := math.Modf(total)

	var denominators = []int{
		2, 4, 8, 32,
	}
	if fractional == 0.0 {
		return Imp{
			Inches:      int(inches),
			Numerator:   0,
			Denominator: 2,
		}, nil
	}

	// not efficient but it works :P
	// Could use something like a lookup table instead, once the program's functional
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
	return Imp{}, errors.New(fmt.Sprintf("Cannot find denominator for: %f", total))
}

func (i Imp) Add(k Imp) Imp {
	added, err := ToFractional(i.ToFloat() + k.ToFloat())
	if err != nil {
		log.Panicln("Could not convert added values to fractional")
	}
	return added
}

func (i Imp) Subtract(k Imp) Imp {
	total := i.ToFloat() - k.ToFloat()
	if total < 0.0 {
		log.Panicln("Cannot represent negative measurements")
	}
	subtracted, err := ToFractional(total)
	if err != nil {
		log.Panicln("Cannot convert value to fractional", err)
	}
	return subtracted
}

func (i Imp) Multiply(quantity int) Imp {
	total := i.ToFloat() * float64(quantity)
	if total < 0.0 {
		log.Panicln("Cannot represent negative measurements")
	}
	multiplied, err := ToFractional(total)
	if err != nil {
		log.Panicln("Cannot convert value to fractional")
	}
	return multiplied
}
