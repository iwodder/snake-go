package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SnakeCanDrawSelfIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: right, mag: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 1, 2, '-', dst)
}

func Test_SnakeCanDrawDownDirectionIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 2, 1, '|', dst)
	requireEqualContents(t, 3, 1, '|', dst)
}

func Test_SnakeCanDrawUpDirectionIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: up, mag: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 0, 1, '|', dst)
}

func Test_SnakeCanDrawLeftDirectionIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: left, mag: 1},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 1, 0, '-', dst)
}

func Test_SnakeCanDrawMultipleVectorsIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
			{dir: right, mag: 2},
		},
	}
	s.draw(dst)

	requireEqualContents(t, 2, 1, '|', dst)
	requireEqualContents(t, 3, 1, '|', dst)
	requireEqualContents(t, 3, 2, '-', dst)
	requireEqualContents(t, 3, 3, '-', dst)
}

func Test_MoveLeftAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
			{dir: right, mag: 2},
		},
	}
	s.left()

	require.Equal(t, s.segments[2], vector{dir: left, mag: 0})
}

func Test_MoveRightAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
		},
	}
	s.right()

	require.Equal(t, s.segments[1], vector{dir: right, mag: 0})
}

func Test_MoveUpAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
		},
	}
	s.up()

	require.Equal(t, s.segments[1], vector{dir: up, mag: 0})
}

func Test_MoveDownAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: up, mag: 2},
		},
	}
	s.down()

	require.Equal(t, s.segments[1], vector{dir: down, mag: 0})
}

func Test_MovingCurrentDirectionDoesNotAddVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: up, mag: 2},
		},
	}
	s.up()

	require.Len(t, s.segments, 1)
	require.Equal(t, s.segments[0], vector{dir: up, mag: 2})
}

func Test_MovingWhenEmptyMovement(t *testing.T) {
	s := snake{
		start:    pos{x: 1, y: 1},
		segments: []vector{},
	}
	s.up()

	require.Len(t, s.segments, 1)
	require.Equal(t, s.segments[0], vector{dir: up, mag: 0})
}

func Test_UpdateMovesPosInDirectionOfFirstVectorAndGrowsLastVector(t *testing.T) {
	tests := []struct {
		name string
		m    vector
		exp  pos
	}{
		{
			name: "down",
			m:    vector{dir: down, mag: 3},
			exp:  pos{x: 40, y: 41},
		},
		{
			name: "up",
			m:    vector{dir: up, mag: 3},
			exp:  pos{x: 40, y: 39},
		},
		{
			name: "right",
			m:    vector{dir: right, mag: 3},
			exp:  pos{x: 41, y: 40},
		},
		{
			name: "left",
			m:    vector{dir: left, mag: 3},
			exp:  pos{x: 39, y: 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := snake{
				start: pos{x: 40, y: 40},
				segments: []vector{
					tt.m,
				},
			}
			s.updatePosition()

			require.Equal(t, tt.exp, s.start)
			exp := vector{
				dir: tt.m.dir,
				mag: tt.m.mag + 1,
			}
			require.Equal(t, exp, s.segments[len(s.segments)-1])
		})
	}

}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position was (%dir,%dir)", x, y)
}

func setupScreen(t *testing.T) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	return ret
}
