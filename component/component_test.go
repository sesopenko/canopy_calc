package component

import (
	"canopy_calc/dimensions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanelFromTarget(t *testing.T) {
	var scenarios = []struct {
		Description                  string
		InputSpecifications          PanelSpecification
		ExpectedHorizontal           float64
		ExpectedVertical             float64
		ExpectedHorizontalFullLength bool
	}{
		{
			Description: "full length horizontal",
			InputSpecifications: PanelSpecification{
				BoardWidth: dimensions.Inches(1),
				AssembledDimensions: dimensions.Rectangle{
					Width:  dimensions.Inches(12),
					Height: dimensions.Inches(12),
				},
				HorizontalFullLength: true,
			},
			ExpectedHorizontal:           12.0,
			ExpectedVertical:             10.0,
			ExpectedHorizontalFullLength: true,
		},
		{
			Description: "short horizontal, full length vertical",
			InputSpecifications: PanelSpecification{
				BoardWidth: dimensions.Inches(1),
				AssembledDimensions: dimensions.Rectangle{
					Width:  dimensions.Inches(12),
					Height: dimensions.Inches(12),
				},
				HorizontalFullLength: false,
			},
			ExpectedHorizontal:           10.0,
			ExpectedVertical:             12.0,
			ExpectedHorizontalFullLength: false,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Description, func(t *testing.T) {
			result := PanelFromTarget(scenario.InputSpecifications)
			assert.Equal(t, scenario.ExpectedHorizontal, result.Horizontal.ToFloat())
			assert.Equal(t, scenario.ExpectedVertical, result.Vertical.ToFloat())
			assert.Equal(t, scenario.ExpectedHorizontalFullLength, result.HorizontalFullLength)
		})
	}
}
