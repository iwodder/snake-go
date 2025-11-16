package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
	scoreHeight = 2
	scoreFormat = "Score: %d"
)

type board struct {
	upperLeft  Position
	lowerRight Position
	scoreBox   *TextBox
}

func (b *board) draw(scn tcell.Screen) {
	b.fill(scn)
	b.drawBorder(scn)
	b.drawScoreArea(scn)
}

func (b *board) fill(scn tcell.Screen) {
	for x := 0; x < b.width(); x++ {
		for y := 0; y < b.height(); y++ {
			scn.SetContent(x, y, ' ', nil, boardStyle)
		}
	}
}

func (b *board) drawBorder(scn tcell.Screen) {
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *board) drawHorizontalEdges(scn tcell.Screen) {
	for x := 0; x < b.width(); x++ {
		scn.SetContent(x, b.upperLeft.y, tcell.RuneHLine, nil, boardStyle)
		scn.SetContent(x, b.lowerRight.y, tcell.RuneHLine, nil, boardStyle)
	}
}

func (b *board) drawVerticalEdges(scn tcell.Screen) {
	for y := 0; y < b.height(); y++ {
		scn.SetContent(b.upperLeft.x, y, tcell.RuneVLine, nil, boardStyle)
		scn.SetContent(b.lowerRight.x, y, tcell.RuneVLine, nil, boardStyle)
	}
}

func (b *board) setCorners(scn tcell.Screen) {
	scn.SetContent(b.upperLeft.x, b.upperLeft.y, tcell.RuneULCorner, nil, boardStyle)
	scn.SetContent(b.lowerRight.x, b.upperLeft.y, tcell.RuneURCorner, nil, boardStyle)
	scn.SetContent(b.upperLeft.x, b.lowerRight.y, tcell.RuneLLCorner, nil, boardStyle)
	scn.SetContent(b.lowerRight.x, b.lowerRight.y, tcell.RuneLRCorner, nil, boardStyle)
}

func (b *board) drawScoreArea(scn tcell.Screen) {
	b.scoreBox.Draw(scn)

	scn.SetContent(b.upperLeft.x, b.scoreBox.BottomEdge(), tcell.RuneLTee, nil, boardStyle)
	scn.SetContent(b.lowerRight.x, b.scoreBox.BottomEdge(), tcell.RuneRTee, nil, boardStyle)
}

func (b *board) leftEdge() int {
	return b.upperLeft.x + borderWidth
}

func (b *board) rightEdge() int {
	return b.lowerRight.x - borderWidth
}

func (b *board) topEdge() int {
	return b.upperLeft.y + borderWidth + scoreHeight
}

func (b *board) bottomEdge() int {
	return b.lowerRight.y - borderWidth
}

func (b *board) width() int {
	return b.lowerRight.x - b.upperLeft.x + 1
}

func (b *board) height() int {
	return b.lowerRight.y - b.upperLeft.y + 1
}

func (b *board) center() Position {
	return Position{x: b.width() / 2, y: b.height() / 2}
}

func (b *board) isInside(pos Position) bool {
	return pos.x > b.leftEdge() && pos.x < b.rightEdge() && pos.y > b.topEdge() && pos.y < b.bottomEdge()
}

func (b *board) setScore(score uint) {
	b.scoreBox.SetText(fmt.Sprintf(scoreFormat, score))
}

func newBoard(ul, lr Position) *board {
	ret := board{
		upperLeft:  ul,
		lowerRight: lr,
	}

	scoreBox := NewTextBox(fmt.Sprintf(scoreFormat, 0), boardStyle)
	scoreBox.SetHeight(scoreHeight)
	scoreBox.SetWidth(ret.width())
	scoreBox.SetPosition(ret.upperLeft)

	ret.scoreBox = scoreBox
	return &ret
}
