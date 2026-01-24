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

func Test_Styles(t *testing.T) {
	require.Equal(t, appleStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	require.Equal(t, boardStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	require.Equal(t, snakeStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
}

func Test_NewGameState(t *testing.T) {
	scn := setupDefaultScreen(t)
	width, height := scn.Size()
	g := newSnakeGame(&Config{}, width, height)

	require.Len(t, g.eventListeners, 1, "snake should be registered for key events")
	require.NotNil(t, g.snake)
	require.NotNil(t, g.apples)
	require.NotNil(t, g.board)
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
	var b *ui.GameBoard
	var a apples
	var s *snake
	var g game

	setup := func() {
		b = ui.NewGameBoard(Position{X: 0, Y: 0}, Position{X: 9, Y: 9})
		pos := b.Center()
		a = apples{apple{pos: Position{X: pos.X + 1, Y: pos.Y}}}
		s = newSnake(b.Center())
		g = game{
			board:          b,
			snake:          s,
			apples:         a,
			remainingLives: DefaultNumberOfLives,
		}
	}

	t.Run("player earns points for eating apples", func(t *testing.T) {
		setup()

		g.Update(moveDelta)

		require.Equal(t, pointsPerApple, g.score)
	})

	t.Run("crashing reduces remainingLives remaining", func(t *testing.T) {
		setup()

		s.body = []cell{
			{x: 3, y: 3},
			{x: 4, y: 3},
			{x: 4, y: 2},
			{x: 3, y: 2},
			{x: 3, y: 3},
		}

		g.Update(moveDelta)

		require.False(t, g.gameOver)
		require.Equal(t, uint(2), g.remainingLives)
	})

	t.Run("crashing and running out of remaining lives snake ends game", func(t *testing.T) {
		setup()
		g.remainingLives = 1
		s.body = []cell{
			{x: 3, y: 3},
			{x: 4, y: 3},
			{x: 4, y: 2},
			{x: 3, y: 2},
			{x: 3, y: 3},
		}

		g.Update(moveDelta)

		require.True(t, g.gameOver)
	})

	t.Run("on game over no entities move", func(t *testing.T) {
		setup()
		g.gameOver = true
		startPos := s.headPos()

		g.Update(moveDelta)

		require.True(t, slices.IndexFunc(a, func(a apple) bool { return a.eaten }) == -1)
		require.Equal(t, startPos, s.headPos())
	})

	t.Run("on game over enter resets game", func(t *testing.T) {
		setup()
		g.gameOver = true
		g.score = 100

		g.Handle(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone))

		assert.False(t, g.gameOver)
		assert.Zero(t, g.score)
		assert.Equal(t, DefaultNumberOfLives, g.remainingLives)
		assert.NotSame(t, g.snake, s)
	})

	t.Run("on game over pressing pause key does nothing", func(t *testing.T) {
		setup()
		g.gameOver = true
		paused := g.paused

		g.Handle(tcell.NewEventKey(tcell.KeyRune, ' ', tcell.ModNone))

		assert.Equal(t, paused, g.paused)
	})
}

type spyScreen struct {
	wasFinialized bool
	tcell.SimulationScreen
}

func (s *spyScreen) Fini() {
	s.wasFinialized = true
	s.SimulationScreen.Fini()
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
