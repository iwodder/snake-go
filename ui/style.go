package ui

import "github.com/gdamore/tcell/v2"

const (
	snakeStyle = "snake"
	foodStyle  = "food"
)

var styles = map[string]tcell.Style{
	snakeStyle: tcell.StyleDefault.Foreground(tcell.ColorGreen),
	foodStyle:  tcell.StyleDefault.Foreground(tcell.ColorRed),
}
