package main

import (
	"context"
	"errors"
	"fmt"
	"slices"
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

const GameOverText = "Game Over"
const GamePausedText = "Game Paused"

type game struct {
	eventMap       EventMap
	eventListeners EventListeners
	board          *board
	snake          *snake
	apples         apples
	score          uint
	remainingLives uint
	finished       bool
	paused         bool
	gameOver       bool
}

func (g *game) Handle(event tcell.Event) {
	ev := g.eventMap.Get(event)
	switch ev {
	case ExitGame:
		g.finished = true
	case PauseGame:
		if !g.gameOver {
			g.paused = !g.paused
		}
	case StartGame:
		if g.gameOver {
			g.reset()
		}
	default:
		g.eventListeners.Notify(ev)
	}
}

func (g *game) Update(delta time.Duration) {
	if g.paused || g.gameOver {
		return
	}
	g.snake.move(g.board, delta)
	g.apples.move(g.board, delta)
	if g.snake.crashed() {
		g.remainingLives -= 1
		if g.remainingLives == 0 {
			g.gameOver = true
		} else {
			g.board.setLives(g.remainingLives)
			g.snake.ResetTo(g.board.center())
		}
	} else {
		applesEaten := g.snake.eat(g.apples)
		g.score += applesEaten * pointsPerApple
		g.board.setScore(g.score)
	}
}

func (g *game) Draw(scrn tcell.Screen) {
	g.board.draw(scrn)
	g.snake.draw(scrn)
	g.apples.draw(scrn)
	if g.paused {
		g.drawTextBox(GamePausedText, scrn)
	} else if g.gameOver {
		g.drawTextBox(GameOverText, scrn)
	}
}

func (g *game) drawTextBox(text string, scrn tcell.Screen) {
	textBox := NewTextBox(text, boardStyle)
	textBox.SetPosition(Position{
		x: (g.board.width() - textBox.Width()) / 2,
		y: (g.board.height() - textBox.Height()) / 2,
	})
	textBox.Draw(scrn)
}

func (g *game) Finished() bool {
	return g.finished
}

func (g *game) reset() {
	g.score = 0
	g.remainingLives = DefaultNumberOfLives
	g.gameOver = false
	g.eventListeners = slices.DeleteFunc(g.eventListeners, func(listener EventListener) bool {
		return listener == g.snake
	})
	g.snake = newSnake(g.board.center())
	g.eventListeners = append(g.eventListeners, g.snake)
}

func newSnakeGame(cfg *Config, scn tcell.Screen) *game {
	width, height := scn.Size()
	b := newBoard(Position{x: 0, y: 0}, Position{x: min(width, maxWidth), y: min(height, maxHeight)})
	s := newSnakeOfLength(b.center(), cfg.SnakeStartingLength())
	a := newApples(b, cfg.MaxNumberOfApples())

	ret := game{
		eventListeners: EventListeners{s},
		eventMap:       EventMap{},
		board:          b,
		snake:          s,
		apples:         a,
		remainingLives: cfg.NumberOfLives(),
	}
	ret.board.setLives(ret.remainingLives)
	return &ret
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
