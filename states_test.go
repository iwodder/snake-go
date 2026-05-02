package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_States(t *testing.T) {
	var g *game

	setup := func() {
		g = newSnakeGame(&Config{}, 10, 10)
	}

	t.Run("menu state", func(t *testing.T) {
		t.Run("transitions to playing on StartGameEvent", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)

			require.IsType(t, new(playingState), g.currentState)
		})

		t.Run("shows modal for menu text", func(t *testing.T) {
			setup()

			g.Update(time.Millisecond * 500)

			require.True(t, g.Manager.ModalVisible())
		})
	})

	t.Run("playing state", func(t *testing.T) {
		t.Run("transitions to pause on PauseEvent", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)
			expSavedState := g.currentState

			g.currentState.handle(g, PauseGame)

			require.IsType(t, new(pausedState), g.currentState)
			require.Equal(t, expSavedState, g.currentState.(*pausedState).currentGame)
		})

		t.Run("transitions to game over when game is over", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)
			g.remainingLives = 0

			g.Update(time.Millisecond * 500)

			require.IsType(t, new(gameOverState), g.currentState)
		})
	})

	t.Run("paused state", func(t *testing.T) {
		t.Run("transitions to playing on PauseEvent", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)
			expSavedState := g.currentState

			g.currentState.handle(g, PauseGame)
			g.currentState.handle(g, PauseGame)

			require.Equal(t, expSavedState, g.currentState)
		})

		t.Run("does nothing with other events", func(t *testing.T) {
			setup()
			g.currentState = new(pausedState)
			expSavedState := g.currentState

			g.currentState.handle(g, StartGame)
			g.currentState.handle(g, MoveRight)

			require.IsType(t, expSavedState, g.currentState)
		})

		t.Run("shows modal for paused text", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)
			g.currentState.handle(g, PauseGame)

			g.Update(time.Millisecond * 500)

			require.True(t, g.Manager.ModalVisible())
		})
	})

	t.Run("game over state", func(t *testing.T) {
		t.Run("does not transitions on any event", func(t *testing.T) {
			setup()
			g.currentState = new(gameOverState)
			expSavedState := g.currentState

			g.currentState.handle(g, PauseGame)
			g.currentState.handle(g, StartGame)

			require.Equal(t, expSavedState, g.currentState)
		})

		t.Run("shows modal for game over text", func(t *testing.T) {
			setup()
			g.currentState.handle(g, StartGame)
			g.currentState.handle(g, PauseGame)

			g.Update(time.Millisecond * 500)

			require.True(t, g.Manager.ModalVisible())
		})

		t.Run("transitions to menu state when delay expires", func(t *testing.T) {
			setup()
			g.currentState = &gameOverState{delay: time.Millisecond * 500}

			g.Update(time.Millisecond * 100)
			g.Update(time.Millisecond * 100)
			g.Update(time.Millisecond * 100)
			g.Update(time.Millisecond * 100)
			g.Update(time.Millisecond * 100)

			require.False(t, g.Manager.ModalVisible())
			require.IsType(t, &menuState{}, g.currentState)
		})
	})
}
