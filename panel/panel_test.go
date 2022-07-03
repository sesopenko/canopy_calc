package panel

import (
	"canopy_calc/dimensions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanelFromTarget(t *testing.T) {
	var scenarios = []struct {
		Description                  string
		InputBuilder                 PanelBuilder
		ExpectedHorizontal           float64
		ExpectedVertical             float64
		ExpectedHorizontalFullLength bool
		ExpectedBoardWidth           dimensions.Imp
		ExpectedCenterColumn         bool
	}{
		{
			Description: "full length horizontal",
			InputBuilder: PanelBuilder{
				BoardWidth: dimensions.Inches(1),
				AssembledDimensions: dimensions.Rectangle{
					Width:  dimensions.Inches(12),
					Height: dimensions.Inches(12),
				},
				HorizontalFullLength: true,
				CenterColumn:         false,
			},
			ExpectedHorizontal:           12.0,
			ExpectedVertical:             10.0,
			ExpectedHorizontalFullLength: true,
			ExpectedBoardWidth:           dimensions.Inches(1),
			ExpectedCenterColumn:         false,
		},
		{
			Description: "short horizontal, full length vertical",
			InputBuilder: PanelBuilder{
				BoardWidth: dimensions.Inches(1),
				AssembledDimensions: dimensions.Rectangle{
					Width:  dimensions.Inches(12),
					Height: dimensions.Inches(12),
				},
				HorizontalFullLength: false,
				CenterColumn:         false,
			},
			ExpectedHorizontal:           10.0,
			ExpectedVertical:             12.0,
			ExpectedHorizontalFullLength: false,
			ExpectedBoardWidth:           dimensions.Inches(1),
			ExpectedCenterColumn:         false,
		},
		{
			Description: "short horizontal, full length vertical, with center column",
			InputBuilder: PanelBuilder{
				BoardWidth: dimensions.Inches(1),
				AssembledDimensions: dimensions.Rectangle{
					Width:  dimensions.Inches(12),
					Height: dimensions.Inches(12),
				},
				HorizontalFullLength: false,
				CenterColumn:         true,
			},
			ExpectedHorizontal:           10.0,
			ExpectedVertical:             12.0,
			ExpectedHorizontalFullLength: false,
			ExpectedBoardWidth:           dimensions.Inches(1),
			ExpectedCenterColumn:         true,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Description, func(t *testing.T) {
			result := scenario.InputBuilder.Build()
			assert.Equal(t, scenario.ExpectedHorizontal, result.Horizontal.ToFloat())
			assert.Equal(t, scenario.ExpectedVertical, result.Vertical.ToFloat())
			assert.Equal(t, scenario.ExpectedHorizontalFullLength, result.HorizontalFullLength)
			assert.Equal(t, scenario.ExpectedBoardWidth, result.BoardWidth)
			assert.Equal(t, scenario.ExpectedCenterColumn, result.CenterColumn)
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
				CenterColumn:         false,
			},
			ExpectedDimensions: dimensions.Rectangle{
				Width:  dimensions.Inches(12),
				Height: dimensions.Inches(12),
			},
		},
		{
			Description: "horizontal full length with center column",
			InputPanel: Panel{
				Horizontal:           dimensions.Inches(12),
				Vertical:             dimensions.Inches(10),
				HorizontalFullLength: false,
				BoardWidth:           dimensions.Inches(1),
				CenterColumn:         true,
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

func TestPanel_GetCuts(t *testing.T) {
	var scenarios = []struct {
		Description  string
		InputPanel   Panel
		ExpectedCuts []dimensions.Imp
	}{
		{
			Description: "no center column",
			InputPanel: Panel{
				Horizontal:           dimensions.BuildImp(12, 3, 8),
				Vertical:             dimensions.BuildImp(14, 1, 4),
				HorizontalFullLength: true,
				BoardWidth:           dimensions.Inches(1),
				CenterColumn:         false,
			},
			ExpectedCuts: []dimensions.Imp{
				dimensions.BuildImp(12, 3, 8),
				dimensions.BuildImp(12, 3, 8),
				dimensions.BuildImp(14, 1, 4),
				dimensions.BuildImp(14, 1, 4),
			},
		},
		{
			Description: "with center column",
			InputPanel: Panel{
				Horizontal:           dimensions.BuildImp(12, 3, 8),
				Vertical:             dimensions.BuildImp(14, 1, 4),
				HorizontalFullLength: true,
				BoardWidth:           dimensions.Inches(1),
				CenterColumn:         true,
			},
			ExpectedCuts: []dimensions.Imp{
				dimensions.BuildImp(12, 3, 8),
				dimensions.BuildImp(12, 3, 8),
				dimensions.BuildImp(14, 1, 4),
				dimensions.BuildImp(14, 1, 4),
				dimensions.BuildImp(14, 1, 4),
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Description, func(t *testing.T) {
			result := s.InputPanel.GetCuts()
			assert.Equal(t, s.ExpectedCuts, result)
		})
	}
}
