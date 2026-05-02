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

type game struct {
	*ui.Manager
	cfg            *Config
	gameBoard      *gameBoard
	score          uint
	remainingLives uint
	finished       bool
	currentState   state
}

func (g *game) keyHandler(key *tcell.EventKey) {
	switch event := eventMap.Get(key); event {
	case ExitGame:
		g.finished = true
	default:
		g.currentState.handle(g, event)
	}
}

func (g *game) Update(delta time.Duration) {
	g.currentState.update(g, delta)
}

func (g *game) Finished() bool {
	return g.finished
}

func (g *game) gameOver() bool {
	return g.remainingLives == 0
}

func (g *game) reset() {
	g.score = 0
	g.remainingLives = g.cfg.NumberOfLives()
	g.gameBoard.reset()
}

func newSnakeGame(cfg *Config, width int, height int) *game {
	b := newGameBoard(
		ui.Position{X: 0, Y: 0},
		ui.Position{X: min(width, maxWidth), Y: min(height, maxHeight)},
		cfg,
	)
	mgr := ui.NewManager()
	mgr.AddView("GameBoard", b)

	ret := game{
		Manager:        mgr,
		cfg:            cfg,
		gameBoard:      b,
		remainingLives: cfg.NumberOfLives(),
		currentState:   new(menuState),
	}
	mgr.SetKeyEventCallback(ret.keyHandler)
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
