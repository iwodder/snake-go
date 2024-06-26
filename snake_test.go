package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const moveDelta = time.Millisecond * 250

func Test_SnakeCanDrawOntoTheScreen(t *testing.T) {
	dst := setupDefaultScreen(t)

	s := snake{
		start: pos{x: 1, y: 1},
		vecs: []vector{
			{dir: down, mag: 2, r: dirRunes[down]},
			{dir: right, mag: 2, r: dirRunes[right]},
			{dir: up, mag: 3, r: dirRunes[up]},
			{dir: left, mag: 3, r: dirRunes[left]},
		},
	}
	s.draw(dst)

	exp := [][]rune{
		{'-', '-', '-', '|', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', '|', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', '|', ' ', '|', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', '|', '-', '-', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	}

	requireEqualScreen(t, exp[:], dst)
}

func Test_HeadLeftAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		vecs: []vector{
			{dir: down, mag: 2, r: dirRunes[down]},
		},
	}
	s.headLeft()

	require.Equal(t, vector{dir: left, mag: 0, r: dirRunes[left]}, s.vecs[1])
}

func Test_HeadRightAppendsNewVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		vecs: []vector{
			{dir: down, mag: 2, r: dirRunes[down]},
		},
	}
	s.headRight()

	require.Equal(t, s.vecs[1], vector{dir: right, mag: 0, r: dirRunes[right]})
}

func Test_HeadUpAppendsNewVector(t *testing.T) {
	s := newSnake(pos{10, 10})
	s.headUp()

	require.Equal(t, s.vecs[1], vector{dir: up, mag: 0, r: dirRunes[up]})
}

func Test_HeadDownAppendsNewVector(t *testing.T) {
	s := newSnake(pos{x: 10, y: 10})
	s.headDown()

	require.Equal(t, s.vecs[1], vector{dir: down, mag: 0, r: dirRunes[down]})
}

func Test_MovingCurrentDirectionDoesNotAddVector(t *testing.T) {
	s := snake{
		start: pos{x: 1, y: 1},
		vecs: []vector{
			{dir: up, mag: 2},
		},
	}
	s.headUp()

	require.Len(t, s.vecs, 1)
	require.Equal(t, s.vecs[0], vector{dir: up, mag: 2})
}

func Test_ChangingDirectionDoesNotAddVectorIfMagnitudeIsZero(t *testing.T) {
	s := newSnake(pos{x: 10, y: 10})

	s.headUp()
	s.headUp()
	s.headRight()
	s.headUp()
	s.headDown()
	s.headDown()
	s.headLeft()
	s.headUp()

	require.Len(t, s.vecs, 2)
	require.Equal(t, vector{dir: right, mag: 1, r: dirRunes[right]}, s.vecs[0])
	require.Equal(t, vector{dir: up, mag: 0, r: dirRunes[up]}, s.vecs[1])
}

func Test_SnakeCantDoubleBackOnSelfRightToLeft(t *testing.T) {
	s := newSnake(pos{x: 10, y: 10})

	s.headLeft()

	require.Len(t, s.vecs, 1)
	require.Equal(t, vector{dir: right, mag: 1, r: dirRunes[right]}, s.vecs[0])
}

func Test_SnakeCantDoubleBackOnSelfLeftToRight(t *testing.T) {
	s := snake{
		vecs: []vector{
			{dir: left, mag: 1},
		},
	}

	s.headRight()

	require.Len(t, s.vecs, 1)
	require.Equal(t, vector{dir: left, mag: 1}, s.vecs[0])
}

func Test_SnakeCantDoubleBackOnSelfUpToDown(t *testing.T) {
	s := snake{
		vecs: []vector{
			{dir: up, mag: 1},
		},
	}

	s.headDown()

	require.Len(t, s.vecs, 1)
	require.Equal(t, vector{dir: up, mag: 1}, s.vecs[0])
}

func Test_SnakeCantDoubleBackOnSelfDownToUp(t *testing.T) {
	s := snake{
		vecs: []vector{
			{dir: down, mag: 1},
		},
	}

	s.headUp()

	require.Len(t, s.vecs, 1)
	require.Equal(t, vector{dir: down, mag: 1}, s.vecs[0])
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
				vecs: []vector{
					tt.m,
				},
			}
			s.move(boundary{
				upperLeft:  pos{x: 0, y: 0},
				lowerRight: pos{x: 80, y: 80},
			}, moveDelta)

			require.Equal(t, tt.exp, s.start)
			require.Equal(t, vector{dir: tt.m.dir, mag: tt.m.mag}, s.vecs[len(s.vecs)-1])
		})
	}

}

