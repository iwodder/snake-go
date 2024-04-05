package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SnakeCanDrawIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: right, magnitude: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 1, 2, '-', dst)
}

func Test_SnakeCanDrawDownMovementIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 2, 1, '|', dst)
	requireEqualContents(t, 3, 1, '|', dst)
}

func Test_SnakeCanDrawUpMovementIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: up, magnitude: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 0, 1, '|', dst)
}

func Test_SnakeCanDrawLeftMovementIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: left, magnitude: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 1, 0, '-', dst)
}

func Test_SnakeCanDrawMultipleMovementsIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
			{d: right, magnitude: 2},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 2, 1, '|', dst)
	requireEqualContents(t, 3, 1, '|', dst)
	requireEqualContents(t, 3, 2, '-', dst)
	requireEqualContents(t, 3, 3, '-', dst)
}

func Test_MoveLeftAppendsNewMagnitude(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
			{d: right, magnitude: 2},
		},
	}
	s.left()

	require.Equal(t, s.movement[2], move{d: left, magnitude: 0})
}

func Test_MoveRightAppendsNewMagnitude(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
		},
	}
	s.right()

	require.Equal(t, s.movement[1], move{d: right, magnitude: 0})
}

func Test_MoveUpAppendsNewMagnitude(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: down, magnitude: 2},
		},
	}
	s.up()

	require.Equal(t, s.movement[1], move{d: up, magnitude: 0})
}

func Test_MoveDownAppendsNewMagnitude(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: up, magnitude: 2},
		},
	}
	s.down()

	require.Equal(t, s.movement[1], move{d: down, magnitude: 0})
}

func Test_MovingCurrentDirectionDoesNotChangeMovement(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		movement: []move{
			{d: up, magnitude: 2},
		},
	}
	s.up()

	require.Len(t, s.movement, 1)
	require.Equal(t, s.movement[0], move{d: up, magnitude: 2})
}

func Test_MovingWhenEmptyMovement(t *testing.T) {
	s := snake{
		start:    pos{x: 1, y: 1},
		movement: []move{},
	}
	s.up()

	require.Len(t, s.movement, 1)
	require.Equal(t, s.movement[0], move{d: up, magnitude: 0})
}

func Test_UpdateMovesPosInDirectionOfFirstMovementAndGrowsLastMagnitude(t *testing.T) {
	tests := []struct {
		name string
		m    move
		exp  pos
	}{
		{
			name: "down",
			m:    move{d: down, magnitude: 3},
			exp:  pos{x: 40, y: 41},
		},
		{
			name: "up",
			m:    move{d: up, magnitude: 3},
			exp:  pos{x: 40, y: 39},
		},
		{
			name: "right",
			m:    move{d: right, magnitude: 3},
			exp:  pos{x: 41, y: 40},
		},
		{
			name: "left",
			m:    move{d: left, magnitude: 3},
			exp:  pos{x: 39, y: 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := snake{
				start: pos{x: 40, y: 40},
				movement: []move{
					tt.m,
				},
			}
			s.updatePosition()

			require.Equal(t, tt.exp, s.start)
			exp := move{
				d:         tt.m.d,
				magnitude: tt.m.magnitude + 1,
			}
			require.Equal(t, exp, s.movement[len(s.movement)-1])
		})
	}

}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position was (%d,%d)", x, y)
}

func setupScreen(t *testing.T) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	return ret
}

//func printFrame(f *frame) {
//	for x := 0; x < len(f); x++ {
//		for y := 0; y < len(f[x]); y++ {
//			fmt.Printf("%c", f[x][y])
//			if y < len(f[x])-1 {
//				fmt.Printf(" ")
//			}
//		}
//		fmt.Println()
//	}
//}
