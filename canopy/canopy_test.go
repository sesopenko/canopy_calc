package canopy

import (
	"canopy_calc/dimensions"
	"canopy_calc/panel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildCanopy(t *testing.T) {
	tankTopWidth := dimensions.Inches(59)
	tankTopHeight := dimensions.BuildImp(22, 5, 8)
	targetClearance := dimensions.Inches(17)
	boardWidth := dimensions.BuildImp(3, 3, 8)
	boardDepth := dimensions.BuildImp(0, 3, 4)
	allowance := dimensions.BuildImp(0, 1, 8).Multiply(2)
	waterLine := dimensions.BuildImp(1, 7, 8)
	var scenarios = []struct {
		Description        string
		InputBuilder       CanopyBuilder
		ExpectedFrontPanel panel.Panel
		ExpectedSidePanel  panel.Panel
		ExpectedRearPanel  panel.Panel
		ExpectedTopPanel   panel.Panel
	}{
		{
			Description: "Sean's Tank",
			InputBuilder: CanopyBuilder{
				AquariumTop: dimensions.Rectangle{
					Width:  tankTopWidth,
					Height: tankTopHeight,
				},
				BoardProfile: dimensions.Rectangle{
					Width:  boardWidth,
					Height: boardDepth,
				},
				DesiredClearance:  targetClearance,
				WaterlineDistance: waterLine,
			},
			ExpectedFrontPanel: panel.PanelBuilder{
				BoardWidth: boardWidth,
				AssembledDimensions: dimensions.Rectangle{
					Width: tankTopWidth.
						Add(allowance).
						// Conceal the side panels.
						Add(boardDepth.Multiply(2)),
					Height: targetClearance.
						// Add water line.
						Add(waterLine).
						// Conceal the top panel.
						Add(boardDepth),
				},
				HorizontalFullLength: true,
			}.Build(),
			// This panel sits underneath the top panel, and conceals the rear panel.
			ExpectedSidePanel: panel.PanelBuilder{
				BoardWidth: boardWidth,
				AssembledDimensions: dimensions.Rectangle{
					Width: tankTopHeight.
						// Conceal the rear panel:
						Add(boardDepth).
						Add(allowance),
					Height: targetClearance.
						Add(waterLine),
				},
				HorizontalFullLength: true,
			}.Build(),
			// This panel sits rests under the top panel, and is concealed by the
			// front panel.
			ExpectedRearPanel: panel.PanelBuilder{
				BoardWidth: boardWidth,
				AssembledDimensions: dimensions.Rectangle{
					// Is concealed by side panels, sitting between them.
					Width: tankTopWidth.
						Add(allowance),
					Height: targetClearance.
						Add(waterLine),
				},
				HorizontalFullLength: true,
			}.Build(),
			// This panel is concealed by the front panel, and sits on top of the side
			// and rear panels.
			ExpectedTopPanel: panel.PanelBuilder{
				BoardWidth: boardWidth,
				AssembledDimensions: dimensions.Rectangle{
					Width: tankTopWidth.
						Add(allowance).
						// Sits between the side panels
						Add(boardDepth.Multiply(2)),
					Height: tankTopHeight.
						Add(allowance).
						Add(boardDepth),
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
			assert.Equal(t, s.ExpectedTopPanel, result.TopPanel, "Should get expected top panel")
		})
	}
}

func TestCanopy_GetCutList(t *testing.T) {
	var scenarios = []struct {
		Description string
		InputCanopy Canopy
		ExpectedQty int
	}{
		{
			Description: "Simple panels",
			InputCanopy: Canopy{
				FrontPanel: panel.PanelBuilder{
					BoardWidth: dimensions.Inches(1),
					AssembledDimensions: dimensions.Rectangle{
						Width:  dimensions.Inches(12),
						Height: dimensions.Inches(12),
					},
					HorizontalFullLength: true,
				}.Build(),
				SidePanel: panel.PanelBuilder{
					BoardWidth: dimensions.Inches(1),
					AssembledDimensions: dimensions.Rectangle{
						Width:  dimensions.Inches(12),
						Height: dimensions.Inches(12),
					},
					HorizontalFullLength: true,
				}.Build(),
				RearPanel: panel.PanelBuilder{
					BoardWidth: dimensions.Inches(1),
					AssembledDimensions: dimensions.Rectangle{
						Width:  dimensions.Inches(12),
						Height: dimensions.Inches(12),
					},
					HorizontalFullLength: true,
				}.Build(),
				TopPanel: panel.PanelBuilder{
					BoardWidth: dimensions.Inches(1),
					AssembledDimensions: dimensions.Rectangle{
						Width:  dimensions.Inches(12),
						Height: dimensions.Inches(12),
					},
					HorizontalFullLength: true,
				}.Build(),
			},
			ExpectedQty: 20,
		},
	}

	for _, s := range scenarios {
		t.Run(s.Description, func(t *testing.T) {
			result := s.InputCanopy.GetCutList()
			assert.Len(t, result, s.ExpectedQty)
		})
	}
}
