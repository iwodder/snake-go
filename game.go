package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

const maxWidth = 40
const maxHeight = maxWidth

var (
	appleStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	boardStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	snakeStyle = tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
)

type keyListener interface {
	notify(event *tcell.EventKey)
}

type keyListeners []keyListener

func (k keyListeners) post(event *tcell.EventKey) {
	for _, listener := range k {
		listener.notify(event)
	}
}

type game struct {
	kl       keyListeners
	events   chan *tcell.EventKey
	board    *board
	snake    *snake
	apples   apples
	finished bool
}

func (g *game) Handle(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyCtrlC {
			g.finished = true
			return
		}
		g.kl.post(ev)
	}
}

func (g *game) Update(delta time.Duration) {
	g.snake.eat(g.apples)
	g.snake.move(g.board, delta)
	g.apples.move(g.board, delta)
}

func (g *game) Draw(scrn tcell.Screen) {
	g.board.draw(scrn)
	g.snake.draw(scrn)
	g.apples.draw(scrn)
}

func (g *game) Finished() bool {
	return g.finished
}

func (g *game) registerKeyListener(kl keyListener) {
	g.kl = append(g.kl, kl)
}

func newGame(scn tcell.Screen) *game {
	x, y := scn.Size()
	b := newBoard(Position{x: 0, y: 0}, Position{x: min(x, maxWidth), y: min(y, maxHeight)})
	s := newSnake(b.center())
	a := newApples(b, 2)

	return &game{
		kl:     keyListeners{s},
		events: make(chan *tcell.EventKey, 1),
		board:  b,
		snake:  s,
		apples: a,
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
