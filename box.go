package main

import "github.com/gdamore/tcell/v2"

type border struct {
	height, width int
	style         tcell.Style
}

func (b *border) draw(scn tcell.Screen) {
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *border) drawHorizontalEdges(scn tcell.Screen) {
	for x := range b.width {
		scn.SetContent(x, 0, tcell.RuneHLine, nil, b.style)
		scn.SetContent(x, b.height-1, tcell.RuneHLine, nil, b.style)
	}
}

func (b *border) drawVerticalEdges(scn tcell.Screen) {
	for y := range b.height {
		scn.SetContent(0, y, tcell.RuneVLine, nil, b.style)
		scn.SetContent(b.width-1, y, tcell.RuneVLine, nil, b.style)
	}
}

func (b *border) setCorners(scn tcell.Screen) {
	scn.SetContent(0, 0, tcell.RuneULCorner, nil, b.style)
	scn.SetContent(b.width-1, 0, tcell.RuneURCorner, nil, b.style)
	scn.SetContent(0, b.height-1, tcell.RuneLLCorner, nil, b.style)
	scn.SetContent(b.width-1, b.height-1, tcell.RuneLRCorner, nil, b.style)
}

func newBorder(height, width int) *border {
	return &border{height: height, width: width, style: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)}
}
