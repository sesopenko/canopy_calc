package dimensions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImp_ToFloat(t *testing.T) {
	var scenarios = []struct {
		Description string
		Input       Imp
		Expected    float64
	}{
		{
			Description: "non fractional",
			Input:       Inches(12),
			Expected:    float64(12.0),
		},
		{
			Description: "Half inch",
			Input: Imp{
				Inches:      0,
				Numerator:   1,
				Denominator: 2,
			},
			Expected: float64(0.5),
		},
		{
			Description: "1.5 inches",
			Input: Imp{
				Inches:      1,
				Numerator:   1,
				Denominator: 2,
			},
			Expected: float64(1.5),
		},
		{
			Description: "Fraction > 1",
			Input: Imp{
				Inches:      0,
				Numerator:   3,
				Denominator: 2,
			},
			Expected: float64(1.5),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Description, func(t *testing.T) {
			assert.Equal(t, scenario.Expected, scenario.Input.ToFloat(), "Should get expected value")
		})
	}

}

func TestToFractional(t *testing.T) {
	var scenarios = []struct {
		Description         string
		Input               float64
		ExpectedInches      int
		ExpectedNumerator   int
		ExpectedDenominator int
	}{
		{
			Description:         "0.5",
			Input:               float64(0.5),
			ExpectedInches:      0,
			ExpectedNumerator:   1,
			ExpectedDenominator: 2,
		},
		{
			Description:         "1.5",
			Input:               float64(1.5),
			ExpectedInches:      1,
			ExpectedNumerator:   1,
			ExpectedDenominator: 2,
		},
		{
			Description:         "1.25",
			Input:               float64(1.25),
			ExpectedInches:      1,
			ExpectedNumerator:   1,
			ExpectedDenominator: 4,
		},
		{
			Description:         "1 3/32",
			Input:               1.0 + (3.0 / 32.0),
			ExpectedInches:      1,
			ExpectedNumerator:   3,
			ExpectedDenominator: 32,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Description, func(t *testing.T) {
			result, err := ToFractional(scenario.Input)
			assert.Nil(t, err)
			assert.Equal(t, scenario.ExpectedInches, result.Inches, "Should get expected inches")
			assert.Equal(t, scenario.ExpectedNumerator, result.Numerator, "Should get expected numerator")
			assert.Equal(t, scenario.ExpectedDenominator, result.Denominator, "Should get expected denominator")
		})
	}
}
