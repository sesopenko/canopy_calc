package canopy

import (
	"canopy_calc/dimensions"
	"canopy_calc/panel"
	"sort"
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
	TopPanel   panel.Panel
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
			Height: sidePanel.Dimensions().Height,
		},
		HorizontalFullLength: true,
		CenterColumn:         true,
	}

	rearPanel := rearPanelBuilder.Build()

	topPanelBuilder := panel.PanelBuilder{
		BoardWidth: b.BoardProfile.Width,
		AssembledDimensions: dimensions.Rectangle{
			Width:  frontPanel.Dimensions().Width,
			Height: sidePanel.Dimensions().Width,
		},
		HorizontalFullLength: true,
	}

	topPanel := topPanelBuilder.Build()

	return Canopy{
		FrontPanel: frontPanel,
		SidePanel:  sidePanel,
		RearPanel:  rearPanel,
		TopPanel:   topPanel,
	}
}

func (c Canopy) GetCutList() []dimensions.Imp {
	cutList := []dimensions.Imp{}
	cutList = append(cutList, c.TopPanel.GetCuts()...)
	cutList = append(cutList, c.FrontPanel.GetCuts()...)
	cutList = append(cutList, c.SidePanel.GetCuts()...)
	cutList = append(cutList, c.SidePanel.GetCuts()...)
	cutList = append(cutList, c.RearPanel.GetCuts()...)
	sort.SliceStable(cutList, func(i, j int) bool {
		return cutList[i].ToFloat() > cutList[j].ToFloat()
	})
	return cutList
}
