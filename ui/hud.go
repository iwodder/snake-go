package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

const (
	livesFormat = "Lives: %d"
	scoreFormat = "Score: %d"
	title       = "Snake"
)

var (
	boardStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
)

// Hud contains information relevant to the player like
// their score, remaining lives, etc.
type Hud struct {
	composite
	pos    Position
	height int
	width  int
	title  *TextBox
	score  *TextBox
	lives  *TextBox
}

func (d *Hud) SetPosition(pos Position) {
	d.pos = pos
}

func (d *Hud) Height() int {
	return max(d.height, d.score.Height()+d.lives.Height()+d.title.Height())
}

func (d *Hud) Width() int {
	return d.width
}

func (d *Hud) Bottom() int {
	return d.pos.Y + d.Height()
}

func (d *Hud) SetTitleBox(title *TextBox) {
	if d.title != nil {
		_ = d.Remove(d.title)
	}
	_ = d.Add(title)
	d.title = title
}

func (d *Hud) SetScoreBox(score *TextBox) {
	if d.score != nil {
		_ = d.Remove(d.score)
	}
	_ = d.Add(score)
	d.score = score
}

func (d *Hud) SetLivesBox(lives *TextBox) {
	if d.lives != nil {
		_ = d.Remove(d.lives)
	}
	_ = d.Add(lives)
	d.lives = lives
}

func (d *Hud) TitleBox() *TextBox {
	return d.title
}

func (d *Hud) ScoreBox() *TextBox {
	return d.score
}

func (d *Hud) LivesBox() *TextBox {
	return d.lives
}

func NewHud(pos Position, height, width int) *Hud {
	boxHeight := height / 3
	titleBox := NewTextBoxWithAlignment(title, CenterAlignment, boardStyle).
		SetHeight(boxHeight).SetPosition(Position{X: pos.X, Y: pos.Y}).
		SetWidth(width).NoBorder()
	scoreBox := NewTextBox(fmt.Sprintf(scoreFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(Position{X: pos.X, Y: titleBox.BottomEdge()}).
		SetWidth(width).NoBorder()
	livesBox := NewTextBox(fmt.Sprintf(livesFormat, 0), boardStyle).
		SetHeight(boxHeight).SetPosition(Position{X: pos.X, Y: scoreBox.BottomEdge()}).
		SetWidth(width).NoBorder()

	ret := Hud{
		composite: composite{},
		pos:       pos,
		height:    height,
		width:     width,
	}
	ret.SetTitleBox(titleBox)
	ret.SetScoreBox(scoreBox)
	ret.SetLivesBox(livesBox)

	return &ret
}
