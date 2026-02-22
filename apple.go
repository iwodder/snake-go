package main

import (
	"math/rand"
	"snake/ui"
	"time"

	"github.com/gdamore/tcell/v2"
)

type apples []apple

func (as apples) draw(scn tcell.Screen) {
	for _, a := range as {
		a.draw(scn)
	}
}

func (as apples) Update(g *game, _ time.Duration) {
	for i := range as {
		as[i].Update(g)
	}
}

func newApples(b *ui.GameBoard, cnt int) apples {
	ret := make(apples, 0, cnt)
	for range cnt {
		ret = append(ret, newApple(b))
	}
	return ret
}

type apple struct {
	pos   ui.Position
	eaten bool
}

func (a *apple) draw(scn tcell.Screen) {
	scn.SetContent(a.pos.X, a.pos.Y, 'A', nil, appleStyle)
}

func (a *apple) Update(g *game) {
	if a.eaten {
		a.setPos(g.GameBoard)
		a.eaten = false
	}
}

func (a *apple) setPos(b *ui.GameBoard) {
	p := ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	for a.pos == p || !b.IsInside(p) {
		p = ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	}
	a.pos = p
}

func newApple(b *ui.GameBoard) apple {
	var ret apple
	ret.setPos(b)
	return ret
}
