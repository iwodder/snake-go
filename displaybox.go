package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
	livesFormat = "Lives: %d"
	scoreFormat = "Score: %d"
	title       = "Snake"
)

// DisplayBox contains information relevant to the player like
// their score, remaining lives, etc.
type DisplayBox struct {
	pos    Position
	height int
	width  int
	title  *TextBox
	score  *TextBox
	lives  *TextBox
}

func (d *DisplayBox) Draw(scrn tcell.Screen) {
	d.title.Draw(scrn)
	d.score.Draw(scrn)
	d.lives.Draw(scrn)
}

func (d *DisplayBox) SetPosition(pos Position) {
	d.pos = pos
}

func (d *DisplayBox) SetScore(score uint) {
	d.score.SetText(fmt.Sprintf(scoreFormat, score))
}

func (d *DisplayBox) SetLives(lives uint) {
	d.lives.SetText(fmt.Sprintf(livesFormat, lives))
}

func (d *DisplayBox) Height() int {
	return max(d.height, d.score.Height()+d.lives.Height()+d.title.Height())
}

func (d *DisplayBox) Width() int {
	return d.width
}

func (d *DisplayBox) Bottom() int {
	return d.pos.y + d.Height()
}

func NewDisplayBox(pos Position, height, width int) *DisplayBox {
	boxHeight := height / 3
	titleBox := NewTextBoxWithAlignment(title, CenterAlignment, boardStyle).
		SetHeight(boxHeight).SetPosition(pos).
		SetWidth(width).NoBorder()
	scoreBox := NewTextBox(fmt.Sprintf(scoreFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(Position{x: pos.x, y: titleBox.BottomEdge()}).
		SetWidth(width).NoBorder()
	livesBox := NewTextBox(fmt.Sprintf(livesFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(Position{x: pos.x, y: scoreBox.BottomEdge()}).
		SetWidth(width).NoBorder()

	return &DisplayBox{
		pos:    pos,
		height: height,
		width:  width,
		title:  titleBox,
		score:  scoreBox,
		lives:  livesBox,
	}
}
