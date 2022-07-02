package main

import (
	"canopy_calc/canopy"
	"canopy_calc/dimensions"
	"encoding/json"
	"fmt"
	"log"
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

	outputDimensions, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Panicln("Unable to convert result to JSON")
	}
	log.Println(string(outputDimensions))

	cutList := result.GetCutList()
	fmt.Println("CUT LIST:")
	for _, cut := range cutList {
		fmt.Printf("%d %d/%d\"\n", cut.Inches, cut.Numerator, cut.Denominator)
	}

}
