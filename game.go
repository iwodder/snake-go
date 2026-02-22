package main

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"snake/ui"
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
const livesFormat = "Lives: %d"
const scoreFormat = "Score: %d"

type game struct {
	*ui.GameBoard
	eventMap       EventMap
	eventListeners EventListeners
	snake          *snake
	apples         apples
	score          uint
	remainingLives uint
	finished       bool
	paused         bool
	gameOver       bool
}

func (g *game) keyEventCallback(event *tcell.EventKey) {
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
	g.snake.move(g.GameBoard, delta)
	g.apples.move(g.GameBoard, delta)
	if g.snake.crashed() {
		g.remainingLives -= 1
		if g.remainingLives == 0 {
			g.gameOver = true
		} else {
			g.GameBoard.LivesBox().SetText(fmt.Sprintf(livesFormat, g.remainingLives))
			g.snake.ResetTo(g.GameBoard.Center())
		}
	} else {
		applesEaten := g.snake.eat(g.apples)
		g.score += applesEaten * pointsPerApple
		g.GameBoard.ScoreBox().SetText(fmt.Sprintf(scoreFormat, g.score))
	}
}

func (g *game) Draw(scrn tcell.Screen) {
	g.GameBoard.Draw(scrn)
	g.apples.draw(scrn)
	switch {
	case g.gameOver:
		ui.ShowMessage(g.GameBoard, GameOverText, scrn)
	case g.paused:
		ui.ShowMessage(g.GameBoard, GamePausedText, scrn)
	}
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
	g.snake = newSnake(g.GameBoard.Center())
	g.eventListeners = append(g.eventListeners, g.snake)
}

func newSnakeGame(cfg *Config, width int, height int) *game {
	b := ui.NewGameBoard(ui.Position{X: 0, Y: 0}, ui.Position{X: min(width, maxWidth), Y: min(height, maxHeight)})
	s := newSnakeOfLength(b.Center(), cfg.SnakeStartingLength())
	a := newApples(b, cfg.MaxNumberOfApples())

	_ = b.Add(s)

	ret := game{
		eventListeners: EventListeners{s},
		eventMap:       EventMap{},
		GameBoard:      b,
		snake:          s,
		apples:         a,
		remainingLives: cfg.NumberOfLives(),
	}
	b.LivesBox().SetText(fmt.Sprintf(livesFormat, ret.remainingLives))
	b.SetKeyEventCallback(ret.keyEventCallback)
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
