package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

// maxWidth and maxHeight are zero-based numbers
const maxWidth = 39
const maxHeight = maxWidth
const pointsPerApple uint = 100

var (
	appleStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	boardStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	snakeStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
)

type game struct {
	eventMap       EventMap
	eventListeners EventListeners
	board          *board
	snake          *snake
	apples         apples
	score          uint
	finished       bool
	paused         bool
}

func (g *game) Handle(event tcell.Event) {
	ev := g.eventMap.Get(event)
	switch ev {
	case ExitGame:
		g.finished = true
	case PauseGame:
		g.paused = !g.paused
	default:
		g.eventListeners.Notify(ev)
	}
}

func (g *game) Update(delta time.Duration) {
	if g.paused {
		return
	}
	applesEaten := g.snake.eat(g.apples)
	g.score += applesEaten * pointsPerApple
	g.board.setScore(g.score)
	g.snake.move(g.board, delta)
	g.apples.move(g.board, delta)
}

func (g *game) Draw(scrn tcell.Screen) {
	g.board.draw(scrn)
	g.snake.draw(scrn)
	g.apples.draw(scrn)
	if g.paused {
		g.drawPausedBox(scrn)
	}
}

func (g *game) drawPausedBox(scrn tcell.Screen) {
	paused := NewTextBox("Game Paused", boardStyle)
	paused.SetPosition(Position{
		x: (g.board.width() - paused.Width()) / 2,
		y: (g.board.height() - paused.Height()) / 2,
	})
	paused.Draw(scrn)
}

func (g *game) Finished() bool {
	return g.finished
}

func newGame(scn tcell.Screen) *game {
	x, y := scn.Size()
	b := newBoard(Position{x: 0, y: 0}, Position{x: min(x, maxWidth), y: min(y, maxHeight)})
	s := newSnake(b.center())
	a := newApples(b, 2)

	return &game{
		eventListeners: EventListeners{s},
		eventMap:       EventMap{},
		board:          b,
		snake:          s,
		apples:         a,
	}
}

type Game interface {
	Handle(event tcell.Event)
	Update(delta time.Duration)
	Draw(scrn tcell.Screen)
	Finished() bool
}

func RunGame(game Game, scrn tcell.Screen) (err error) {
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
		game.Update(delta)
		game.Draw(scrn)
		scrn.Show()
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