func Test_MoveDoesNotChangeMagnitudeWhenOnlyOneSegmentExists(t *testing.T) {
	sn := newSnake(pos{x: 10, y: 10})
	sn.move(boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}, moveDelta)

	require.Equal(t, vector{dir: right, mag: 1, r: dirRunes[right]}, sn.vecs[0])
}

func Test_MoveDoesChangeMagnitudeWhenMoreThanOneSegmentExists(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		vecs: []vector{
			{dir: right, mag: 1},
			{dir: down, mag: 1},
		},
	}
	s.move(boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}, moveDelta)

	require.Len(t, s.vecs, 1)
	require.Equal(t, vector{dir: down, mag: 2}, s.vecs[0])
}

func Test_SnakeWontMoveOutsideOfThe(t *testing.T) {
	tests := []struct {
		name string
		vec  vector
	}{
		{
			name: "right edge of screen",
			vec:  vector{dir: right, mag: 9},
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
			vec:  vector{dir: down, mag: 9},
		},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := snake{
				start: pos{10, 10},
				vecs:  []vector{tt.vec},
			}

			s.move(b, moveDelta)

			require.Len(t, s.vecs, 1)
			require.Equal(t, pos{10, 10}, s.start, "starting pos shouldn't have changed")
		})
	}
}

func Test_SnakeWontMoveUntilDirectionIsAwayFromRightEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		vecs: []vector{
			{dir: right, mag: 10, r: dirRunes[right]},
		},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b, moveDelta)
	s.headDown()
	s.move(b, moveDelta)

	require.Len(t, s.vecs, 2)
	require.Equal(t, vector{dir: right, mag: 9, r: dirRunes[right]}, s.vecs[0])
	require.Equal(t, vector{dir: down, mag: 1, r: dirRunes[down]}, s.vecs[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromLeftEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		vecs: []vector{
			{dir: left, mag: 10, r: dirRunes[left]},
		},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b, moveDelta)
	s.headUp()
	s.move(b, moveDelta)

	require.Len(t, s.vecs, 2)
	require.Equal(t, vector{dir: left, mag: 9, r: dirRunes[left]}, s.vecs[0])
	require.Equal(t, vector{dir: up, mag: 1, r: dirRunes[up]}, s.vecs[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromTopEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		vecs: []vector{
			{dir: up, mag: 10, r: dirRunes[up]},
		},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b, moveDelta)
	s.headLeft()
	s.move(b, moveDelta)

	require.Len(t, s.vecs, 2)
	require.Equal(t, vector{dir: up, mag: 9, r: dirRunes[up]}, s.vecs[0])
	require.Equal(t, vector{dir: left, mag: 1, r: dirRunes[left]}, s.vecs[1])

}

func Test_SnakeWontMoveUntilDirectionIsAwayFromBottomEdgeOfScreen(t *testing.T) {
	s := snake{
		start: pos{10, 10},
		vecs: []vector{
			{dir: down, mag: 10, r: dirRunes[down]},
		},
	}
	b := boundary{
		upperLeft:  pos{x: 0, y: 0},
		lowerRight: pos{x: 20, y: 20},
	}
	s.move(b, moveDelta)
	s.headRight()
	s.move(b, moveDelta)

	require.Len(t, s.vecs, 2)
	require.Equal(t, vector{dir: down, mag: 9, r: dirRunes[down]}, s.vecs[0])
	require.Equal(t, vector{dir: right, mag: 1, r: dirRunes[right]}, s.vecs[1])

}

func Test_SnakeGrowsByEatingApples(t *testing.T) {
	s := newSnake(pos{x: 10, y: 10})
	as := apples{
		{pos: pos{x: 11, y: 10}, eaten: false},
	}
	s.eat(as)

	require.Equal(t, 2, s.vecs[0].mag)
	require.Len(t, as, 1)
	require.True(t, as[0].eaten)
}

func Test_SnakeMovesFourSquaresPerSecond(t *testing.T) {
	rate := time.Second / 20
	ticker := time.NewTicker(rate)

	s := newSnake(pos{x: 10, y: 10})
	ticks := 0
	for range ticker.C {
		s.move(boundary{upperLeft: pos{x: 0, y: 0}, lowerRight: pos{x: 20, y: 20}}, rate)
		ticks += 1
		if ticks == 20 {
			ticker.Stop()
			break
		}
	}
	require.Equal(t, pos{x: 14, y: 10}, s.start, "starting pos should have moved 4 squares right")
}

func Test_NewSnakeState(t *testing.T) {
	scn := setupDefaultScreen(t)

	sn := newSnake(pos{x: 40, y: 40})
	sn.draw(scn)

	require.Equal(t, pos{x: 40, y: 40}, sn.start)
	require.Equal(t, 24, cap(sn.vecs))
	require.Equal(t, vector{dir: right, mag: 1, r: dirRunes[right]}, sn.vecs[0])
	requireEqualContents(t, 41, 40, '-', scn)
}

func Test_BoundsMethods(t *testing.T) {
	tests := []struct {
		name   string
		method func(boundary) int
		exp    int
	}{
		{
			name:   "top edge",
			method: boundary.topEdge,
			exp:    2,
		},
		{
			name:   "left edge",
			method: boundary.leftEdge,
			exp:    1,
		},
		{
			name:   "right edge",
			method: boundary.rightEdge,
			exp:    19,
		},
		{
			name:   "bottom edge",
			method: boundary.bottomEdge,
			exp:    19,
		},
		{
			name:   "height",
			method: boundary.height,
			exp:    18,
		},
		{
			name:   "width",
			method: boundary.width,
			exp:    19,
		},
	}
	b := boundary{
		upperLeft:  pos{x: 1, y: 2},
		lowerRight: pos{x: 20, y: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, tt.method(b))
		})
	}
}

