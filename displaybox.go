package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// DisplayBox contains information relevant to the player like
// their score, remaining lives, etc.
type DisplayBox struct {
	pos    Position
	height int
	width  int
	score  *TextBox
}

func (d *DisplayBox) Draw(scrn tcell.Screen) {
	d.score.Draw(scrn)
}

func (d *DisplayBox) SetPosition(pos Position) {
	d.pos = pos
}

func (d *DisplayBox) SetScore(score uint) {
	d.score.SetText(fmt.Sprintf(scoreFormat, score))
}

func (d *DisplayBox) Height() int {
	return max(d.height, d.score.Height())
}

func (d *DisplayBox) Width() int {
	return d.width
}

func (d *DisplayBox) Bottom() int {
	return d.pos.y + d.Height()
}

func NewDisplayBox(pos Position, height, width int) *DisplayBox {
	scoreBox := NewTextBox(fmt.Sprintf(scoreFormat, 0), boardStyle).
		SetHeight(height / 2).SetPosition(pos).SetWidth(width).NoBorder()

	return &DisplayBox{
		pos:    pos,
		height: height,
		width:  width,
		score:  scoreBox,
	}
}
