package main

import (
	"snake/ui"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GameBoard(t *testing.T) {
	board := newGameBoard(ui.Position{X: 0, Y: 0}, ui.Position{X: 9, Y: 9})

	t.Run("center", func(t *testing.T) {
		t.Run("upper left not at origin", func(t *testing.T) {
			board := newGameBoard(ui.Position{X: 10, Y: 20}, ui.Position{X: 19, Y: 29})

			require.Equal(t, ui.Position{X: 14, Y: 26}, board.Center())
		})

		t.Run("upper left at origin", func(t *testing.T) {
			require.Equal(t, ui.Position{X: 4, Y: 6}, board.Center())
		})
	})
	t.Run("isInside", func(t *testing.T) {
		t.Run("returns true for playable area", func(t *testing.T) {
			require.True(t, board.IsInside(ui.Position{X: 2, Y: 6}))
		})

		t.Run("returns false for left wall", func(t *testing.T) {
			require.False(t, board.IsInside(ui.Position{X: board.Left(), Y: 6}))
		})

		t.Run("returns false for right wall", func(t *testing.T) {
			require.False(t, board.IsInside(ui.Position{X: board.Right(), Y: 6}))
		})

		t.Run("returns false for top wall", func(t *testing.T) {
			require.False(t, board.IsInside(ui.Position{X: 2, Y: board.Top()}))
		})

		t.Run("returns false for bottom wall", func(t *testing.T) {
			require.False(t, board.IsInside(ui.Position{X: 2, Y: board.Bottom()}))
		})
	})
}
