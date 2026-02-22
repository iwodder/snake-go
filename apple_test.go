package main

import (
	"snake/ui"
	"testing"

	"github.com/stretchr/testify/require"
)

var testBoard = ui.NewGameBoard(ui.Position{X: 0, Y: 0}, ui.Position{X: 20, Y: 20})

func Test_CanDrawApple(t *testing.T) {
	scn := setupDefaultScreen(t)

	a := apple{
		pos: ui.Position{X: 1, Y: 1},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
}

func Test_CanDrawApples(t *testing.T) {
	scn := setupDefaultScreen(t)

	a := apples{
		{pos: ui.Position{X: 1, Y: 1}},
		{pos: ui.Position{X: 2, Y: 2}},
		{pos: ui.Position{X: 3, Y: 3}},
	}

	a.draw(scn)

	requireEqualContents(t, 1, 1, 'A', scn)
	requireEqualContents(t, 2, 2, 'A', scn)
	requireEqualContents(t, 3, 3, 'A', scn)
}

func Test_IfAppleIsEatenThenPositionIsUpdatedAndItsNotEaten(t *testing.T) {
	a := apple{pos: ui.Position{X: 10, Y: 10}, eaten: true}

	a.move(testBoard)

	require.NotEqual(t, ui.Position{X: 10, Y: 10}, a.pos)
	require.False(t, a.eaten)
	requireWithinBounds(t, testBoard, a.pos)
}

func Test_IfAppleIsNotEatenThenPositionDoesNotChange(t *testing.T) {
	a := apple{pos: ui.Position{X: 10, Y: 10}, eaten: false}

	a.move(testBoard)

	require.Equal(t, ui.Position{X: 10, Y: 10}, a.pos)
	require.False(t, a.eaten)
}

func Test_CanMoveApples(t *testing.T) {
	a := apples{
		{pos: ui.Position{X: 1, Y: 1}, eaten: true},
		{pos: ui.Position{X: testBoard.Left() + 1, Y: testBoard.Top() + 1}, eaten: false},
		{pos: ui.Position{X: 3, Y: 3}, eaten: true},
	}

	a.move(testBoard, 0)

	require.NotEqual(t, apple{pos: ui.Position{X: 1, Y: 1}, eaten: true}, a[0])
	requireWithinBounds(t, testBoard, a[0].pos)

	require.Equal(t, apple{pos: ui.Position{X: testBoard.Left() + 1, Y: testBoard.Top() + 1}, eaten: false}, a[1])
	requireWithinBounds(t, testBoard, a[1].pos)

	require.NotEqual(t, apple{pos: ui.Position{X: 3, Y: 3}, eaten: true}, a[0])
	requireWithinBounds(t, testBoard, a[2].pos)
}

func Test_NewApplesHasSizeOfTwo(t *testing.T) {
	as := newApples(testBoard, 2)

	require.Len(t, as, 2)
	require.NotEqual(t, as[0], as[1])
}

func requireWithinBounds(t *testing.T, b *ui.GameBoard, p ui.Position) {
	require.Truef(t, b.IsInside(p), "%#v was not inside board", p)
}
