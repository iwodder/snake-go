package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type apples []apple

func (as apples) draw(scn tcell.Screen) {
	for _, a := range as {
		a.draw(scn)
	}
}

func (as apples) move(b *board, _ time.Duration) {
	for i := range as {
		as[i].move(b)
	}
}

func newApples(b *board, cnt int) apples {
	ret := make(apples, 0, cnt)
	for range cnt {
		ret = append(ret, newApple(b))
	}
	return ret
}

type apple struct {
	pos   Position
	eaten bool
}

func (a *apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.x, a.pos.y, 'A', nil, appleStyle)
}

func (a *apple) move(b *board) {
	if a.eaten {
		a.setPos(b)
		a.eaten = false
	}
}

func (a *apple) setPos(b *board) {
	p := Position{x: rand.Intn(b.rightEdge()), y: rand.Intn(b.bottomEdge())}
	for a.pos == p || !b.isInside(p) {
		p = Position{x: rand.Intn(b.lowerRight.x), y: rand.Intn(b.lowerRight.y)}
	}
	a.pos = p
}

func newApple(b *board) apple {
	var ret apple
	ret.setPos(b)
	return ret
}
