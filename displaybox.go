package main

import (
	"fmt"
	"snake/ui"

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
	title  *ui.TextBox
	score  *ui.TextBox
	lives  *ui.TextBox
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
	return d.pos.Y + d.Height()
}

func NewDisplayBox(pos Position, height, width int) *DisplayBox {
	boxHeight := height / 3
	titleBox := ui.NewTextBoxWithAlignment(title, ui.CenterAlignment, boardStyle).
		SetHeight(boxHeight).SetPosition(ui.Position{X: pos.X, Y: pos.Y}).
		SetWidth(width).NoBorder()
	scoreBox := ui.NewTextBox(fmt.Sprintf(scoreFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(ui.Position{X: pos.X, Y: titleBox.BottomEdge()}).
		SetWidth(width).NoBorder()
	livesBox := ui.NewTextBox(fmt.Sprintf(livesFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(ui.Position{X: pos.X, Y: scoreBox.BottomEdge()}).
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
