package main

import (
	"github.com/gdamore/tcell/v2"
)

// TextBox displays immutable text on the screen. The TextBox is wrapped with a border
// and has no padding.
type TextBox struct {
	upperLeft Position
	text      string
	style     tcell.Style
}

func (p *TextBox) Draw(scrn tcell.Screen) {
	p.fill(scrn)
	p.drawBorder(scrn)
	p.drawText(scrn)
}

func (p *TextBox) Height() int {
	// top & bottom borders + padding
	return 3
}

func (p *TextBox) Width() int {
	return len(p.text) + 2
}

func (p *TextBox) SetPosition(pos Position) {
	p.upperLeft = pos
}

func (p *TextBox) Location() (x, y int) {
	return p.upperLeft.x, p.upperLeft.y
}

func (p *TextBox) fill(scrn tcell.Screen) {
	for y := p.topEdge(); y < p.bottomEdge(); y++ {
		for x := p.rightEdge(); x < p.leftEdge(); x++ {
			scrn.SetContent(x, y, ' ', nil, p.style)
		}
	}
}

func (p *TextBox) drawBorder(scrn tcell.Screen) {
	for x := p.upperLeft.x; x < p.upperLeft.x+p.Width(); x++ {
		scrn.SetContent(x, p.upperLeft.y, tcell.RuneHLine, nil, p.style)
		scrn.SetContent(x, p.upperLeft.y+p.Height()-1, tcell.RuneHLine, nil, p.style)
	}
	for y := p.upperLeft.y; y < p.upperLeft.y+p.Height(); y++ {
		scrn.SetContent(p.upperLeft.x, y, tcell.RuneVLine, nil, p.style)
		scrn.SetContent(p.upperLeft.x+p.Width()-1, y, tcell.RuneVLine, nil, p.style)
	}
	scrn.SetContent(p.upperLeft.x, p.upperLeft.y, tcell.RuneULCorner, nil, p.style)
	scrn.SetContent(p.upperLeft.x+p.Width()-1, p.upperLeft.y, tcell.RuneURCorner, nil, p.style)
	scrn.SetContent(p.upperLeft.x+p.Width()-1, p.upperLeft.y+p.Height()-1, tcell.RuneLRCorner, nil, p.style)
	scrn.SetContent(p.upperLeft.x, p.upperLeft.y+p.Height()-1, tcell.RuneLLCorner, nil, p.style)
}

func (p *TextBox) drawText(scrn tcell.Screen) {
	x, y := p.upperLeft.x+1, p.upperLeft.y+1
	for i, ch := range p.text {
		scrn.SetContent(x+i, y, ch, nil, p.style)
	}
}

func (p *TextBox) bottomEdge() int {
	return p.upperLeft.y + 2
}

func (p *TextBox) topEdge() int {
	return p.upperLeft.y
}

func (p *TextBox) leftEdge() int {
	return p.upperLeft.x
}

func (p *TextBox) rightEdge() int {
	return p.upperLeft.x + len(p.text) + 1
}

func NewTextBox(text string, style tcell.Style) *TextBox {
	return &TextBox{
		text:  text,
		style: style,
	}
}
