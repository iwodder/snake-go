package main

import "github.com/gdamore/tcell/v2"

type apple struct {
	pos pos
}

func (a apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, tcell.StyleDefault)
}
