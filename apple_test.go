package main

import (
	"testing"
)

func Test_CanDrawApple(t *testing.T) {
	scn := setupScreen(t)

	a := apple{
		pos: pos{x: 1, y: 1},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
}

func Test_CanDrawApples(t *testing.T) {
	scn := setupScreen(t)

	a := apples{
		{pos: pos{x: 1, y: 1}},
		{pos: pos{x: 2, y: 2}},
		{pos: pos{x: 3, y: 3}},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
	requireEqualContents(t, 2, 2, 'A', scn)
	requireEqualContents(t, 3, 3, 'A', scn)
}
