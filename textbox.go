package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

const (
	MinTextboxHeightWithBorder = 3
	MinTextboxHeightNoBorder   = 1
)

const (
	NoAlignment TextAlignment = iota + 0
	LeftAlignment
	RightAlignment
	CenterAlignment
)

// TextAlignment is used to for adjusting how text is aligned within a TextBox
type TextAlignment int

func (t TextAlignment) Align(width int, text string) string {
	padding := width - len(text)
	if padding < 0 {
		return text
	}

	switch t {
	case NoAlignment:
		return text
	case LeftAlignment:
		return text + strings.Repeat(" ", padding)
	case RightAlignment:
		return strings.Repeat(" ", padding) + text
	case CenterAlignment:
		left := strings.Repeat(" ", padding/2)
		right := strings.Repeat(" ", padding-len(left))
		return left + text + right
	default:
		panic("unknown alignment value used")
	}
}

// TextBox displays immutable text on the screen. The TextBox is wrapped with a border
// and has no padding.
type TextBox struct {
	upperLeft Position
	alignment TextAlignment
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
	return p.alignment.Align(p.width, p.text)
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
	for i, ch := range p.Text() {
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
	return p.upperLeft.y + p.Height()
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
	return NewTextBoxWithAlignment(text, NoAlignment, style)
}

func NewTextBoxWithAlignment(text string, ta TextAlignment, style tcell.Style) *TextBox {
	return &TextBox{
		text:      text,
		style:     style,
		alignment: ta,
		border:    true,
	}
}
