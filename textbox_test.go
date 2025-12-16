package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
)

func Test_TextBox(t *testing.T) {
	text := "AAA"
	textBoxNoBorder := NewTextBox(text, tcell.StyleDefault).NoBorder()
	textBoxBorder := NewTextBox(text, tcell.StyleDefault)

	t.Run("height includes border cells", func(t *testing.T) {
		require.Equal(t, MinTextboxHeightNoBorder, textBoxNoBorder.Height())
		require.Equal(t, MinTextboxHeightWithBorder, textBoxBorder.Height())
	})

	t.Run("width includes border cells", func(t *testing.T) {
		require.Equal(t, len(text), textBoxNoBorder.Width())
		require.Equal(t, len(text)+2, textBoxBorder.Width())
	})

	t.Run("height can't be set smaller than minimum", func(t *testing.T) {
		require.Equal(t, MinTextboxHeightNoBorder, textBoxNoBorder.SetHeight(0).Height())
		require.Equal(t, MinTextboxHeightWithBorder, textBoxBorder.SetHeight(0).Height())
	})

	t.Run("reports presence of border", func(t *testing.T) {
		require.False(t, textBoxNoBorder.border)
		require.True(t, textBoxBorder.border)
	})

	t.Run("reports bottom edge text box area", func(t *testing.T) {
		pos := Position{x: 1, y: 1}
		require.Equal(t, 2, textBoxNoBorder.SetPosition(pos).BottomEdge())
		require.Equal(t, 4, textBoxBorder.SetPosition(pos).BottomEdge())
	})
}