func Test_ShrinkBounds(t *testing.T) {
	b := boundary{
		upperLeft:  pos{x: 1, y: 1},
		lowerRight: pos{x: 20, y: 20},
	}
	exp := boundary{
		upperLeft:  pos{x: 2, y: 2},
		lowerRight: pos{x: 19, y: 19},
	}
	require.Equal(t, exp, b.shrink(1))
}

func Test_BoundsCenter(t *testing.T) {
	b := boundary{
		upperLeft: pos{
			x: 0, y: 0,
		},
		lowerRight: pos{
			x: 20, y: 20,
		},
	}

	require.Equal(t, pos{10, 10}, b.center())
}

func Test_SnakeRespondsToKeyEvents(t *testing.T) {
	tests := []struct {
		name   string
		ev     *tcell.EventKey
		snake  snake
		expDir direction
	}{
		{
			name:   "down arrow",
			snake:  snake{vecs: []vector{{dir: right, mag: 1, r: dirRunes[right]}}},
			ev:     tcell.NewEventKey(tcell.KeyDown, tcell.RuneDArrow, tcell.ModNone),
			expDir: down,
		},
		{
			name:   "up arrow",
			snake:  snake{vecs: []vector{{dir: right, mag: 1, r: dirRunes[right]}}},
			ev:     tcell.NewEventKey(tcell.KeyUp, tcell.RuneUArrow, tcell.ModNone),
			expDir: up,
		},
		{
			name:   "left arrow",
			snake:  snake{vecs: []vector{{dir: up, mag: 1, r: dirRunes[up]}}},
			ev:     tcell.NewEventKey(tcell.KeyLeft, tcell.RuneLArrow, tcell.ModNone),
			expDir: left,
		},
		{
			name:   "right arrow",
			snake:  snake{vecs: []vector{{dir: up, mag: 1, r: dirRunes[up]}}},
			ev:     tcell.NewEventKey(tcell.KeyRight, tcell.RuneRArrow, tcell.ModNone),
			expDir: right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.snake.notify(tt.ev)

			require.Len(t, tt.snake.vecs, 2)
			require.Equal(t, tt.expDir, tt.snake.head().dir, "head direction didn't match")
		})
	}
}

func requireEqualScreen(t *testing.T, exp [][]rune, act tcell.SimulationScreen) {
	for y := range exp {
		for x := range exp[y] {
			requireEqualContents(t, x, y, exp[y][x], act)
		}
	}
}

func requireEqualContents(t *testing.T, x, y int, exp rune, scn tcell.SimulationScreen) {
	act, _, _, _ := scn.GetContent(x, y)
	require.EqualValues(t, exp, act, "position (%d,%d) expected '%c', but was '%c'", x, y, exp, act)
}

func setupDefaultScreen(t *testing.T) tcell.SimulationScreen {
	return setupScreen(t, 80, 80)
}

func setupScreen(t *testing.T, height, width int) tcell.SimulationScreen {
	ret := tcell.NewSimulationScreen("")
	require.NoError(t, ret.Init())
	ret.SetSize(height, width)
	return ret
}
