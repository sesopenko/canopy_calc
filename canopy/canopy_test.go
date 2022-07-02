package canopy

import (
	"canopy_calc/dimensions"
	"canopy_calc/panel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildCanopy(t *testing.T) {
	var scenarios = []struct {
		Description        string
		InputBuilder       CanopyBuilder
		ExpectedFrontPanel panel.Panel
	}{
		{
			Description: "Sean's Tank",
			InputBuilder: CanopyBuilder{
				AquariumTop: dimensions.Rectangle{
					Width:  dimensions.Inches(59),
					Height: dimensions.BuildImp(22, 5, 8),
				},
				BoardProfile: dimensions.Rectangle{
					Width:  dimensions.BuildImp(3, 3, 8),
					Height: dimensions.BuildImp(0, 3, 4),
				},
				DesiredClearance:  dimensions.Inches(17),
				WaterlineDistance: dimensions.BuildImp(1, 7, 8),
			},
			ExpectedFrontPanel: panel.Panel{
				// Should extend beyond the width of the tank, to cover the edges of the other panels
				Horizontal: dimensions.Inches(59).
					Add(RestingAllowance.Multiply(2)).
					Add(dimensions.BuildImp(0, 3, 4).Multiply(2)),
				Vertical: dimensions.Inches(17).
					Add(dimensions.BuildImp(1, 7, 8)).
					Add(dimensions.BuildImp(0, 3, 4)),
				HorizontalFullLength: true,
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Description, func(t *testing.T) {
			result := s.InputBuilder.Build()
			assert.Equal(t, s.ExpectedFrontPanel.HorizontalFullLength, result.FrontPanel.HorizontalFullLength)
			assert.Equal(t, s.ExpectedFrontPanel.Horizontal, result.FrontPanel.Horizontal)
		})
	}
}
