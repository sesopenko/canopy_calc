package main

import (
	"canopy_calc/dimensions"
	"log"
)

func main() {
	// top dimensions 59" x 22 5/8"
	//aquariumTop := dimensions.Rectangle{
	//	X: dimensions.Imp{
	//		Inches:      59,
	//		Numerator:   0,
	//		Denominator: 32,
	//	},
	//	Y: dimensions.Imp{
	//		Inches:      22,
	//		Numerator:   5,
	//		Denominator: 8,
	//	},
	//}
	// board is 3 3/8" x 3/4"
	boardProfile := dimensions.Rectangle{
		X: dimensions.Imp{
			Inches:      3,
			Numerator:   3,
			Denominator: 8,
		},
		Y: dimensions.Imp{
			Inches:      0,
			Numerator:   3,
			Denominator: 4,
		},
	}

	// desired height 17"
	desiredClearance := dimensions.Inches(17)

	// distance from top of tank to water line, to completely cover water line
	waterlineDistance := dimensions.BuildImp(1, 7, 8)

	frontPanelHeight := desiredClearance.Add(waterlineDistance)
	frontPanelVerticalLength := frontPanelHeight.Subtract(boardProfile.X).Subtract(boardProfile.X)
	log.Printf("front panel vertical length: %d %d/%d\"",
		frontPanelVerticalLength.Inches,
		frontPanelVerticalLength.Numerator,
		frontPanelVerticalLength.Denominator)
}
