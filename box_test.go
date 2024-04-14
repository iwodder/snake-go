package main

import (
	"github.com/gdamore/tcell/v2"
	"testing"
)

func Test_BoxCanDraw(t *testing.T) {
	dst := setupScreen(t, 10, 10)

	b := newBorder(10, 10)

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
