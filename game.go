package main

import (
	"github.com/gdamore/tcell/v2"
	"os"
	"time"
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
	scn      tcell.Screen
	events   chan *tcell.EventKey
	board    *board
	snake    *snake
	apples   apples
	exitFunc func(int)
}

func (g *game) start() {
	go g.eventPoller()

	now := time.Now()
	var delta time.Duration
	for {
		next := time.Now()
		delta = next.Sub(now)
		now = next

		select {
		case ev := <-g.events:
			g.kl.post(ev)
		default:
			g.snake.eat(g.apples)

			g.snake.move(g.board.boundary(), delta)
			g.apples.move(g.board.boundary(), delta)

			g.board.draw(g.scn)
			g.snake.draw(g.scn)
			g.apples.draw(g.scn)
			g.scn.Show()

			time.Sleep(time.Second / 60)
		}
	}
}

func (g *game) eventPoller() {
	g.events = make(chan *tcell.EventKey)
	for {
		switch ev := g.scn.PollEvent().(type) {
		case *tcell.EventKey:
			g.events <- ev
		case nil: // screen finalized
			return
		}
	}
}

func (g *game) registerKeyListener(kl keyListener) {
	g.kl = append(g.kl, kl)
}

func (g *game) notify(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyCtrlC, tcell.KeyEscape:
		g.scn.Fini()
		g.exitFunc(0)
	}
}

func newGame(scn tcell.Screen) *game {
	ret := game{
		kl:       keyListeners{},
		scn:      scn,
		events:   make(chan *tcell.EventKey),
		exitFunc: os.Exit,
	}

	x, y := ret.scn.Size()
	ret.board = newBoard(pos{x: 0, y: 0}, pos{x: min(x, maxWidth), y: min(y, maxHeight)})

	ret.snake = newSnake(ret.board.boundary().center())
	ret.apples = newApples(ret.board.boundary(), 2)

	ret.registerKeyListener(ret.snake)
	ret.registerKeyListener(&ret)

	return &ret
}
