package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DisplayBox(t *testing.T) {
	const width = 10
	const height = 0
	pos := Position{x: 0, y: 0}

	displayBox := NewDisplayBox(pos, height, width)

	t.Run("height includes components", func(t *testing.T) {
		require.Equal(t, 1, displayBox.Height())
	})

	t.Run("bottom is offset from position", func(t *testing.T) {
		require.Equal(t, 1, displayBox.Bottom())
	})
}
