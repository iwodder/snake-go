package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
)

func Test_DefaultEvent(t *testing.T) {
	require.Equal(t, Event(0), Unknown, "zero value must not map to expected event(s)")
}

func Test_EventMappings(t *testing.T) {
	eventMap := EventMap{}

	t.Run("up event", func(t *testing.T) {
		require.Equal(t, MoveUp, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'w', tcell.ModNone)))
		require.Equal(t, MoveUp, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'W', tcell.ModNone)))
		require.Equal(t, MoveUp, eventMap.Get(tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)))
	})

	t.Run("down event", func(t *testing.T) {
		require.Equal(t, MoveDown, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 's', tcell.ModNone)))
		require.Equal(t, MoveDown, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'S', tcell.ModNone)))
		require.Equal(t, MoveDown, eventMap.Get(tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)))
	})

	t.Run("right event", func(t *testing.T) {
		require.Equal(t, MoveRight, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'd', tcell.ModNone)))
		require.Equal(t, MoveRight, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'd', tcell.ModNone)))
		require.Equal(t, MoveRight, eventMap.Get(tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone)))
	})

	t.Run("left event", func(t *testing.T) {
		require.Equal(t, MoveLeft, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone)))
		require.Equal(t, MoveLeft, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, 'A', tcell.ModNone)))
		require.Equal(t, MoveLeft, eventMap.Get(tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone)))
	})

	t.Run("pause event", func(t *testing.T) {
		require.Equal(t, PauseGame, eventMap.Get(tcell.NewEventKey(tcell.KeyRune, ' ', tcell.ModNone)))
	})

	t.Run("exit event", func(t *testing.T) {
		require.Equal(t, ExitGame, eventMap.Get(tcell.NewEventKey(tcell.KeyCtrlC, 0, tcell.ModNone)))
	})
}
