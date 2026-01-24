package main

import (
	"fmt"
	"snake/ui"

	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
	livesFormat = "Lives: %d"
	scoreFormat = "Score: %d"
)

type board struct {
	upperLeft  Position
	lowerRight Position
	hud        *ui.Hud
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

func (b *board) drawScoreArea(scn tcell.Screen) {
	b.hud.Draw(scn)

	for i := range b.width() {
		scn.SetContent(b.upperLeft.X+i, b.hud.Bottom(), tcell.RuneHLine, nil, boardStyle)
	}
	scn.SetContent(b.upperLeft.X, b.hud.Bottom(), tcell.RuneLTee, nil, boardStyle)
	scn.SetContent(b.lowerRight.X, b.hud.Bottom(), tcell.RuneRTee, nil, boardStyle)
}

func (b *board) drawHorizontalEdges(scn tcell.Screen) {
	for x := 0; x < b.width(); x++ {
		scn.SetContent(x, b.upperLeft.Y, tcell.RuneHLine, nil, boardStyle)
		scn.SetContent(x, b.lowerRight.Y, tcell.RuneHLine, nil, boardStyle)
	}
}

func (b *board) drawVerticalEdges(scn tcell.Screen) {
	for y := 0; y < b.height(); y++ {
		scn.SetContent(b.upperLeft.X, y, tcell.RuneVLine, nil, boardStyle)
		scn.SetContent(b.lowerRight.X, y, tcell.RuneVLine, nil, boardStyle)
	}
}

func (b *board) setCorners(scn tcell.Screen) {
	scn.SetContent(b.upperLeft.X, b.upperLeft.Y, tcell.RuneULCorner, nil, boardStyle)
	scn.SetContent(b.lowerRight.X, b.upperLeft.Y, tcell.RuneURCorner, nil, boardStyle)
	scn.SetContent(b.upperLeft.X, b.lowerRight.Y, tcell.RuneLLCorner, nil, boardStyle)
	scn.SetContent(b.lowerRight.X, b.lowerRight.Y, tcell.RuneLRCorner, nil, boardStyle)
}

func (b *board) leftEdge() int {
	return b.upperLeft.X + borderWidth
}

func (b *board) rightEdge() int {
	return b.lowerRight.X - borderWidth
}

func (b *board) topEdge() int {
	return b.upperLeft.Y + borderWidth + b.hud.Height() + borderWidth
}

func (b *board) bottomEdge() int {
	return b.lowerRight.Y - borderWidth
}

func (b *board) width() int {
	return b.lowerRight.X - b.upperLeft.X + 1
}

func (b *board) height() int {
	return b.lowerRight.Y - b.upperLeft.Y + 1
}

func (b *board) center() Position {
	return Position{X: b.width() / 2, Y: b.height() / 2}
}

func (b *board) isInside(pos Position) bool {
	return pos.X > b.leftEdge() && pos.X < b.rightEdge() && pos.Y > b.topEdge() && pos.Y < b.bottomEdge()
}

func (b *board) setScore(score uint) {
	b.hud.ScoreBox().SetText(fmt.Sprintf(scoreFormat, score))
}

func (b *board) setLives(lives uint) {
	b.hud.LivesBox().SetText(fmt.Sprintf(livesFormat, lives))
}

func newBoard(ul, lr Position) *board {
	ret := board{
		upperLeft:  ul,
		lowerRight: lr,
	}
	ret.hud = ui.NewHud(Position{X: ul.X + 1, Y: ul.Y + 1}, 0, ret.width()-2)
	return &ret
}
