package main

import (
	"github.com/gdamore/tcell/v2"
)

const borderWidth = 1

type board struct {
	upperLeft  pos
	lowerRight pos
}

func (b *board) draw(scn tcell.Screen) {
	b.fill(scn)
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *board) fill(scn tcell.Screen) {
	for x := range b.width() {
		for y := range b.height() {
			scn.SetContent(x, y, ' ', nil, boardStyle)
		}
	}
}

func (b *board) drawHorizontalEdges(scn tcell.Screen) {
	for x := range b.width() {
		scn.SetContent(x, 0, tcell.RuneHLine, nil, boardStyle)
		scn.SetContent(x, b.height()-1, tcell.RuneHLine, nil, boardStyle)
	}
}

func (b *board) drawVerticalEdges(scn tcell.Screen) {
	for y := range b.height() {
		scn.SetContent(0, y, tcell.RuneVLine, nil, boardStyle)
		scn.SetContent(b.width()-1, y, tcell.RuneVLine, nil, boardStyle)
	}
}

func (b *board) setCorners(scn tcell.Screen) {
	scn.SetContent(0, 0, tcell.RuneULCorner, nil, boardStyle)
	scn.SetContent(b.width()-1, 0, tcell.RuneURCorner, nil, boardStyle)
	scn.SetContent(0, b.height()-1, tcell.RuneLLCorner, nil, boardStyle)
	scn.SetContent(b.width()-1, b.height()-1, tcell.RuneLRCorner, nil, boardStyle)
}

func (b *board) leftEdge() int {
	return b.upperLeft.x + borderWidth
}

func (b *board) rightEdge() int {
	return b.lowerRight.x - borderWidth
}

func (b *board) topEdge() int {
	return b.upperLeft.y + borderWidth
}

func (b *board) bottomEdge() int {
	return b.lowerRight.y - borderWidth
}

func (b *board) width() int {
	return b.lowerRight.x - b.upperLeft.x
}

func (b *board) height() int {
	return b.lowerRight.y - b.upperLeft.y
}

func (b *board) center() pos {
	return pos{x: b.width() / 2, y: b.height() / 2}
}

func (b *board) isInside(p pos) bool {
	return p.x > b.leftEdge() && p.x < b.rightEdge() && p.y > b.topEdge() && p.y < b.bottomEdge()
}

func newBoard(ul, lr pos) *board {
	return &board{
		upperLeft:  ul,
		lowerRight: lr,
	}
}
