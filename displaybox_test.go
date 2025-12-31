package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DisplayBox(t *testing.T) {
	const width = 10
	const height = 0
	pos := Position{x: 0, y: 0}

	displayBox := NewDisplayBox(pos, height, width)

	t.Run("height includes components", func(t *testing.T) {
		require.Equal(t, 3, displayBox.Height())
	})

	t.Run("bottom is offset from position", func(t *testing.T) {
		require.Equal(t, 3, displayBox.Bottom())
	})

	t.Run("sets text of lives box", func(t *testing.T) {
		const numLives = 2
		displayBox.SetLives(numLives)

		require.Equal(t, fmt.Sprintf(livesFormat, numLives), displayBox.lives.Text())
	})

	t.Run("sets text of score box", func(t *testing.T) {
		const score = 2000
		displayBox.SetScore(score)

		require.Equal(t, fmt.Sprintf(scoreFormat, score), displayBox.score.Text())
	})
}
