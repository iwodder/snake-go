package main

import "github.com/gdamore/tcell/v2"

type board struct {
	height, width int
	style         tcell.Style
}

func (b *board) draw(scn tcell.Screen) {
	b.fill(scn)
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *board) fill(scn tcell.Screen) {
	for x := range b.width {
		for y := range b.height {
			scn.SetContent(x, y, ' ', nil, b.style)
		}
	}
}

func (b *board) drawHorizontalEdges(scn tcell.Screen) {
	for x := range b.width {
		scn.SetContent(x, 0, tcell.RuneHLine, nil, b.style)
		scn.SetContent(x, b.height-1, tcell.RuneHLine, nil, b.style)
	}
}

func (b *board) drawVerticalEdges(scn tcell.Screen) {
	for y := range b.height {
		scn.SetContent(0, y, tcell.RuneVLine, nil, b.style)
		scn.SetContent(b.width-1, y, tcell.RuneVLine, nil, b.style)
	}
}

func (b *board) setCorners(scn tcell.Screen) {
	scn.SetContent(0, 0, tcell.RuneULCorner, nil, b.style)
	scn.SetContent(b.width-1, 0, tcell.RuneURCorner, nil, b.style)
	scn.SetContent(0, b.height-1, tcell.RuneLLCorner, nil, b.style)
	scn.SetContent(b.width-1, b.height-1, tcell.RuneLRCorner, nil, b.style)
}

func newBoard(height, width int) *board {
	return &board{height: height, width: width, style: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)}
}
