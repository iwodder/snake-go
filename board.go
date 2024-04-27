package main

import "github.com/gdamore/tcell/v2"

type board struct {
	bounds boundary
}

func (b *board) draw(scn tcell.Screen) {
	b.fill(scn)
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *board) fill(scn tcell.Screen) {
	for x := range b.bounds.width() {
		for y := range b.bounds.height() {
			scn.SetContent(x, y, ' ', nil, boardStyle)
		}
	}
}

func (b *board) drawHorizontalEdges(scn tcell.Screen) {
	for x := range b.bounds.width() {
		scn.SetContent(x, 0, tcell.RuneHLine, nil, boardStyle)
		scn.SetContent(x, b.bounds.height()-1, tcell.RuneHLine, nil, boardStyle)
	}
}

func (b *board) drawVerticalEdges(scn tcell.Screen) {
	for y := range b.bounds.height() {
		scn.SetContent(0, y, tcell.RuneVLine, nil, boardStyle)
		scn.SetContent(b.bounds.width()-1, y, tcell.RuneVLine, nil, boardStyle)
	}
}

func (b *board) setCorners(scn tcell.Screen) {
	scn.SetContent(0, 0, tcell.RuneULCorner, nil, boardStyle)
	scn.SetContent(b.bounds.width()-1, 0, tcell.RuneURCorner, nil, boardStyle)
	scn.SetContent(0, b.bounds.height()-1, tcell.RuneLLCorner, nil, boardStyle)
	scn.SetContent(b.bounds.width()-1, b.bounds.height()-1, tcell.RuneLRCorner, nil, boardStyle)
}

func (b *board) boundary() boundary {
	return b.bounds.shrink(1)
}

func newBoard(ul, lr pos) *board {
	return &board{
		bounds: boundary{upperLeft: ul, lowerRight: lr},
	}
}
