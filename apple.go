package main

import (
	"github.com/gdamore/tcell/v2"
	"math/rand"
)

type apples []apple

func (as apples) draw(scn tcell.Screen) {
	for _, a := range as {
		a.draw(scn)
	}
}

func (as apples) move(b bounds) {
	for i := range as {
		as[i].move(b)
	}
}

type apple struct {
	pos   pos
	eaten bool
}

func (a *apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, tcell.StyleDefault)
}

func (a *apple) move(b bounds) {
	if a.eaten {
		p := pos{x: rand.Intn(b.lowerRight.x), y: rand.Intn(b.lowerRight.y)}
		for a.pos == p || !b.isInside(p) {
			p = pos{x: rand.Intn(b.lowerRight.x), y: rand.Intn(b.lowerRight.y)}
		}
		a.pos = p
		a.eaten = false
	}
}