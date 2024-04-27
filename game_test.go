package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func Test_Styles(t *testing.T) {
	require.Equal(t, appleStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	require.Equal(t, boardStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	require.Equal(t, snakeStyle, tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
}

func Test_GameFinializesScreenOnCtrlC(t *testing.T) {
	didExit := new(bool)
	exitCode := new(int)
	spy := spyScreen{
		wasFinialized:    false,
		SimulationScreen: setupDefaultScreen(t),
	}

	g := newGame(&spy)
	g.exitFunc = func(val int) {
		*didExit = true
		*exitCode = val
	}

	g.notify(tcell.NewEventKey(tcell.KeyCtrlC, 0, 0))

	require.True(t, *didExit)
	require.True(t, spy.wasFinialized)
}

func Test_EventPollerLoopExitsOnNil(t *testing.T) {
	scn := setupDefaultScreen(t)
	g := newGame(scn)

	var wg sync.WaitGroup
	go func() {
		g.eventPoller()
		wg.Done()
	}()

	require.NoError(t, scn.PostEvent(tcell.NewEventKey(tcell.KeyCtrlC, 0, 0)))
	require.NoError(t, scn.PostEvent(nil))

	wg.Wait()
}

func Test_NewGameState(t *testing.T) {
	g := newGame(setupDefaultScreen(t))

	require.Len(t, g.kl, 2, "snake and game should be registered for key events")
	require.NotNil(t, g.events)
	require.NotNil(t, g.scn)
	require.NotNil(t, g.snake)
	require.NotNil(t, g.apples)
	require.NotNil(t, g.board)
	require.NotNil(t, g.exitFunc)
}

type spyScreen struct {
	wasFinialized bool
	tcell.SimulationScreen
}

func (s *spyScreen) Fini() {
	s.wasFinialized = true
	s.SimulationScreen.Fini()
}
