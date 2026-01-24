package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
)

func Test_BoardCanDraw(t *testing.T) {
	dst := setupScreen(t, 10, 10)

	b := NewGameBoard(Position{X: 0, Y: 0}, Position{X: 9, Y: 9})

	b.Draw(dst)

	exp := [][]rune{
		{tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneURCorner},
		{tcell.RuneVLine, ' ', 'S', 'n', 'a', 'k', 'e', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, 'S', 'c', 'o', 'r', 'e', ':', ' ', '0', tcell.RuneVLine},
		{tcell.RuneVLine, 'L', 'i', 'v', 'e', 's', ':', ' ', '0', tcell.RuneVLine},
		{tcell.RuneLTee, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneRTee},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneLLCorner, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneLRCorner},
	}

	requireEqualScreen(t, exp[:], dst)
}

func Test_Board(t *testing.T) {
	t.Run("height", func(t *testing.T) {
		board := NewGameBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		require.Equal(t, 11, board.Height())
	})

	t.Run("width", func(t *testing.T) {
		board := NewGameBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		require.Equal(t, 6, board.Width())
	})

	t.Run("test edges", func(t *testing.T) {
		board := NewGameBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		t.Run("top", func(t *testing.T) {
			require.Equal(t, 5, board.Top())
		})

		t.Run("bottom", func(t *testing.T) {
			require.Equal(t, 9, board.Bottom())
		})

		t.Run("left", func(t *testing.T) {
			require.Equal(t, 1, board.Left())
		})

		t.Run("right", func(t *testing.T) {
			require.Equal(t, 4, board.Right())
		})
	})
}

func requireEqualScreen(t *testing.T, exp [][]rune, act tcell.SimulationScreen) {
	for y := range exp {
		for x := range exp[y] {
			requireEqualContents(t, x, y, exp[y][x], act)
		}
	}
}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position (x=%d,Y=%d) expected '%c', but was '%c'", x, y, exp, act)
}

func setupScreen(t *testing.T, height, width int) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	ret.SetSize(height, width)
	return ret
}
