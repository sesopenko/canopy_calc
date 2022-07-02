package component

import "canopy_calc/dimensions"

type PanelSpecification struct {
	BoardWidth           dimensions.Imp
	AssembledDimensions  dimensions.Rectangle
	HorizontalFullLength bool
}

type Panel struct {
	Horizontal           dimensions.Imp
	Vertical             dimensions.Imp
	HorizontalFullLength bool
}

func PanelFromTarget(specs PanelSpecification) Panel {
	var verticalLength dimensions.Imp
	var horizontalLength dimensions.Imp
	if specs.HorizontalFullLength {
		verticalLength = specs.AssembledDimensions.Height.Subtract(specs.BoardWidth).Subtract(specs.BoardWidth)
		horizontalLength = specs.AssembledDimensions.Width
	} else {
		verticalLength = specs.AssembledDimensions.Height
		horizontalLength = specs.AssembledDimensions.Width.Subtract(specs.BoardWidth).Subtract(specs.BoardWidth)
	}
	return Panel{
		Horizontal:           horizontalLength,
		Vertical:             verticalLength,
		HorizontalFullLength: specs.HorizontalFullLength,
	}
}
