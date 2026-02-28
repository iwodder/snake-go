package main

import (
	"math/rand"
	"snake/ui"
	"time"
)

type apples []apple

func (a apples) Update(g *game, _ time.Duration) {
	for i := range a {
		a[i].Update(g)
	}
}

func (a apples) ForEach(f func(*apple)) {
	for i := range a {
		f(&a[i])
	}
}

func newApples(b *ui.GameBoard, cnt int) apples {
	ret := make([]apple, 0, cnt)
	for range cnt {
		ret = append(ret, newApple(b))
	}
	return ret
}

type apple struct {
	ui.AppleRenderer
	eaten bool
}

func (a *apple) Update(g *game) {
	if a.eaten {
		a.setPos(g.GameBoard)
		a.eaten = false
	}
}

func (a *apple) setPos(b *ui.GameBoard) {
	p := ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	for a.Pos == p || !b.IsInside(p) {
		p = ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	}
	a.Pos = p
}

func newApple(b *ui.GameBoard) apple {
	var ret apple
	ret.setPos(b)
	return ret
}
