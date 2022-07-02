package panel

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
	BoardWidth           dimensions.Imp
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
		BoardWidth:           s.BoardWidth,
	}
}

func (p Panel) Dimensions() dimensions.Rectangle {
	if p.HorizontalFullLength {
		return dimensions.Rectangle{
			Width:  p.Horizontal,
			Height: p.Vertical.Add(p.BoardWidth.Multiply(2)),
		}
	} else {
		return dimensions.Rectangle{
			Height: p.Vertical,
			Width:  p.Horizontal.Add(p.BoardWidth.Multiply(2)),
		}
	}
}
