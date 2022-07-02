package main

import (
	"canopy_calc/canopy"
	"canopy_calc/dimensions"
	"fmt"
)

func main() {

	canopyBuilder := canopy.CanopyBuilder{
		AquariumTop: dimensions.Rectangle{
			Width: dimensions.Inches(59),
			Height: dimensions.Imp{
				Inches:      22,
				Numerator:   5,
				Denominator: 8,
			},
		},
		BoardProfile: dimensions.Rectangle{
			Width: dimensions.Imp{
				Inches:      3,
				Numerator:   3,
				Denominator: 8,
			},
			Height: dimensions.Imp{
				Inches:      0,
				Numerator:   3,
				Denominator: 4,
			},
		},
		DesiredClearance:  dimensions.Inches(17),
		WaterlineDistance: dimensions.BuildImp(1, 7, 8),
	}

	result := canopyBuilder.Build()

	fmt.Printf("Board Width: %s\n", canopyBuilder.BoardProfile.Width.Format())
	fmt.Printf("Aquarium top: %s x %s\n", canopyBuilder.AquariumTop.Width.Format(), canopyBuilder.AquariumTop.Height.Format())
	fmt.Println("")

	result.FrontPanel.PrettyPrint("FRONT PANEL")
	result.SidePanel.PrettyPrint("SIDE PANEL")
	result.RearPanel.PrettyPrint("REAR PANEL")
	result.TopPanel.PrettyPrint("TOP PANEL")

	cutList := result.GetCutList()
	fmt.Println("CUT LIST:")
	for _, cut := range cutList {
		fmt.Printf("%d %d/%d\"\n", cut.Inches, cut.Numerator, cut.Denominator)
	}

}
