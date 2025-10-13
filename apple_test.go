package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	b := &board{
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

	a.move(&board{
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
	b := &board{
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

func Test_NewApplesHasSizeOfTwo(t *testing.T) {
	as := newApples(&board{upperLeft: pos{x: 0, y: 0}, lowerRight: pos{x: 20, y: 20}}, 2)

	require.Len(t, as, 2)
	require.NotEqual(t, as[0], as[1])
}

func requireWithinBounds(t *testing.T, b *board, p pos) {
	require.Less(t, p.x, b.rightEdge())
	require.Less(t, p.y, b.bottomEdge())
	require.Greater(t, p.x, b.leftEdge())
	require.Greater(t, p.y, b.topEdge())
}
