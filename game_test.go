package main

import (
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
	g := newGame(setupDefaultScreen(t))

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
	t.Run("player earns points for eating apples", func(t *testing.T) {
		b := board{
			upperLeft:  Position{x: 0, y: 0},
			lowerRight: Position{x: 9, y: 9},
			scoreBox:   NewTextBox("", boardStyle),
		}
		a := apples{apple{pos: Position{x: 4, y: 4}}}
		s := newSnake(Position{x: 4, y: 4})

		g := game{
			board:  &b,
			snake:  s,
			apples: a,
		}

		g.Update(moveDelta)

		require.Equal(t, pointsPerApple, g.score)
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
