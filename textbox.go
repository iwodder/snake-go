package main

import (
	"github.com/gdamore/tcell/v2"
)

const (
	MinTextboxHeightWithBorder = 3
	MinTextboxHeightNoBorder   = 1
)

// TextBox displays immutable text on the screen. The TextBox is wrapped with a border
// and has no padding.
type TextBox struct {
	upperLeft Position
	text      string
	style     tcell.Style
	height    int
	width     int
	border    bool
}

func (p *TextBox) Draw(scrn tcell.Screen) {
	p.fill(scrn)
	if p.border {
		p.drawBorder(scrn)
	}
	p.drawText(scrn)
}

func (p *TextBox) Height() int {
	if p.border {
		return max(MinTextboxHeightWithBorder, p.height)
	}
	return max(MinTextboxHeightNoBorder, p.height)
}

func (p *TextBox) Width() int {
	minWidth := len(p.text)
	if p.border {
		minWidth += 2
	}
	return max(minWidth, p.width)
}

func (p *TextBox) SetHeight(height int) *TextBox {
	p.height = height
	return p
}

func (p *TextBox) SetWidth(width int) *TextBox {
	p.width = width
	return p
}

func (p *TextBox) Position() (x, y int) {
	return p.upperLeft.x, p.upperLeft.y
}

func (p *TextBox) SetPosition(pos Position) *TextBox {
	p.upperLeft = pos
	return p
}

func (p *TextBox) Text() string {
	return p.text
}

func (p *TextBox) SetText(text string) *TextBox {
	p.text = text
	return p
}

func (p *TextBox) NoBorder() *TextBox {
	p.border = false
	return p
}

func (p *TextBox) fill(scrn tcell.Screen) {
	for y := p.topEdge(); y < p.BottomEdge(); y++ {
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
	x, y := p.getTextPos()
	for i, ch := range p.text {
		scrn.SetContent(x+i, y, ch, nil, p.style)
	}
}

func (p *TextBox) getTextPos() (x, y int) {
	if p.border {
		return p.upperLeft.x + 1, p.upperLeft.y + 1
	}
	return p.upperLeft.x, p.upperLeft.y
}

func (p *TextBox) BottomEdge() int {
	return p.upperLeft.y + p.Height() - 1
}

func (p *TextBox) topEdge() int {
	return p.upperLeft.y
}

func (p *TextBox) leftEdge() int {
	return p.upperLeft.x
}

func (p *TextBox) rightEdge() int {
	return p.upperLeft.x + p.Width() - 1
}

func NewTextBox(text string, style tcell.Style) *TextBox {
	return &TextBox{
		text:   text,
		style:  style,
		border: true,
	}
}
