package canopy

import (
	"canopy_calc/dimensions"
	"canopy_calc/panel"
)

var RestingAllowance = dimensions.Imp{
	Inches:      0,
	Numerator:   1,
	Denominator: 8,
}

type CanopyBuilder struct {
	AquariumTop       dimensions.Rectangle
	BoardProfile      dimensions.Rectangle
	DesiredClearance  dimensions.Imp
	WaterlineDistance dimensions.Imp
}

type Canopy struct {
	FrontPanel panel.Panel
	SidePanel  panel.Panel
	RearPanel  panel.Panel
}

func (b CanopyBuilder) Build() Canopy {
	frontPanelBuilder := panel.PanelBuilder{
		BoardWidth: b.BoardProfile.Width,
		AssembledDimensions: dimensions.Rectangle{
			Width: b.AquariumTop.Width.
				Add(RestingAllowance.Multiply(2)).
				Add(b.BoardProfile.Height.Multiply(2)),
			Height: b.DesiredClearance.Add(b.WaterlineDistance).Add(b.BoardProfile.Height),
		},
		HorizontalFullLength: true,
	}

	frontPanel := frontPanelBuilder.Build()

	sidePanelBuilder := panel.PanelBuilder{
		BoardWidth: b.BoardProfile.Width,
		AssembledDimensions: dimensions.Rectangle{
			Width: b.AquariumTop.Height.Add(RestingAllowance.Multiply(2)).
				// Extend off the back to enclose the rear panel.
				Add(b.BoardProfile.Height),
			Height: frontPanel.Dimensions().Height.
				// Top panel sits on top
				Subtract(b.BoardProfile.Height),
		},
		HorizontalFullLength: true,
	}
	sidePanel := sidePanelBuilder.Build()

	rearPanelBuilder := panel.PanelBuilder{
		BoardWidth: b.BoardProfile.Width,
		AssembledDimensions: dimensions.Rectangle{
			Width: b.AquariumTop.Width.
				Add(RestingAllowance.Multiply(2)),
			Height: sidePanel.Dimensions().Height.
				// Give it space for the top panel to sit on top of it
				Subtract(dimensions.BuildImp(0, 3, 4)),
		},
		HorizontalFullLength: true,
	}

	rearPanel := rearPanelBuilder.Build()

	return Canopy{
		FrontPanel: frontPanel,
		SidePanel:  sidePanel,
		RearPanel:  rearPanel,
	}
}
