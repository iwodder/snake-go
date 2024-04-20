package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CanDrawApple(t *testing.T) {
	scn := setupDefaultScreen(t)

	a := apple{
		pos: pos{x: 1, y: 1},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
}

func Test_CanDrawApples(t *testing.T) {
	scn := setupDefaultScreen(t)

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

func Test_IfAppleIsEatenThenPositionIsUpdatedAndItsNotEaten(t *testing.T) {
	a := apple{pos: pos{x: 10, y: 10}, eaten: true}

	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	a.move(b)

	require.NotEqual(t, pos{x: 10, y: 10}, a.pos)
	require.False(t, a.eaten)
	requireWithinBounds(t, b, a.pos)
}

func Test_IfAppleIsNotEatenThenPositionDoesNotChange(t *testing.T) {
	a := apple{pos: pos{x: 10, y: 10}, eaten: false}

	a.move(boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	})

	require.Equal(t, pos{x: 10, y: 10}, a.pos)
	require.False(t, a.eaten)
}

func Test_CanMoveApples(t *testing.T) {
	a := apples{
		{pos: pos{x: 1, y: 1}, eaten: true},
		{pos: pos{x: 2, y: 2}, eaten: false},
		{pos: pos{x: 3, y: 3}, eaten: true},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	a.move(b, 0)

	require.NotEqual(t, apple{pos: pos{x: 1, y: 1}, eaten: true}, a[0])
	requireWithinBounds(t, b, a[0].pos)

	require.Equal(t, apple{pos: pos{x: 2, y: 2}, eaten: false}, a[1])
	requireWithinBounds(t, b, a[1].pos)

	require.NotEqual(t, apple{pos: pos{x: 3, y: 3}, eaten: true}, a[0])
	requireWithinBounds(t, b, a[2].pos)
}

func requireWithinBounds(t *testing.T, b boundary, p pos) {
	require.Less(t, p.x, b.lowerRight.x)
	require.Less(t, p.y, b.lowerRight.y)
	require.Greater(t, p.x, b.upperLeft.x)
	require.Greater(t, p.y, b.upperLeft.y)
}
