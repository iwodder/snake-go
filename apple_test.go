package main

import (
	"snake/ui"
	"testing"

	"github.com/stretchr/testify/require"
)

var testGame = &game{
	GameBoard: ui.NewGameBoard(ui.Position{X: 0, Y: 0}, ui.Position{X: 20, Y: 20}),
}

func Test_IfAppleIsEatenThenPositionIsUpdatedAndItsNotEaten(t *testing.T) {
	a := apple{
		AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: 10, Y: 10}},
		eaten:         true,
	}

	a.Update(testGame)

	require.NotEqual(t, ui.Position{X: 10, Y: 10}, a.Pos)
	require.False(t, a.eaten)
	requireWithinBounds(t, testGame.GameBoard, a.Pos)
}

func Test_IfAppleIsNotEatenThenPositionDoesNotChange(t *testing.T) {
	a := apple{
		AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: 10, Y: 10}},
		eaten:         false,
	}

	a.Update(testGame)

	require.Equal(t, ui.Position{X: 10, Y: 10}, a.Pos)
	require.False(t, a.eaten)
}

func Test_CanUpdateApples(t *testing.T) {
	as := newApples(testGame.GameBoard, 3)
	as[0].eaten = true
	as[0].Pos = ui.Position{X: 1, Y: 1}
	as[1].Pos = ui.Position{X: testGame.Left() + 1, Y: testGame.Top() + 1}
	as[2].eaten = true
	as[2].Pos = ui.Position{X: 3, Y: 3}

	as.Update(testGame, 0)

	require.NotEqual(t, apple{AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: 1, Y: 1}}, eaten: true}, as[0])
	requireWithinBounds(t, testGame.GameBoard, as[0].Pos)

	require.Equal(t, apple{AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: testGame.Left() + 1, Y: testGame.Top() + 1}}, eaten: false}, as[1])
	requireWithinBounds(t, testGame.GameBoard, as[1].Pos)

	require.NotEqual(t, apple{AppleRenderer: ui.AppleRenderer{Pos: ui.Position{X: 3, Y: 3}}, eaten: true}, as[2])
	requireWithinBounds(t, testGame.GameBoard, as[2].Pos)
}

func Test_NewApplesHasSizeOfTwo(t *testing.T) {
	as := newApples(testGame.GameBoard, 2)

	require.Len(t, as, 2)
	require.NotEqual(t, as[0], as[1])
}

func Test_Apples(t *testing.T) {
	t.Run("ForEach", func(t *testing.T) {
		t.Run("iterates over all apples", func(t *testing.T) {
			const numApples = 3
			set := make(map[*apple]struct{})
			wasRun := false

			as := newApples(testGame.GameBoard, numApples)
			as.ForEach(func(a *apple) {
				wasRun = true
				set[a] = struct{}{}
				require.NotNil(t, a, "apple was nil")
			})

			require.True(t, wasRun, "callback was not run")
			require.Equal(t, numApples, len(set))
		})
	})
}

func requireWithinBounds(t *testing.T, b *ui.GameBoard, p ui.Position) {
	require.Truef(t, b.IsInside(p), "%#v was not inside board", p)
}
