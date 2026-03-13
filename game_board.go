package main

import "snake/ui"

type gameBoard struct {
	*ui.GameBoardRenderer
}

func (b *gameBoard) Center() ui.Position {
	return ui.Position{
		X: b.Left() + (b.Right()-b.Left())/2,
		Y: b.Top() + (b.Bottom()-b.Top())/2,
	}
}

func (b *gameBoard) IsInside(pos ui.Position) bool {
	return pos.X > b.Left() && pos.X < b.Right() &&
		pos.Y > b.Top() && pos.Y < b.Bottom()
}

func newGameBoard(ul, lr ui.Position) *gameBoard {
	return &gameBoard{
		GameBoardRenderer: ui.NewGameBoardRenderer(ul, lr),
	}
}
