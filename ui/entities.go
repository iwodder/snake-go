package ui

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

const (
	appleRune = 'A'
	snakeRune = 'X'
)

type SnakeRenderer struct {
	leaf
	Body []Position
}

func (s *SnakeRenderer) Draw(scrn tcell.Screen) {
	for _, c := range slices.All(s.Body) {
		scrn.SetContent(c.X, c.Y, snakeRune, nil, tcell.StyleDefault)
	}
}

func (s *SnakeRenderer) Width() int {
	//TODO implement me
	panic("implement me")
}

func (s *SnakeRenderer) Height() int {
	//TODO implement me
	panic("implement me")
}

type AppleRenderer struct {
	leaf
	Pos Position
}

func (a *AppleRenderer) Draw(scn tcell.Screen) {
	scn.SetContent(a.Pos.X, a.Pos.Y, appleRune, nil, tcell.StyleDefault)
}

func (a *AppleRenderer) Width() int {
	return 1
}

func (a *AppleRenderer) Height() int {
	return 1
}
