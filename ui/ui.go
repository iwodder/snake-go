package ui

import "github.com/gdamore/tcell/v2"

// drawBorder renders a rectangular border onto the screen using the given position, dimensions, and style.
func drawBorder(start Position, width int, height int, style tcell.Style, scrn tcell.Screen) {
	if width <= 0 {
		panic("width must be greater than zero")
	}
	if height <= 0 {
		panic("height must be greater than zero")
	}

	for x := start.X; x < start.X+width; x++ {
		scrn.SetContent(x, start.Y, tcell.RuneHLine, nil, style)
		scrn.SetContent(x, start.Y+height-1, tcell.RuneHLine, nil, style)
	}
	for y := start.Y; y < start.Y+height; y++ {
		scrn.SetContent(start.X, y, tcell.RuneVLine, nil, style)
		scrn.SetContent(start.X+width-1, y, tcell.RuneVLine, nil, style)
	}
	scrn.SetContent(start.X, start.Y, tcell.RuneULCorner, nil, style)
	scrn.SetContent(start.X+width-1, start.Y, tcell.RuneURCorner, nil, style)
	scrn.SetContent(start.X+width-1, start.Y+height-1, tcell.RuneLRCorner, nil, style)
	scrn.SetContent(start.X, start.Y+height-1, tcell.RuneLLCorner, nil, style)
}
