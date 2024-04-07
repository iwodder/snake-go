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
