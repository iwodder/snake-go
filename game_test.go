package main

import (
	"slices"
	"snake/ui"
	"sync"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewGameState(t *testing.T) {
	g := newSnakeGame(&Config{}, 10, 10)

	require.NotNil(t, g.gameBoard)
}

func Test_RunGame(t *testing.T) {
	simScreen := setupScreen(t, 20, 20)
	game := spyGame{}

	t.Run("executes game lifecycle", func(t *testing.T) {
		go func() {
			require.NoError(t, simScreen.PostEvent(tcell.NewEventKey(tcell.KeyUp, tcell.RuneUArrow, tcell.ModNone)))
			require.NoError(t, simScreen.PostEvent(tcell.NewEventKey(tcell.KeyCtrlC, 'C', tcell.ModCtrl)))
		}()
		require.NoError(t, RunGame(&game, simScreen))
		game.assertNotified(t)
		game.assertUpdated(t)
		game.assertDrawn(t)
	})

	t.Run("game runs until finished", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-time.After(time.Millisecond * 500)
			require.NoError(t, simScreen.PostEvent(tcell.NewEventKey(tcell.KeyCtrlC, 'C', tcell.ModCtrl)))
		}()
		require.NoError(t, RunGame(&game, simScreen))
		wg.Wait()
	})
}

func Test_Game(t *testing.T) {
	var b *gameBoard
	var a apples
	var s *snake
	var g game

	setup := func() {
		b = &gameBoard{
			GameBoardRenderer: ui.NewGameBoardRenderer(ui.Position{X: 0, Y: 0}, ui.Position{X: 9, Y: 9}),
		}
		pos := b.Center()
		a = apples{
			{AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: pos.X + 1, Y: pos.Y}}, eaten: false},
		}
		s = newSnake(b.Center())
		b.snake = s
		b.apples = a
		g = game{
			Manager:        ui.NewManager(),
			gameBoard:      b,
			cfg:            &Config{},
			remainingLives: DefaultNumberOfLives,
			currentState:   new(menuState),
		}
	}

	simulateEvent := func(g *game, event Event) {
		g.currentState.handle(g, event)
	}

	t.Run("player earns points for eating apples", func(t *testing.T) {
		setup()
		simulateEvent(&g, StartGame)

		g.Update(moveDelta)

		require.Equal(t, pointsPerApple, g.score)
	})

	t.Run("crashing reduces remainingLives remaining", func(t *testing.T) {
		setup()
		simulateEvent(&g, StartGame)

		simulate(g.gameBoard.snake, &g, MoveRight, MoveDown, MoveLeft, MoveUp)

		require.False(t, g.gameOver())
		require.Equal(t, uint(2), g.remainingLives)
	})

	t.Run("crashing and running out of remaining lives snake ends game", func(t *testing.T) {
		setup()
		g.remainingLives = 1

		simulate(g.gameBoard.snake, &g, MoveRight, MoveDown, MoveLeft, MoveUp)

		require.True(t, g.gameOver())
	})

	t.Run("on game over no entities move", func(t *testing.T) {
		setup()
		g.remainingLives = 0
		startPos := s.head()

		g.Update(moveDelta)

		require.True(t, slices.IndexFunc(a, func(a apple) bool { return a.eaten }) == -1)
		require.Equal(t, startPos, s.head())
	})

	t.Run("on game over enter resets game", func(t *testing.T) {
		setup()
		g.currentState = new(gameOverState)

		g.keyHandler(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone))

		assert.False(t, g.gameOver())
		assert.Zero(t, g.score)
		assert.Equal(t, DefaultNumberOfLives, g.remainingLives)
	})

	t.Run("on game over pressing pause key does nothing", func(t *testing.T) {
		setup()
		exp := g

		g.keyHandler(tcell.NewEventKey(tcell.KeyRune, ' ', tcell.ModNone))

		assert.Equal(t, exp, g)
	})
}

type spyGame struct {
	notified bool
	updated  bool
	drawn    bool
	finished bool
}

func (s *spyGame) Handle(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyCtrlC {
			s.finished = true
		}
	}
	s.notified = true
}

func (s *spyGame) Update(time.Duration) {
	s.updated = true
}

func (s *spyGame) Draw(tcell.Screen) {
	s.drawn = true
}

func (s *spyGame) Finished() bool {
	return s.finished
}

func (s *spyGame) assertNotified(t *testing.T) {
	assert.True(t, s.notified, "game was never notified of any events")
}

func (s *spyGame) assertUpdated(t *testing.T) {
	assert.True(t, s.updated, "game was never updated")
}

func (s *spyGame) assertDrawn(t *testing.T) {
	assert.True(t, s.drawn, "game was never drawn")
}

func setupScreen(t *testing.T, height, width int) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	ret.SetSize(height, width)
	return ret
}
