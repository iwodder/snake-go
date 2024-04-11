package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SnakeCanDrawSelfIntoTheFrame(t *testing.T) {
	dst := setupScreen(t)
	s := newSnake(1, 1)
	s.draw(dst)

	requireEqualContents(t, 2, 1, '-', dst)
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

	requireEqualContents(t, 1, 2, '|', dst)
	requireEqualContents(t, 1, 3, '|', dst)
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

	requireEqualContents(t, 1, 0, '|', dst)
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

	requireEqualContents(t, 0, 1, '-', dst)
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

	requireEqualContents(t, 1, 2, '|', dst)
	requireEqualContents(t, 1, 3, '|', dst)
	requireEqualContents(t, 2, 3, '-', dst)
	requireEqualContents(t, 3, 3, '-', dst)
}

func Test_HeadLeftAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
			{dir: right, mag: 2},
		},
	}
	s.headLeft()

	require.Equal(t, s.segments[2], vector{dir: left, mag: 0})
}

func Test_HeadRightAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
		},
	}
	s.headRight()

	require.Equal(t, s.segments[1], vector{dir: right, mag: 0})
}

func Test_HeadUpAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: down, mag: 2},
		},
	}
	s.headUp()

	require.Equal(t, s.segments[1], vector{dir: up, mag: 0})
}

func Test_HeadDownAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: up, mag: 2},
		},
	}
	s.headDown()

	require.Equal(t, s.segments[1], vector{dir: down, mag: 0})
}

func Test_MovingCurrentDirectionDoesNotAddVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		segments: []vector{
			{dir: up, mag: 2},
		},
	}
	s.headUp()

	require.Len(t, s.segments, 1)
	require.Equal(t, s.segments[0], vector{dir: up, mag: 2})
}

func Test_MovingWhenEmptyMovement(t *testing.T) {
	s := snake{
		start:    pos{x: 1, y: 1},
		segments: []vector{},
	}
	s.headUp()

	require.Len(t, s.segments, 1)
	require.Equal(t, s.segments[0], vector{dir: up, mag: 0})
}

func Test_UpdateMovesPosInDirectionOfFirstVector(t *testing.T) {
	tests := []struct {
		name string
		m    vector
		exp  pos
	}{
		{
			name: "headDown",
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
			s.move(bounds{
				upperLeft:  pos{x: 0, y: 0},
				lowerRight: pos{x: 80, y: 80},
			})

			require.Equal(t, tt.exp, s.start)
			require.Equal(t, vector{dir: tt.m.dir, mag: tt.m.mag}, s.segments[len(s.segments)-1])
		})
	}

}

func Test_SnakeDoesNotChangeMagnitudeWhenOnlyOneSegmentExists(t *testing.T) {
	sn := newSnake(10, 10)
	sn.move(bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	})

	require.Equal(t, vector{dir: right, mag: 1}, sn.segments[0])
}

func Test_SnakeChangesMagnitudeWhenMoreThanOneSegmentExists(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		segments: []vector{
			{dir: right, mag: 1},
			{dir: down, mag: 1},
		},
	}
	s.move(bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	})

	require.Len(t, s.segments, 1)
	require.Equal(t, vector{dir: down, mag: 2}, s.segments[0])
}

func Test_SnakeWontMoveOutsideOfThe(t *testing.T) {
	tests := []struct {
		name string
		vec  vector
	}{
		{
			name: "right edge of screen",
			vec:  vector{dir: right, mag: 10},
		},
		{
			name: "left edge of screen",
			vec:  vector{dir: left, mag: 10},
		},
		{
			name: "top edge of screen",
			vec:  vector{dir: up, mag: 10},
		},
		{
			name: "bottom edge of screen",
			vec:  vector{dir: down, mag: 10},
		},
	}
	b := bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := snake{
				start:    pos{10, 10},
				segments: []vector{tt.vec},
			}

			s.move(b)

			require.Len(t, s.segments, 1)
			require.Equal(t, pos{10, 10}, s.start, "starting pos shouldn't have changed")
		})
	}
}

func Test_SnakeWontMoveUntilDirectionIsAwayFromRightEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		segments: []vector{
			{dir: right, mag: 10},
		},
	}
	b := bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b)
	s.headDown()
	s.move(b)

	require.Len(t, s.segments, 2)
	require.Equal(t, vector{dir: right, mag: 9}, s.segments[0])
	require.Equal(t, vector{dir: down, mag: 1}, s.segments[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromLeftEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		segments: []vector{
			{dir: left, mag: 10},
		},
	}
	b := bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b)
	s.headUp()
	s.move(b)

	require.Len(t, s.segments, 2)
	require.Equal(t, vector{dir: left, mag: 9}, s.segments[0])
	require.Equal(t, vector{dir: up, mag: 1}, s.segments[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromTopEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		segments: []vector{
			{dir: up, mag: 10},
		},
	}
	b := bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b)
	s.headLeft()
	s.move(b)

	require.Len(t, s.segments, 2)
	require.Equal(t, vector{dir: up, mag: 9}, s.segments[0])
	require.Equal(t, vector{dir: left, mag: 1}, s.segments[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromBottomEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		segments: []vector{
			{dir: down, mag: 10},
		},
	}
	b := bounds{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b)
	s.headRight()
	s.move(b)

	require.Len(t, s.segments, 2)
	require.Equal(t, vector{dir: down, mag: 9}, s.segments[0])
	require.Equal(t, vector{dir: right, mag: 1}, s.segments[1])

}

func Test_SnakeGrowsByEatingApples(t *testing.T) {
	s := newSnake(10, 10)
	as := apples{
		{pos: pos{x: 11, y: 10}, eaten: false},
	}
	s.eat(as)

	require.Equal(t, 2, s.segments[0].mag)
	require.Len(t, as, 1)
	require.True(t, as[0].eaten)
}

func Test_NewSnakeState(t *testing.T) {
	scn := setupScreen(t)

	sn := newSnake(40, 40)
	sn.draw(scn)

	require.Equal(t, pos{x: 40, y: 40}, sn.start)
	require.Equal(t, 24, cap(sn.segments))
	require.Equal(t, vector{dir: right, mag: 1}, sn.segments[0])
	requireEqualContents(t, 41, 40, '-', scn)
}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position (%d,%d) expected %c, but was %c", x, y, exp, act)
}

func setupScreen(t *testing.T) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	ret.SetSize(80, 80)
	return ret
}
