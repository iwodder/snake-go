package main

import (
	"github.com/gdamore/tcell/v2"
	"math/rand"
	"time"
)

type apples []apple

func (as apples) draw(scn tcell.Screen) {
	for _, a := range as {
		a.draw(scn)
	}
}

func (as apples) move(b boundary, _ time.Duration) {
	for i := range as {
		as[i].move(b)
	}
}

func newApples(b boundary, cnt int) apples {
	ret := make(apples, 0, cnt)
	for range cnt {
		ret = append(ret, newApple(b))
	}
	return ret
}

type apple struct {
	pos   pos
	eaten bool
}

func (a *apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, appleStyle)
}

func (a *apple) move(b boundary) {
	if a.eaten {
		a.setPos(b)
		a.eaten = false
	}
}

func (a *apple) setPos(b boundary) {
	p := pos{x: rand.Intn(b.lowerRight.x), y: rand.Intn(b.lowerRight.y)}
	for a.pos == p || !b.isInside(p) {
		p = pos{x: rand.Intn(b.lowerRight.x), y: rand.Intn(b.lowerRight.y)}
	}
	a.pos = p
}

func newApple(b boundary) apple {
	var ret apple
	ret.setPos(b)
	return ret
}
