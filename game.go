package main

import (
	"context"
	"errors"
	"fmt"
	"snake/ui"
	"time"

	"github.com/gdamore/tcell/v2"
)

// maxWidth and maxHeight are zero-based numbers
const maxWidth = 39
const maxHeight = maxWidth
const pointsPerApple uint = 100

const GameOverText = "Game Over"
const GamePausedText = "Game Paused"
const livesFormat = "Lives: %d"
const scoreFormat = "Score: %d"

type game struct {
	*gameBoard
	score          uint
	remainingLives uint
	finished       bool
	paused         bool
}

func (g *game) Handle(event tcell.Event) {
	switch eventMap.Get(event) {
	case ExitGame:
		g.finished = true
	case PauseGame:
		if !g.gameOver() {
			g.paused = !g.paused
		}
	case StartGame:
		if g.gameOver() {
			g.reset()
		}
	default:
		g.gameBoard.Handle(event)
	}
}

func (g *game) Update(delta time.Duration) {
	if g.paused || g.gameOver() {
		return
	}
	g.gameBoard.Update(g, delta)
}

func (g *game) Draw(scrn tcell.Screen) {
	g.gameBoard.Draw(scrn)
	switch {
	case g.gameOver():
		ui.ShowMessage(g.gameBoard, GameOverText, scrn)
	case g.paused:
		ui.ShowMessage(g.gameBoard, GamePausedText, scrn)
	}
}

func (g *game) Finished() bool {
	return g.finished
}

func (g *game) gameOver() bool {
	return g.remainingLives == 0
}

func (g *game) reset() {
	g.score = 0
	g.remainingLives = DefaultNumberOfLives
	g.snake = newSnake(g.gameBoard.Center())
}

func newSnakeGame(cfg *Config, width int, height int) *game {
	b := newGameBoard(
		ui.Position{X: 0, Y: 0},
		ui.Position{X: min(width, maxWidth), Y: min(height, maxHeight)},
		cfg,
	)

	ret := game{
		gameBoard:      b,
		remainingLives: cfg.NumberOfLives(),
	}
	return &ret
}

type Game interface {
	ui.EventHandler
	Update(delta time.Duration)
	Draw(scrn tcell.Screen)
	Finished() bool
}

func RunGame(game Game, scrn tcell.Screen) (err error) {
	const FramesPerSecond = 60
	const FrameDuration = time.Second / FramesPerSecond

	ctx, cancel := context.WithCancel(context.Background())
	eventQueue := runEventPoller(ctx, scrn)
	defer func() {
		cancel()
		if r := recover(); r != nil {
			err = errors.Join(err, fmt.Errorf("panic caught in game loop: %v", r))
		}
	}()

	now := time.Now()
	var delta time.Duration
	for !game.Finished() {
		next := time.Now()
		delta = next.Sub(now)
		now = next

		select {
		case ev := <-eventQueue:
			game.Handle(ev)
		default:
		}
		scrn.Clear()
		game.Update(delta)
		game.Draw(scrn)
		scrn.Show()

		time.Sleep(FrameDuration - delta)
	}
	return nil
}

func runEventPoller(ctx context.Context, scrn tcell.Screen) <-chan tcell.Event {
	ret := make(chan tcell.Event, 1)
	go func() {
		scrn.ChannelEvents(ret, ctx.Done())
	}()
	return ret
}
