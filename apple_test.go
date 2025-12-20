package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testBoard = board{
	upperLeft:  Position{x: 0, y: 0},
	lowerRight: Position{x: 20, y: 20},
	hud:        NewDisplayBox(Position{x: 0, y: 0}, 0, 20),
}

func Test_CanDrawApple(t *testing.T) {
	scn := setupDefaultScreen(t)

	a := apple{
		pos: Position{x: 1, y: 1},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
}

func Test_CanDrawApples(t *testing.T) {
	scn := setupDefaultScreen(t)

	a := apples{
		{pos: Position{x: 1, y: 1}},
		{pos: Position{x: 2, y: 2}},
		{pos: Position{x: 3, y: 3}},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
	requireEqualContents(t, 2, 2, 'A', scn)
	requireEqualContents(t, 3, 3, 'A', scn)
}

func Test_IfAppleIsEatenThenPositionIsUpdatedAndItsNotEaten(t *testing.T) {
	a := apple{pos: Position{x: 10, y: 10}, eaten: true}

	a.move(&testBoard)

	require.NotEqual(t, Position{x: 10, y: 10}, a.pos)
	require.False(t, a.eaten)
	requireWithinBounds(t, &testBoard, a.pos)
}

func Test_IfAppleIsNotEatenThenPositionDoesNotChange(t *testing.T) {
	a := apple{pos: Position{x: 10, y: 10}, eaten: false}

	a.move(&board{
		upperLeft:  Position{x: 0, y: 0},
		lowerRight: Position{x: 20, y: 20},
	})

	require.Equal(t, Position{x: 10, y: 10}, a.pos)
	require.False(t, a.eaten)
}

func Test_CanMoveApples(t *testing.T) {
	a := apples{
		{pos: Position{x: 1, y: 1}, eaten: true},
		{pos: Position{x: testBoard.leftEdge() + 1, y: testBoard.topEdge() + 1}, eaten: false},
		{pos: Position{x: 3, y: 3}, eaten: true},
	}

	a.move(&testBoard, 0)

	require.NotEqual(t, apple{pos: Position{x: 1, y: 1}, eaten: true}, a[0])
	requireWithinBounds(t, &testBoard, a[0].pos)

	require.Equal(t, apple{pos: Position{x: testBoard.leftEdge() + 1, y: testBoard.topEdge() + 1}, eaten: false}, a[1])
	requireWithinBounds(t, &testBoard, a[1].pos)

	require.NotEqual(t, apple{pos: Position{x: 3, y: 3}, eaten: true}, a[0])
	requireWithinBounds(t, &testBoard, a[2].pos)
}

func Test_NewApplesHasSizeOfTwo(t *testing.T) {
	as := newApples(&testBoard, 2)

	require.Len(t, as, 2)
	require.NotEqual(t, as[0], as[1])
}

func requireWithinBounds(t *testing.T, b *board, p Position) {
	require.Truef(t, b.isInside(p), "%#v was not inside board", p)
}
