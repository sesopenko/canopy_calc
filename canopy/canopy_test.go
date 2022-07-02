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
		ExpectedSidePanel  panel.Panel
		ExpectedRearPanel  panel.Panel
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
			ExpectedFrontPanel: panel.PanelBuilder{
				BoardWidth: dimensions.BuildImp(3, 3, 8),
				AssembledDimensions: dimensions.Rectangle{
					Width: dimensions.Inches(59).
						Add(dimensions.BuildImp(0, 1, 8).Multiply(2)).
						Add(dimensions.BuildImp(0, 3, 4).Multiply(2)),
					Height: dimensions.Inches(17).
						Add(dimensions.BuildImp(1, 7, 8)).
						Add(dimensions.BuildImp(0, 3, 4)),
				},
				HorizontalFullLength: true,
			}.Build(),
			ExpectedSidePanel: panel.PanelBuilder{
				BoardWidth: dimensions.BuildImp(3, 3, 8),
				AssembledDimensions: dimensions.Rectangle{
					Width: dimensions.BuildImp(22, 5, 8).
						Add(dimensions.BuildImp(0, 3, 4)).
						Add(dimensions.BuildImp(0, 1, 8).Multiply(2)),
					Height: dimensions.Inches(17).
						Add(dimensions.BuildImp(1, 7, 8)),
				},
				HorizontalFullLength: true,
			}.Build(),
			ExpectedRearPanel: panel.PanelBuilder{
				BoardWidth: dimensions.BuildImp(3, 3, 8),
				AssembledDimensions: dimensions.Rectangle{
					Width: dimensions.Inches(59).
						Add(dimensions.BuildImp(0, 1, 8).Multiply(2)),
					Height: dimensions.Inches(17).
						Add(dimensions.BuildImp(1, 7, 8)).
						Subtract(dimensions.BuildImp(0, 3, 4)),
				},
				HorizontalFullLength: true,
			}.Build(),
		},
	}

	for _, s := range scenarios {
		t.Run(s.Description, func(t *testing.T) {
			result := s.InputBuilder.Build()
			assert.Equal(t, s.ExpectedFrontPanel, result.FrontPanel, "Should get expected front panel")
			assert.Equal(t, s.ExpectedSidePanel, result.SidePanel, "Should get expected side panel")
			assert.Equal(t, s.ExpectedRearPanel, result.RearPanel, "Should get expected rear panel")
		})
	}
}
