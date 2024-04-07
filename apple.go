package main

import "github.com/gdamore/tcell/v2"

type apples []apple

func (as apples) draw(scn tcell.Screen) {
	for _, a := range as {
		a.draw(scn)
	}
}

type apple struct {
	pos   pos
	eaten bool
}

func (a apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, tcell.StyleDefault)
}
