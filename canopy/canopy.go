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

	return Canopy{
		FrontPanel: frontPanel,
	}
}
