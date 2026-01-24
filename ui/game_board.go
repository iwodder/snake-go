package ui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
)

type GameBoard struct {
	composite
	ul  Position
	lr  Position
	hud *Hud
}

func (b *GameBoard) Draw(scn tcell.Screen) {
	b.drawBorder(scn)
	b.drawScoreArea(scn)
}

func (b *GameBoard) drawBorder(scn tcell.Screen) {
	b.drawHorizontalEdges(scn)
	b.drawVerticalEdges(scn)
	b.setCorners(scn)
}

func (b *GameBoard) drawScoreArea(scn tcell.Screen) {
	b.hud.Draw(scn)

	for i := range b.Width() {
		scn.SetContent(b.ul.X+i, b.hud.Bottom(), tcell.RuneHLine, nil, boardStyle)
	}
	scn.SetContent(b.ul.X, b.hud.Bottom(), tcell.RuneLTee, nil, boardStyle)
	scn.SetContent(b.lr.X, b.hud.Bottom(), tcell.RuneRTee, nil, boardStyle)
}

func (b *GameBoard) drawHorizontalEdges(scn tcell.Screen) {
	for x := 0; x < b.Width(); x++ {
		scn.SetContent(x, b.ul.Y, tcell.RuneHLine, nil, boardStyle)
		scn.SetContent(x, b.lr.Y, tcell.RuneHLine, nil, boardStyle)
	}
}

func (b *GameBoard) drawVerticalEdges(scn tcell.Screen) {
	for y := 0; y < b.Height(); y++ {
		scn.SetContent(b.ul.X, y, tcell.RuneVLine, nil, boardStyle)
		scn.SetContent(b.lr.X, y, tcell.RuneVLine, nil, boardStyle)
	}
}

func (b *GameBoard) setCorners(scn tcell.Screen) {
	scn.SetContent(b.ul.X, b.ul.Y, tcell.RuneULCorner, nil, boardStyle)
	scn.SetContent(b.lr.X, b.ul.Y, tcell.RuneURCorner, nil, boardStyle)
	scn.SetContent(b.ul.X, b.lr.Y, tcell.RuneLLCorner, nil, boardStyle)
	scn.SetContent(b.lr.X, b.lr.Y, tcell.RuneLRCorner, nil, boardStyle)
}

func (b *GameBoard) Left() int {
	return b.ul.X + borderWidth
}

func (b *GameBoard) Right() int {
	return b.lr.X - borderWidth
}

func (b *GameBoard) Top() int {
	return b.ul.Y + borderWidth + b.hud.Height() + borderWidth
}

func (b *GameBoard) Bottom() int {
	return b.lr.Y - borderWidth
}

func (b *GameBoard) Width() int {
	return b.lr.X - b.ul.X + 1
}

func (b *GameBoard) Height() int {
	return b.lr.Y - b.ul.Y + 1
}

func (b *GameBoard) Center() Position {
	return Position{X: b.Width() / 2, Y: b.Height() / 2}
}

func (b *GameBoard) IsInside(pos Position) bool {
	return pos.X > b.Left() && pos.X < b.Right() && pos.Y > b.Top() && pos.Y < b.Bottom()
}

func (b *GameBoard) setHud(hud *Hud) {
	if b.hud != nil {
		_ = b.Remove(b.hud)
	}
	_ = b.Add(hud)
	b.hud = hud
}

func (b *GameBoard) ScoreBox() *TextBox {
	return b.hud.ScoreBox()
}

func (b *GameBoard) LivesBox() *TextBox {
	return b.hud.LivesBox()
}

func NewGameBoard(ul, lr Position) *GameBoard {
	ret := GameBoard{
		ul: ul,
		lr: lr,
	}
	ret.setHud(NewHud(Position{X: ul.X + 1, Y: ul.Y + 1}, 0, ret.Width()-2))
	return &ret
}
