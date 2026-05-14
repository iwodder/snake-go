package ui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	borderWidth = 1
)

type GameBoardRenderer struct {
	composite
	ul            Position
	height, width int
	hud           *Hud
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
	scn.SetContent(b.Left(), b.hud.Bottom(), tcell.RuneLTee, nil, boardStyle)
	scn.SetContent(b.Right(), b.hud.Bottom(), tcell.RuneRTee, nil, boardStyle)
}

func (b *GameBoardRenderer) Left() int {
	return b.ul.X
}

func (b *GameBoardRenderer) Right() int {
	return b.ul.X + b.width - 1
}

func (b *GameBoardRenderer) Top() int {
	return b.ul.Y + b.hud.Height() + borderWidth
}

func (b *GameBoardRenderer) Bottom() int {
	return b.ul.Y + b.height - borderWidth
}

func (b *GameBoardRenderer) Width() int {
	return b.width
}

func (b *GameBoardRenderer) Height() int {
	return b.height
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

func NewGameBoardRenderer(ul Position, width int, height int) *GameBoardRenderer {
	ret := GameBoardRenderer{
		ul:     ul,
		width:  width,
		height: height,
	}
	ret.setHud(NewHud(Position{X: ul.X + 1, Y: ul.Y + 1}, 0, ret.Width()-2))
	return &ret
}
