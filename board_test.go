package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BoxCanDraw(t *testing.T) {
	dst := setupScreen(t, 10, 10)

	b := newBoard(Position{X: 0, Y: 0}, Position{X: 9, Y: 9})

	b.draw(dst)

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

func Test_BoxFillsEachCell(t *testing.T) {
	dst := setupScreen(t, 10, 10)

	b := newBoard(Position{X: 0, Y: 0}, Position{X: 10, Y: 10})

	b.draw(dst)

	for x := range 10 {
		for y := range 10 {
			_, _, style, _ := dst.GetContent(x, y)
			assert.Equal(t, boardStyle, style, "cell (x=%d,Y=%d) didn't have the correct style", x, y)
		}
	}
}

func Test_Board(t *testing.T) {
	t.Run("height", func(t *testing.T) {
		board := newBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		require.Equal(t, 11, board.height())
	})

	t.Run("width", func(t *testing.T) {
		board := newBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		require.Equal(t, 6, board.width())
	})

	t.Run("test edges", func(t *testing.T) {
		board := newBoard(Position{X: 0, Y: 0}, Position{X: 5, Y: 10})

		t.Run("top", func(t *testing.T) {
			require.Equal(t, 5, board.topEdge())
		})

		t.Run("bottom", func(t *testing.T) {
			require.Equal(t, 9, board.bottomEdge())
		})

		t.Run("left", func(t *testing.T) {
			require.Equal(t, 1, board.leftEdge())
		})

		t.Run("right", func(t *testing.T) {
			require.Equal(t, 4, board.rightEdge())
		})
	})
}
