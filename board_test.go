package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func Test_BoxCanDraw(t *testing.T) {
	dst := setupScreen(t, 10, 10)

	b := newBoard(pos{x: 0, y: 0}, pos{x: 10, y: 10})

	b.draw(dst)

	exp := [][]rune{
		{tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneHLine, tcell.RuneURCorner},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
		{tcell.RuneVLine, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', tcell.RuneVLine},
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

	b := newBoard(pos{x: 0, y: 0}, pos{x: 10, y: 10})

	b.draw(dst)

	for x := range 10 {
		for y := range 10 {
			_, _, style, _ := dst.GetContent(x, y)
			assert.Equal(t, boardStyle, style, "cell (x=%d,y=%d) didn't have the correct style", x, y)
		}
	}
}
