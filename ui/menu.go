package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Menu struct {
	composite
	ul      Position
	width   int
	height  int
	title   *TextBox
	entries []*TextBox
}

func (m *Menu) Draw(scn tcell.Screen) {
	fill(m.ul, m.Width(), m.Height(), boardStyle, scn)
	drawBorder(m.ul, m.Width(), m.Height(), boardStyle, scn)
	m.composite.Draw(scn)
}

func (m *Menu) Width() int {
	return m.width
}

func (m *Menu) Height() int {
	ret := m.title.Height() + borderWidth*2
	for _, entry := range m.entries {
		ret += entry.Height()
	}
	return min(ret, m.height)
}

func (m *Menu) contentWidth() int {
	return m.Width() - borderWidth*2
}

func (m *Menu) AddEntry(text string) {
	pos := m.calculatePosOfNextEntry()
	entryBox := NewTextBoxWithAlignment(text, CenterAlignment, boardStyle).
		SetPosition(pos).
		SetWidth(m.contentWidth()).
		NoBorder()
	_ = m.Add(entryBox)
	m.entries = append(m.entries, entryBox)
}

func (m *Menu) calculatePosOfNextEntry() Position {
	var pos Position
	if len(m.entries) == 0 {
		pos.X, pos.Y = m.title.Position()
		pos.Y += m.title.Height()
	} else {
		lastEntry := m.entries[len(m.entries)-1]
		pos.X, pos.Y = lastEntry.Position()
		pos.Y += lastEntry.Height()
	}
	return pos
}

func NewMenu(ul Position, width, height int, title string) *Menu {
	maxEntries := height - 1
	ret := Menu{
		composite: composite{},
		ul:        ul,
		width:     width,
		height:    height,
		entries:   make([]*TextBox, 0, maxEntries),
	}

	titleBox := NewTextBoxWithAlignment(title, CenterAlignment, boardStyle).
		NoBorder().
		SetPosition(Position{X: ul.X + borderWidth, Y: ul.Y + borderWidth}).
		SetWidth(ret.contentWidth())
	_ = ret.Add(titleBox)
	ret.title = titleBox
	return &ret
}
