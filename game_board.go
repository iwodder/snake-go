package main

import (
	"fmt"
	"snake/ui"
	"time"

	"github.com/gdamore/tcell/v2"
)

const livesFormat = "Lives: %d"
const scoreFormat = "Score: %d"

type gameBoard struct {
	*ui.GameBoardRenderer
	snake  *snake
	apples apples
}

func (b *gameBoard) Update(g *game, delta time.Duration) {
	b.snake.Update(b, g, delta)
	b.apples.Update(b, delta)
	b.GameBoardRenderer.LivesBox().SetText(fmt.Sprintf(livesFormat, g.remainingLives))
	b.GameBoardRenderer.ScoreBox().SetText(fmt.Sprintf(scoreFormat, g.score))
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

func (b *gameBoard) keyHandler(key *tcell.EventKey) {
	b.snake.Notify(eventMap.GetEventFromKey(key))
}

func (b *gameBoard) reset() {
	b.snake.ResetTo(b.Center())
}

func newGameBoard(ul, lr ui.Position, cfg *Config) *gameBoard {
	ret := gameBoard{
		GameBoardRenderer: ui.NewGameBoardRenderer(ul, lr),
	}
	ret.SetKeyEventCallback(ret.keyHandler)
	ret.LivesBox().SetText(fmt.Sprintf(livesFormat, cfg.NumberOfLives()))

	s := newSnakeOfLength(ret.Center(), cfg.SnakeStartingLength())
	ret.snake = s
	a := newApples(&ret, cfg.MaxNumberOfApples())
	ret.apples = a

	_ = ret.Add(s)
	a.ForEach(func(a *apple) {
		_ = ret.Add(a)
	})

	return &ret
}
