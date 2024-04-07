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

func (as apples) move() {
	for i := range as {
		as[i].move()
	}
}

type apple struct {
	pos   pos
	eaten bool
}

func (a *apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, tcell.StyleDefault)
}

func (a *apple) move() {
	if a.eaten {
		p := pos{x: rand.Int(), y: rand.Int()}
		for i := 0; a.pos == p; i++ {
			if i%2 == 0 {
				p.x = rand.Int()
			} else {
				p.y = rand.Int()
			}
		}
		a.pos = p
		a.eaten = false
	}
}
