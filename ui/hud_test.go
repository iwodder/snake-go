package ui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Hud(t *testing.T) {
	const width = 10
	const height = 0
	pos := Position{X: 0, Y: 0}

	displayBox := NewHud(pos, height, width)

	t.Run("height includes components", func(t *testing.T) {
		require.Equal(t, 3, displayBox.Height())
	})

	t.Run("bottom is offset from position", func(t *testing.T) {
		require.Equal(t, 3, displayBox.Bottom())
	})
}
