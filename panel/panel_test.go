package panel

import (
	"canopy_calc/dimensions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanelFromTarget(t *testing.T) {
	var scenarios = []struct {
		Description                  string
		InputSpecifications          PanelBuilder
		ExpectedHorizontal           float64
		ExpectedVertical             float64
		ExpectedHorizontalFullLength bool
		ExpectedBoardWidth           dimensions.Imp
	}{
		{
			Description: "full length horizontal",
			InputSpecifications: PanelBuilder{
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
			ExpectedBoardWidth:           dimensions.Inches(1),
		},
		{
			Description: "short horizontal, full length vertical",
			InputSpecifications: PanelBuilder{
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
			ExpectedBoardWidth:           dimensions.Inches(1),
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Description, func(t *testing.T) {
			result := scenario.InputSpecifications.Build()
			assert.Equal(t, scenario.ExpectedHorizontal, result.Horizontal.ToFloat())
			assert.Equal(t, scenario.ExpectedVertical, result.Vertical.ToFloat())
			assert.Equal(t, scenario.ExpectedHorizontalFullLength, result.HorizontalFullLength)
			assert.Equal(t, scenario.ExpectedBoardWidth, result.BoardWidth)
		})
	}
}

func TestPanel_Dimensions(t *testing.T) {
	var scenarios = []struct {
		Description        string
		InputPanel         Panel
		ExpectedDimensions dimensions.Rectangle
	}{
		{
			Description: "horizontal full length",
			InputPanel: Panel{
				Horizontal:           dimensions.Inches(12),
				Vertical:             dimensions.Inches(10),
				HorizontalFullLength: true,
				BoardWidth:           dimensions.Inches(1),
			},
			ExpectedDimensions: dimensions.Rectangle{
				Width:  dimensions.Inches(12),
				Height: dimensions.Inches(12),
			},
		},
		{
			Description: "horizontal full length",
			InputPanel: Panel{
				Horizontal:           dimensions.Inches(12),
				Vertical:             dimensions.Inches(10),
				HorizontalFullLength: false,
				BoardWidth:           dimensions.Inches(1),
			},
			ExpectedDimensions: dimensions.Rectangle{
				Width:  dimensions.Inches(14),
				Height: dimensions.Inches(10),
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Description, func(t *testing.T) {
			result := s.InputPanel.Dimensions()
			assert.Equal(t, s.ExpectedDimensions.Width.ToFloat(), result.Width.ToFloat())
			assert.Equal(t, s.ExpectedDimensions.Height.ToFloat(), result.Height.ToFloat())
		})
	}
}
