package panel

import (
	"canopy_calc/dimensions"
	"fmt"
)

type PanelBuilder struct {
	BoardWidth           dimensions.Imp
	AssembledDimensions  dimensions.Rectangle
	HorizontalFullLength bool
	CenterColumn         bool
}

type Panel struct {
	Horizontal           dimensions.Imp
	Vertical             dimensions.Imp
	HorizontalFullLength bool
	BoardWidth           dimensions.Imp
	CenterColumn         bool
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
		CenterColumn:         s.CenterColumn,
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

func (p Panel) GetCuts() []dimensions.Imp {
	cuts := []dimensions.Imp{
		p.Horizontal,
		p.Horizontal,
		p.Vertical,
		p.Vertical,
	}
	if p.CenterColumn {
		cuts = append(cuts, p.Vertical)
	}
	return cuts
}

func (p Panel) PrettyPrint(name string) {
	fmt.Println(name)
	fmt.Printf("Horizontal: %s\n", p.Horizontal.Format())
	fmt.Printf("Vertical: %s\n", p.Vertical.Format())
	fmt.Printf("Assembled: %s x %s\n", p.Dimensions().Width.Format(), p.Dimensions().Height.Format())
	if p.CenterColumn {
		fmt.Printf("Has center column\n")
	}
	fmt.Println("")
}
