package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
)

func Test_Styles(t *testing.T) {
	t.Run("base style follows system", func(t *testing.T) {
		require.Equal(t, styles[snakeStyle], tcell.StyleDefault.Foreground(tcell.ColorGreen))
		require.Equal(t, styles[foodStyle], tcell.StyleDefault.Foreground(tcell.ColorRed))
	})
}
