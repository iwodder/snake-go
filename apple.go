package main

import (
	"math/rand"
	"snake/ui"
	"time"
)

type apples []apple

func (a apples) Update(board *gameBoard, _ time.Duration) {
	for i := range a {
		a[i].Update(board)
	}
}

func (a apples) ForEach(f func(*apple)) {
	for i := range a {
		f(&a[i])
	}
}

func newApples(b *gameBoard, cnt int) apples {
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

func (a *apple) Update(board *gameBoard) {
	if a.eaten {
		a.setPos(board)
		a.eaten = false
	}
}

func (a *apple) setPos(b *gameBoard) {
	p := ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	for a.Pos == p || !b.IsInside(p) {
		p = ui.Position{X: rand.Intn(b.Right()), Y: rand.Intn(b.Bottom())}
	}
	a.Pos = p
}

func newApple(b *gameBoard) apple {
	var ret apple
	ret.setPos(b)
	return ret
}
