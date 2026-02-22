package ui

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

const snakeRune = 'X'

type SnakeComponent struct {
	leaf
	Body []Position
}

func (s *SnakeComponent) Draw(scrn tcell.Screen) {
	for _, c := range slices.All(s.Body) {
		scrn.SetContent(c.X, c.Y, snakeRune, nil, tcell.StyleDefault)
	}
}

func (s *SnakeComponent) Width() int {
	//TODO implement me
	panic("implement me")
}

func (s *SnakeComponent) Height() int {
	//TODO implement me
	panic("implement me")
}
