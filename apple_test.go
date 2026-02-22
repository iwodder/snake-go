package main

import (
	"snake/ui"
	"testing"

	"github.com/stretchr/testify/require"
)

var testGame = &game{
	GameBoard: ui.NewGameBoard(ui.Position{X: 0, Y: 0}, ui.Position{X: 20, Y: 20}),
}

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

	a.Update(testGame)

	require.NotEqual(t, ui.Position{X: 10, Y: 10}, a.pos)
	require.False(t, a.eaten)
	requireWithinBounds(t, testGame.GameBoard, a.pos)
}

func Test_IfAppleIsNotEatenThenPositionDoesNotChange(t *testing.T) {
	a := apple{pos: ui.Position{X: 10, Y: 10}, eaten: false}

	a.Update(testGame)

	require.Equal(t, ui.Position{X: 10, Y: 10}, a.pos)
	require.False(t, a.eaten)
}

func Test_CanMoveApples(t *testing.T) {
	a := apples{
		{pos: ui.Position{X: 1, Y: 1}, eaten: true},
		{pos: ui.Position{X: testGame.Left() + 1, Y: testGame.Top() + 1}, eaten: false},
		{pos: ui.Position{X: 3, Y: 3}, eaten: true},
	}

	a.Update(testGame, 0)

	require.NotEqual(t, apple{pos: ui.Position{X: 1, Y: 1}, eaten: true}, a[0])
	requireWithinBounds(t, testGame.GameBoard, a[0].pos)

	require.Equal(t, apple{pos: ui.Position{X: testGame.Left() + 1, Y: testGame.Top() + 1}, eaten: false}, a[1])
	requireWithinBounds(t, testGame.GameBoard, a[1].pos)

	require.NotEqual(t, apple{pos: ui.Position{X: 3, Y: 3}, eaten: true}, a[0])
	requireWithinBounds(t, testGame.GameBoard, a[2].pos)
}

func Test_NewApplesHasSizeOfTwo(t *testing.T) {
	as := newApples(testGame.GameBoard, 2)

	require.Len(t, as, 2)
	require.NotEqual(t, as[0], as[1])
}

func requireWithinBounds(t *testing.T, b *ui.GameBoard, p ui.Position) {
	require.Truef(t, b.IsInside(p), "%#v was not inside board", p)
}
