package ui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
)

type GameBoardRenderer struct {
	composite
	ul  Position
	lr  Position
	hud *Hud
}

func (b *GameBoardRenderer) Draw(scn tcell.Screen) {
	drawBorder(b.ul, b.Width(), b.Height(), boardStyle, scn)
	b.drawScoreArea(scn)
	b.composite.Draw(scn)
}

func (b *GameBoardRenderer) drawScoreArea(scn tcell.Screen) {
	b.hud.Draw(scn)

	for i := range b.Width() {
		scn.SetContent(b.ul.X+i, b.hud.Bottom(), tcell.RuneHLine, nil, boardStyle)
	}
	scn.SetContent(b.ul.X, b.hud.Bottom(), tcell.RuneLTee, nil, boardStyle)
	scn.SetContent(b.lr.X, b.hud.Bottom(), tcell.RuneRTee, nil, boardStyle)
}

func (b *GameBoardRenderer) Left() int {
	return b.ul.X
}

func (b *GameBoardRenderer) Right() int {
	return b.lr.X
}

func (b *GameBoardRenderer) Top() int {
	return b.ul.Y + b.hud.Height() + borderWidth
}

func (b *GameBoardRenderer) Bottom() int {
	return b.lr.Y
}

func (b *GameBoardRenderer) Width() int {
	return b.lr.X - b.ul.X + 1
}

func (b *GameBoardRenderer) Height() int {
	return b.lr.Y - b.ul.Y + 1
}

func (b *GameBoardRenderer) setHud(hud *Hud) {
	if b.hud != nil {
		_ = b.Remove(b.hud)
	}
	_ = b.Add(hud)
	b.hud = hud
}

func (b *GameBoardRenderer) ScoreBox() *TextBox {
	return b.hud.ScoreBox()
}

func (b *GameBoardRenderer) LivesBox() *TextBox {
	return b.hud.LivesBox()
}

func NewGameBoardRenderer(ul, lr Position) *GameBoardRenderer {
	ret := GameBoardRenderer{
		ul: ul,
		lr: lr,
	}
	ret.setHud(NewHud(Position{X: ul.X + 1, Y: ul.Y + 1}, 0, ret.Width()-2))
	return &ret
}
