package component

import "canopy_calc/dimensions"

type PanelBuilder struct {
	BoardWidth           dimensions.Imp
	AssembledDimensions  dimensions.Rectangle
	HorizontalFullLength bool
}

type Panel struct {
	Horizontal           dimensions.Imp
	Vertical             dimensions.Imp
	HorizontalFullLength bool
}

func (s PanelBuilder) Build() Panel {
	var verticalLength dimensions.Imp
	var horizontalLength dimensions.Imp
	if s.HorizontalFullLength {
		verticalLength = s.AssembledDimensions.Height.Subtract(s.BoardWidth).Subtract(s.BoardWidth)
		horizontalLength = s.AssembledDimensions.Width
	} else {
		verticalLength = s.AssembledDimensions.Height
		horizontalLength = s.AssembledDimensions.Width.Subtract(s.BoardWidth).Subtract(s.BoardWidth)
	}
	return Panel{
		Horizontal:           horizontalLength,
		Vertical:             verticalLength,
		HorizontalFullLength: s.HorizontalFullLength,
	}
}
