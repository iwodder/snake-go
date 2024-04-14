package main

import "github.com/gdamore/tcell/v2"

type box struct {
	height, width int
	vecs          []vector
}

func (b *box) draw(scn tcell.Screen) {
	var r rune
	for y := range b.height {
		if y == 0 || y == b.height-1 {
			for x := range b.width {
				if y == 0 && x == 0 {
					r = tcell.RuneULCorner
				} else if y == 0 && x == b.width-1 {
					r = tcell.RuneURCorner
				} else if y == b.height-1 && x == 0 {
					r = tcell.RuneLLCorner
				} else if y == b.height-1 && x == b.width-1 {
					r = tcell.RuneLRCorner
				} else {
					r = tcell.RuneHLine
				}
				scn.SetContent(x, y, r, nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
			}
		} else {
			scn.SetContent(0, y, tcell.RuneVLine, nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
			scn.SetContent(b.width-1, y, tcell.RuneVLine, nil, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
		}
	}
}

func newBox(height, width int) *box {
	return &box{
		height: height, width: width,
	}
}
